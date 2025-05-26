package services

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

//start kafka consumer

func (s *DriverService) StartKafkaConsumer(ctx context.Context, reader *kafka.Reader) {
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
				log.Printf("kakfa err : %v", err)
				continue
			}
			switch baseEvent.Event {
			case "ride_created":
				var event RideCreatedEvent
				if err := json.Unmarshal(m.Value, &event); err != nil {
					log.Printf("kafka is unable to unmarshal ride created event, err : %v", err)
					continue
				}
				//ride to redis
				rideKey := "ride:" + event.RideID
				rideData, _ := json.Marshal(event)
				if s.RedisRides != nil {
					_ = s.RedisRides.Set(ctx, rideKey, rideData, time.Hour*2).Err()
				}
				log.Printf("Ride created: ride_id=%s, passenger_id=%s", event.RideID, event.PassengerID)
			case "ride_canceled":
				var event RideCanceledEvent
				if err := json.Unmarshal(m.Value, &event); err != nil {
					log.Printf("kafka is unable to unmarshal ride canceled event, err : %v", err)
					continue
				}
				rideKey := "ride:" + event.RideID
				if s.RedisRides != nil {
					_ = s.RedisRides.Del(ctx, rideKey).Err()
				}
				log.Printf("Ride canceled: ride_id=%s, driver_id=%s, reason=%s", event.RideID, event.DriverID, event.Reason)
			case "ride_completed":
				var event RideCompletedEvent
				if err := json.Unmarshal(m.Value, &event); err != nil {
					log.Printf("kafka is unable to unmarshal ride completed event, err : %v", err)
					continue
				}
				rideKey := "ride:" + event.RideID
				if s.RedisRides != nil {
					_ = s.RedisRides.Del(ctx, rideKey).Err()
				}
				log.Printf("Ride completed: ride_id=%s, driver_id=%s", event.RideID, event.DriverID)
			}
		}
	}()
}
