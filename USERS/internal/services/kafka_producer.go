package services

import (
	"context"
	"encoding/json"
	"log"

	"github.com/segmentio/kafka-go"
)

type KafkaProducer struct {
	writer *kafka.Writer
}

func NewKafkaProducer(brokers []string, topic string) *KafkaProducer {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  brokers,
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	})
	return &KafkaProducer{writer: writer}
}

func (p *KafkaProducer) SendEvent(ctx context.Context, event interface{}) error {
	data, err := json.Marshal(event)
	if err != nil {
		log.Printf("Failed to marshal event: %v", err)
		return err
	}

	msg := kafka.Message{
		Value: data,
	}

	err = p.writer.WriteMessages(ctx, msg)
	if err != nil {
		log.Printf("Failed to write message to Kafka: %v", err)
	}
	return err
}

func (p *KafkaProducer) Close() error {
	return p.writer.Close()
}
