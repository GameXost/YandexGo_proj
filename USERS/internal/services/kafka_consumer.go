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
			case "ride_completed":
				var event RideCompletedEvent
				if err := json.Unmarshal(m.Value, &event); err != nil {
					log.Printf("Kafka unmarshal ride_completed error: %v", err)
					continue
				}
				// 1. Найти поездку в БД пользователя по event.RideID
				// 2. Обновить статус поездки на "completed"
				// 3. (Опционально) Уведомить пользователя
				log.Printf("Ride completed: ride_id=%s, driver_id=%s", event.RideID, event.DriverID)
			case "ride_canceled":
				var event RideCanceledEvent
				if err := json.Unmarshal(m.Value, &event); err != nil {
					log.Printf("Kafka unmarshal ride_canceled error: %v", err)
					continue
				}
				// 1. Найти поездку в БД пользователя по event.RideID
				// 2. Обновить статус поездки на "canceled"
				// 3. (Опционально) Уведомить пользователя
				log.Printf("Ride canceled: ride_id=%s, driver_id=%s, reason=%s", event.RideID, event.DriverID, event.Reason)
				// Добавь другие события, если нужно
			}
		}
	}()
}
