package services

import (
	"context"
	"encoding/json"
	"log"
	"time"

	pb "github.com/GameXost/YandexGo_proj/USERS/API/generated/clients"
	"github.com/segmentio/kafka-go"
)

type RideWaiter struct {
	Ch      chan string // канал для сигнализации о принятии или отмене
	Timeout time.Duration
}

func (s *UserService) StartKafkaConsumer(ctx context.Context, reader *kafka.Reader) {
	go func() {
		for {
			m, err := reader.ReadMessage(ctx)
			if err != nil {
				log.Printf("Kafka can't read msg, err : %v", err)
				time.Sleep(time.Second)
				continue
			}
			var baseEvent BaseEvent
			if err := json.Unmarshal(m.Value, &baseEvent); err != nil {
				log.Printf("kafka err : %v", err)
				continue
			}
			switch baseEvent.Event {
			case "ride_created":
				var event RideCreatedEvent
				if err := json.Unmarshal(m.Value, &event); err != nil {
					log.Printf("Kafka unmarshal ride_created error: %v", err)
					continue
				}
				log.Printf("Ride created: ride_id=%s, passenger_id=%s", event.RideID, event.PassengerID)
				s.startRideWaiter(ctx, event.RideID, 600*time.Second) // 30 секунд ожидания

			case "ride_accepted":
				var event RideAcceptedEvent
				if err := json.Unmarshal(m.Value, &event); err != nil {
					log.Printf("Kafka unmarshal ride_accepted error: %v", err)
					continue
				}
				s.signalRide(event.RideID, "accepted")
				log.Printf("Ride accepted: ride_id=%s, driver_id=%s", event.RideID, event.DriverID)

			case "ride_canceled":
				var event RideCanceledEvent
				if err := json.Unmarshal(m.Value, &event); err != nil {
					log.Printf("Kafka unmarshal ride_canceled error: %v", err)
					continue
				}
				s.signalRide(event.RideID, "canceled")
				log.Printf("Ride canceled: ride_id=%s, driver_id=%s, reason=%s", event.RideID, event.DriverID, event.Reason)

			case "ride_completed":
				var event RideCompletedEvent
				if err := json.Unmarshal(m.Value, &event); err != nil {
					log.Printf("Kafka unmarshal ride_completed error: %v", err)
					continue
				}
				log.Printf("Ride completed: ride_id=%s, driver_id=%s", event.RideID, event.DriverID)

			case "get_user_profile":
				var req struct {
					BaseEvent
					UserID string `json:"user_id"`
				}
				if err := json.Unmarshal(m.Value, &req); err != nil {
					log.Printf("kafka is unable to unmarshal get_user_profile event, err : %v", err)
					continue
				}
				user, err := s.GetUserProfile(ctx, req.UserID)
				if err != nil {
					log.Printf("Failed to get user profile: %v", err)
					continue
				}
				if req.ReplyTo != "" && req.CorrelationID != "" {
					response := struct {
						BaseEvent
						User *pb.User `json:"user"`
					}{
						BaseEvent: BaseEvent{
							Event:         "user_profile_response",
							CorrelationID: req.CorrelationID,
							Timestamp:     time.Now().Unix(),
						},
						User: user,
					}
					err := s.sendKafkaResponse(ctx, req.ReplyTo, response)
					if err != nil {
						log.Printf("Failed to send Kafka response: %v", err)
					}
				}

			default:
				log.Printf("Unknown event type: %s", baseEvent.Event)
			}
		}
	}()
}

// Запуск ожидания принятия заказа с таймаутом
func (s *UserService) startRideWaiter(ctx context.Context, rideID string, timeout time.Duration) {
	s.WaitersMutex.Lock()
	if s.OrderWaiters == nil {
		s.OrderWaiters = make(map[string]*RideWaiter)
	}
	if _, exists := s.OrderWaiters[rideID]; exists {
		s.WaitersMutex.Unlock()
		return // уже есть
	}
	waiter := &RideWaiter{Ch: make(chan string, 1), Timeout: timeout}
	s.OrderWaiters[rideID] = waiter
	s.WaitersMutex.Unlock()

	go func() {
		select {
		case result := <-waiter.Ch:
			log.Printf("Ride %s finished with result: %s", rideID, result)
			// Здесь можно обновить статус заказа в БД или уведомить пользователя
		case <-time.After(timeout):
			log.Printf("Ride %s timed out, canceling", rideID)
			// Отправляем событие отмены
			cancelEvent := RideCanceledEvent{
				BaseEvent: BaseEvent{
					Event:     "ride_canceled",
					Timestamp: time.Now().Unix(),
				},
				RideID: rideID,
				Reason: "timeout",
			}
			if s.Kafka != nil {
				err := s.PublishEvent(ctx, "ride-events", cancelEvent, rideID)
				if err != nil {
					log.Printf("Failed to publish ride_canceled: %v", err)
				}
			}
		}
		s.WaitersMutex.Lock()
		delete(s.OrderWaiters, rideID)
		s.WaitersMutex.Unlock()
	}()
}

// Сигнализировать о принятии или отмене заказа
func (s *UserService) signalRide(rideID, result string) {
	s.WaitersMutex.Lock()
	defer s.WaitersMutex.Unlock()
	if waiter, ok := s.OrderWaiters[rideID]; ok {
		select {
		case waiter.Ch <- result:
		default:
		}
	}
}

func (s *UserService) sendKafkaResponse(ctx context.Context, topic string, event interface{}) error {
	data, err := json.Marshal(event)
	if err != nil {
		return err
	}
	msg := kafka.Message{
		Value: data,
	}
	return s.Kafka.WriteMessages(ctx, msg)
}

func WaitForKafkaReply(ctx context.Context, brokers []string, replyTo, correlationId string, timeout time.Duration) (*pb.Driver, error) {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: brokers,
		Topic:   replyTo,
		GroupID: "", // индивидуальный consumer
	})
	defer reader.Close()

	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	for {
		m, err := reader.ReadMessage(ctx)
		if err != nil {
			return nil, err
		}
		var resp struct {
			BaseEvent
			Driver *pb.Driver `json:"driver"`
		}
		if err := json.Unmarshal(m.Value, &resp); err != nil {
			continue
		}
		if resp.CorrelationID == correlationId {
			return resp.Driver, nil
		}
	}
}
