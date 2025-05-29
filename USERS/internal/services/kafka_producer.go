package services

import (
	"context"
	"encoding/json"
	"time"

	"github.com/segmentio/kafka-go"
)

// Универсальная функция для отправки любого события
func PublishEvent(ctx context.Context, writer *kafka.Writer, topic string, event interface{}, key string) error {
	data, err := json.Marshal(event)
	if err != nil {
		return err
	}
	msg := kafka.Message{
		Key:   []byte(key),
		Value: data,
		Topic: topic,
	}
	return writer.WriteMessages(ctx, msg)
}

// Пример thin-обёртки для создания поездки
func PublishRideCreated(ctx context.Context, writer *kafka.Writer, event RideCreatedEvent) error {
	event.Event = "ride_created"
	event.Timestamp = time.Now().Unix()
	return PublishEvent(ctx, writer, "ride-events", event, event.RideID)
}

// Аналогично для других событий, если потребуется
func PublishRideAccepted(ctx context.Context, writer *kafka.Writer, event RideAcceptedEvent) error {
	event.Event = "ride_accepted"
	event.Timestamp = time.Now().Unix()
	return PublishEvent(ctx, writer, "ride-events", event, event.RideID)
}

func PublishRideCanceled(ctx context.Context, writer *kafka.Writer, event RideCanceledEvent) error {
	event.Event = "ride_canceled"
	event.Timestamp = time.Now().Unix()
	return PublishEvent(ctx, writer, "ride-events", event, event.RideID)
}
