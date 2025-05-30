package services

import (
	"context"
	"encoding/json"
	"time"

	"github.com/segmentio/kafka-go"
)

// Универсальная функция для отправки любого события
func (s *DriverService) PublishEvent(ctx context.Context, topic string, event interface{}, key string) error {
	data, err := json.Marshal(event)
	if err != nil {
		return err
	}
	msg := kafka.Message{
		Key:   []byte(key),
		Value: data,
	}
	return s.Kafka.WriteMessages(ctx, msg)
}

func (s *DriverService) PublishRideAccepted(ctx context.Context, event RideAcceptedEvent) error {
	event.Event = "ride_accepted"
	event.Timestamp = time.Now().Unix()
	return s.PublishEvent(ctx, "ride-events", event, event.RideID)
}

func (s *DriverService) PublishRideCompleted(ctx context.Context, event RideCompletedEvent) error {
	event.Event = "ride_completed"
	event.Timestamp = time.Now().Unix()
	return s.PublishEvent(ctx, "ride-events", event, event.RideID)
}

func (s *DriverService) PublishRideCanceled(ctx context.Context, event RideCanceledEvent) error {
	event.Event = "ride_canceled"
	event.Timestamp = time.Now().Unix()
	return s.PublishEvent(ctx, "ride-events", event, event.RideID)
}
