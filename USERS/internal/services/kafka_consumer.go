package services

import (
	"context"
	"encoding/json"
	"log"
	"time"

	pb "github.com/GameXost/YandexGo_proj/USERS/API/generated/clients"
	"github.com/segmentio/kafka-go"
)

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

			case "ride_accepted":
				var event RideAcceptedEvent
				if err := json.Unmarshal(m.Value, &event); err != nil {
					log.Printf("Kafka unmarshal ride_accepted error: %v", err)
					continue
				}
				log.Printf("Ride accepted: ride_id=%s, driver_id=%s", event.RideID, event.DriverID)

			case "ride_completed":
				var event RideCompletedEvent
				if err := json.Unmarshal(m.Value, &event); err != nil {
					log.Printf("Kafka unmarshal ride_completed error: %v", err)
					continue
				}
				log.Printf("Ride completed: ride_id=%s, driver_id=%s", event.RideID, event.DriverID)

			case "ride_canceled":
				var event RideCanceledEvent
				if err := json.Unmarshal(m.Value, &event); err != nil {
					log.Printf("Kafka unmarshal ride_canceled error: %v", err)
					continue
				}
				log.Printf("Ride canceled: ride_id=%s, driver_id=%s, reason=%s", event.RideID, event.DriverID, event.Reason)

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

func (s *UserService) sendKafkaResponse(ctx context.Context, topic string, event interface{}) error {
	data, err := json.Marshal(event)
	if err != nil {
		return err
	}
	msg := kafka.Message{
		Value: data,
		Topic: topic,
	}
	return s.Kafka.WriteMessages(ctx, msg)
}
