package services

import (
    "context"
    "encoding/json"
    "log"
    "time"

    pb "github.com/GameXost/YandexGo_proj/DRIVERS/API/generated/drivers"
    "github.com/segmentio/kafka-go"
)

func (s *DriverService) StartKafkaConsumer(ctx context.Context, reader *kafka.Reader) {
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
            case "get_driver_location":
                var req struct {
                    BaseEvent
                    RideID   string `json:"ride_id"`
                    DriverID string `json:"driver_id"`
                }
                if err := json.Unmarshal(m.Value, &req); err != nil {
                    log.Printf("kafka is unable to unmarshal get_driver_location event, err : %v", err)
                    continue
                }

                var lat, lon float64
                location, err := s.GetDriverLocation(ctx, req.DriverID)
                if err != nil {
                    log.Printf("Failed to get driver location: %v", err)
                } else {
                    lat = location.Latitude
                    lon = location.Longitude
                }

                if req.ReplyTo != "" && req.CorrelationID != "" {
                    response := struct {
                        BaseEvent
                        RideID   string  `json:"ride_id"`
                        DriverID string  `json:"driver_id"`
                        Lat      float64 `json:"lat"`
                        Lon      float64 `json:"lon"`
                    }{
                        BaseEvent: BaseEvent{
                            Event:         "driver_location_response",
                            CorrelationID: req.CorrelationID,
                            Timestamp:     time.Now().Unix(),
                        },
                        RideID:   req.RideID,
                        DriverID: req.DriverID,
                        Lat:      lat,
                        Lon:      lon,
                    }
                    err := s.sendKafkaResponse(ctx, req.ReplyTo, response)
                    if err != nil {
                        log.Printf("Failed to send Kafka response: %v", err)
                    }
                }
            case "get_driver_info":
                var req struct {
                    BaseEvent
                    DriverID string `json:"driver_id"`
                }
                if err := json.Unmarshal(m.Value, &req); err != nil {
                    log.Printf("kafka is unable to unmarshal get_driver_info event, err : %v", err)
                    continue
                }
                driver, err := s.GetDriverProfile(ctx, req.DriverID)
                if err != nil {
                    log.Printf("Failed to get driver profile: %v", err)
                    continue
                }
                if req.ReplyTo != "" && req.CorrelationID != "" {
                    response := struct {
                        BaseEvent
                        Driver *pb.Driver `json:"driver"`
                    }{
                        BaseEvent: BaseEvent{
                            Event:         "driver_info_response",
                            CorrelationID: req.CorrelationID,
                            Timestamp:     time.Now().Unix(),
                        },
                        Driver: driver,
                    }
                    err := s.sendKafkaResponse(ctx, req.ReplyTo, response)
                    if err != nil {
                        log.Printf("Failed to send Kafka response: %v", err)
                    }
                }

            }

        }
    }()
}

func (s *DriverService) sendKafkaResponse(ctx context.Context, topic string, event interface{}) error {
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
