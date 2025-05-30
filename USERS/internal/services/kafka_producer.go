package services

import (
	"context"
	"encoding/json"
	"time"

	"github.com/segmentio/kafka-go"
)

// Универсальная функция для отправки любого события
func (s *UserService) PublishEvent(ctx context.Context, topic string, event interface{}, key string) error {
	data, err := json.Marshal(event)
	if err != nil {
		return err
	}
	msg := s.kafkaMessage(key, data)
	return s.Kafka.WriteMessages(ctx, msg)
}

func (s *UserService) kafkaMessage(key string, value []byte) kafka.Message {
	return kafka.Message{
		Key:   []byte(key),
		Value: value,
	}
}

// Обёртка для создания поездки
func (s *UserService) PublishRideCreated(ctx context.Context, event RideCreatedEvent) error {
	event.Event = "ride_created"
	event.Timestamp = time.Now().Unix()
	return s.PublishEvent(ctx, "ride-events", event, event.RideID)
}

// Обёртка для принятия поездки
func (s *UserService) PublishRideAccepted(ctx context.Context, event RideAcceptedEvent) error {
	event.Event = "ride_accepted"
	event.Timestamp = time.Now().Unix()
	return s.PublishEvent(ctx, "ride-events", event, event.RideID)
}

// Обёртка для завершения поездки
func (s *UserService) PublishRideCompleted(ctx context.Context, event RideCompletedEvent) error {
	event.Event = "ride_completed"
	event.Timestamp = time.Now().Unix()
	return s.PublishEvent(ctx, "ride-events", event, event.RideID)
}

// Обёртка для отмены поездки
func (s *UserService) PublishRideCanceled(ctx context.Context, event RideCanceledEvent) error {
	event.Event = "ride_canceled"
	event.Timestamp = time.Now().Unix()
	return s.PublishEvent(ctx, "ride-events", event, event.RideID)
}
