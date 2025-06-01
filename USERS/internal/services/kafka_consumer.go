package services

import (
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

type RideWaiter struct {
	Ch      chan string
	Timeout time.Duration
}

func (s *UserService) StartKafkaConsumer(ctx context.Context, reader *kafka.Reader) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				log.Println("Closing Kafka Consumer...")
				return
			default:
				m, err := reader.ReadMessage(ctx)
				if err != nil {
					if ctx.Err() != nil {
						log.Printf("Kafka consumer: context done during ReadMessage: %v", ctx.Err())
						return
					}
					log.Printf("Kafka reader error: %v", err)
					time.Sleep(time.Second)
					continue
				}
				var baseEvent BaseEvent
				if err := json.Unmarshal(m.Value, &baseEvent); err != nil {
					log.Printf("Kafka consumer: failed to unmarshal base event: %v", err)
					continue
				}
				s.ResponseWaitersMutex.Lock()
				responseChan, found := s.ResponseWaiters[baseEvent.CorrelationID]
				s.ResponseWaitersMutex.Unlock()

				if found {
					select {
					case responseChan <- m.Value:
						log.Printf("Kafka consumer: routed response with CorrelationID: %s to waiting channel.", baseEvent.CorrelationID)
					default:
						log.Printf("Kafka consumer: channel for CorrelationID %s not ready or closed, response dropped.", baseEvent.CorrelationID)
					}
					continue
				}
				switch baseEvent.Event {
				case "ride_created":
					var event RideCreatedEvent
					if err := json.Unmarshal(m.Value, &event); err != nil {
						log.Printf("Kafka consumer: unmarshal ride_created error: %v", err)
						continue
					}
					log.Printf("Event: Ride created - RideID: %s, PassengerID: %s", event.RideID, event.PassengerID)
					s.startRideWaiter(ctx, event.RideID, 600*time.Second)
				case "ride_accepted":
					var event RideAcceptedEvent
					if err := json.Unmarshal(m.Value, &event); err != nil {
						log.Printf("Kafka consumer: unmarshal ride_accepted error: %v", err)
						continue
					}
					log.Printf("Event: Ride accepted - RideID: %s, DriverID: %s", event.RideID, event.DriverID)
					s.signalRide(event.RideID, "accepted")
				case "ride_canceled":
					var event RideCanceledEvent
					if err := json.Unmarshal(m.Value, &event); err != nil {
						log.Printf("Kafka consumer: unmarshal ride_canceled error: %v", err)
						continue
					}
					log.Printf("Event: Ride canceled - RideID: %s, DriverID: %s, Reason: %s", event.RideID, event.DriverID, event.Reason)
					s.signalRide(event.RideID, "canceled")

				case "ride_completed":
					var event RideCompletedEvent
					if err := json.Unmarshal(m.Value, &event); err != nil {
						log.Printf("Kafka consumer: unmarshal ride_completed error: %v", err)
						continue
					}
					log.Printf("Event: Ride completed - RideID: %s, DriverID: %s", event.RideID, event.DriverID)
					s.signalRide(event.RideID, "completed")
				case "get_user_profile_request":
					var req GetUserProfileRequest
					if err := json.Unmarshal(m.Value, &req); err != nil {
						log.Printf("Kafka consumer: failed to unmarshal get_user_profile_request: %v", err)
						continue
					}
					log.Printf("Kafka consumer: Received get_user_profile_request for UserID: %s with CorrelationID: %s", req.UserID, req.CorrelationID)

					user, err := s.Repo.GetUserByID(ctx, req.UserID)
					var userProfileResp UserProfileResponse
					if err != nil {
						log.Printf("Failed to get user profile for UserID %s: %v", req.UserID, err)
						userProfileResp = UserProfileResponse{
							BaseEvent: NewBaseEvent("user_profile_response", req.CorrelationID, ""),
							Error:     err.Error(),
						}
					} else {
						userProfileResp = UserProfileResponse{
							BaseEvent: NewBaseEvent("user_profile_response", req.CorrelationID, ""),
							User:      modelToProtoUser(user),
						}
					}

					if req.ReplyTo != "" {
						if sendErr := s.sendKafkaMessage(ctx, req.ReplyTo, userProfileResp); sendErr != nil {
							log.Printf("Failed to send user_profile_response to topic %s: %v", req.ReplyTo, sendErr)
						} else {
							log.Printf("Sent user_profile_response for CorrelationID %s to topic %s", req.CorrelationID, req.ReplyTo)
						}
					} else {
						log.Printf("Cannot send user_profile_response for CorrelationID %s: ReplyTo topic is empty.", req.CorrelationID)
					}

				default:
					log.Printf("Kafka consumer: Unknown event type: %s with CorrelationID: %s. Value: %s", baseEvent.Event, baseEvent.CorrelationID, string(m.Value))
				}
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
		defer func() {
			s.WaitersMutex.Lock()
			delete(s.OrderWaiters, rideID)
			close(waiter.Ch)
			s.WaitersMutex.Unlock()
		}()
		select {
		case result := <-waiter.Ch:
			log.Printf("Ride %s finished with result: %s", rideID, result)
		case <-time.After(timeout):
			log.Printf("Ride %s timed out, canceling", rideID)
			cancelEvent := RideCanceledEvent{
				BaseEvent: NewBaseEvent("ride_canceled", uuid.New().String(), ""),
				RideID:    rideID,
				Reason:    "timeout",
			}
			if s.Kafka != nil {
				err := s.sendKafkaMessage(ctx, s.RidesTopic, cancelEvent)
				if err != nil {
					log.Printf("Failed to publish ride_canceled: %v", err)
				}
			}
		case <-ctx.Done():
			log.Printf("Ride %s cancelled", rideID)
			return
		}
	}()
}

func (s *UserService) signalRide(rideID, result string) {
	s.WaitersMutex.Lock()
	defer s.WaitersMutex.Unlock()
	if waiter, ok := s.OrderWaiters[rideID]; ok {
		select {
		case waiter.Ch <- result:
		default:
			log.Printf("Ride %s waiter channel not ready to receive signal: %s", rideID, result)
		}
	} else {
		log.Printf("No waiter found for ride ID: %s", rideID)
	}
}
