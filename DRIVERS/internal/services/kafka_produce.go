package services

import (
	"context"
	"encoding/json"
	"time"

	"github.com/segmentio/kafka-go"
)

func PublishRideAccepted(ctx context.Context, writer *kafka.Writer, event RideAcceptedEvent) error {
	event.Event = "ride_accepted"
	event.Timestamp = time.Now().Unix()
	data, err := json.Marshal(event)
	if err != nil {
		return err
	}
	msg := kafka.Message{
		Key:   []byte(event.RideID),
		Value: data,
		Topic: "ride-events",
	}
	return writer.WriteMessages(ctx, msg)
}

func PublishRideCompleted(ctx context.Context, writer *kafka.Writer, event RideCompletedEvent) error {
	event.Event = "ride_completed"
	event.Timestamp = time.Now().Unix()
	data, err := json.Marshal(event)
	if err != nil {
		return err
	}
	msg := kafka.Message{
		Key:   []byte(event.RideID),
		Value: data,
		Topic: "ride-events",
	}
	return writer.WriteMessages(ctx, msg)
}

func PublishRideCanceled(ctx context.Context, writer *kafka.Writer, event RideCanceledEvent) error {
	event.Event = "ride_canceled"
	event.Timestamp = time.Now().Unix()
	data, err := json.Marshal(event)
	if err != nil {
		return err
	}
	msg := kafka.Message{
		Key:   []byte(event.RideID),
		Value: data,
		Topic: "ride-events",
	}
	return writer.WriteMessages(ctx, msg)
}
