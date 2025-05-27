package services

import (
	"context"
	"encoding/json"
	"log"
	"time"

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
			var baseEvent struct {
				Event string `json:"event"`
			}
			if err := json.Unmarshal(m.Value, &baseEvent); err != nil {
				log.Printf("Kafka unmarshal error: %v", err)
				continue
			}
			switch baseEvent.Event {
			case "ride_created":
				var event RideCreatedEvent
				if err := json.Unmarshal(m.Value, &event); err != nil {
					log.Printf("Kafka unmarshal ride_created error: %v", err)
					continue
				}
				// TODO: сохранить информацию о поездке в БД пользователя
				log.Printf("Ride created: ride_id=%s, passenger_id=%s", event.RideID, event.PassengerID)

			case "ride_accepted":
				var event RideAcceptedEvent
				if err := json.Unmarshal(m.Value, &event); err != nil {
					log.Printf("Kafka unmarshal ride_accepted error: %v", err)
					continue
				}
				// TODO: обновить поездку в БД пользователя: назначить водителя, статус на "accepted"
				log.Printf("Ride accepted: ride_id=%s, driver_id=%s", event.RideID, event.DriverID)

			case "ride_completed":
				var event RideCompletedEvent
				if err := json.Unmarshal(m.Value, &event); err != nil {
					log.Printf("Kafka unmarshal ride_completed error: %v", err)
					continue
				}
				// TODO: обновить статус поездки в БД пользователя на "completed"
				log.Printf("Ride completed: ride_id=%s, driver_id=%s", event.RideID, event.DriverID)

			case "ride_canceled":
				var event RideCanceledEvent
				if err := json.Unmarshal(m.Value, &event); err != nil {
					log.Printf("Kafka unmarshal ride_canceled error: %v", err)
					continue
				}
				// TODO: обновить статус поездки в БД пользователя на "canceled"
				log.Printf("Ride canceled: ride_id=%s, driver_id=%s, reason=%s", event.RideID, event.DriverID, event.Reason)

			default:
				log.Printf("Unknown event type: %s", baseEvent.Event)
			}
		}
	}()
}
