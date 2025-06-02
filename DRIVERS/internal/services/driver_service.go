package services

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"sync"
	"time"

	"github.com/google/uuid"

	"github.com/GameXost/YandexGo_proj/DRIVERS/internal/config"
	"github.com/GameXost/YandexGo_proj/DRIVERS/internal/prometh"

	drivers_pb "github.com/GameXost/YandexGo_proj/DRIVERS/API/generated/drivers"
	"github.com/GameXost/YandexGo_proj/DRIVERS/internal/models"
	"github.com/GameXost/YandexGo_proj/DRIVERS/internal/repository"
	users_pb "github.com/GameXost/YandexGo_proj/USERS/API/generated/clients" // Corrected import based on new go_package
	"github.com/redis/go-redis/v9"
	"github.com/segmentio/kafka-go"
)

type RedisClient interface {
	Get(ctx context.Context, key string) *redis.StringCmd
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd
	Del(ctx context.Context, keys ...string) *redis.IntCmd
	Scan(ctx context.Context, cursor uint64, match string, count int64) *redis.ScanIterator
}

type EventWrapper struct {
	Event         string `json:"event"`
	CorrelationID string `json:"correlation_id,omitempty"`
	ReplyTo       string `json:"reply_to,omitempty"`
	RideID        string `json:"ride_id,omitempty"`
	DriverID      string `json:"driver_id,omitempty"`
	PassengerID   string `json:"passenger_id,omitempty"`
}

type DriverService struct {
	Repo                     repository.DriverRepositoryInterface
	RedisDrivers             RedisClient
	RedisRides               RedisClient
	kafkaWriter              *kafka.Writer
	kafkaBrokers             []string
	ridesTopic               string
	userRequestsTopic        string
	userNotificationsTopic   string
	driverRequestsTopic      string
	driverNotificationsTopic string
	responseChannels         *sync.Map
}

func NewDriverService(
	repo repository.DriverRepositoryInterface,
	redisDrivers RedisClient,
	redisRides RedisClient,
	kafkaWriter *kafka.Writer,
	cfg *config.Config,
) *DriverService {
	return &DriverService{
		Repo:                     repo,
		RedisDrivers:             redisDrivers,
		RedisRides:               redisRides,
		kafkaWriter:              kafkaWriter,
		kafkaBrokers:             cfg.Kafka.Brokers,
		ridesTopic:               cfg.Kafka.Topics.Rides,
		userRequestsTopic:        cfg.Kafka.Topics.UserRequests,
		userNotificationsTopic:   cfg.Kafka.Topics.UserNotifications,
		driverRequestsTopic:      cfg.Kafka.Topics.DriverRequests,
		driverNotificationsTopic: cfg.Kafka.Topics.DriverNotifications,
		responseChannels:         &sync.Map{},
	}
}

func (s *DriverService) sendKafkaMessage(ctx context.Context, topic string, eventData interface{}) error {
	payload, err := json.Marshal(eventData)
	if err != nil {
		return fmt.Errorf("failed to marshal event data for topic %s: %w", topic, err)
	}

	msg := kafka.Message{
		Topic: topic,
		Value: payload,
		Time:  time.Now(),
	}

	if err := s.kafkaWriter.WriteMessages(ctx, msg); err != nil {
		prometh.KafkaProduceErrors.Inc()
		return fmt.Errorf("failed to write message to Kafka topic %s: %w", topic, err)
	}

	log.Printf("Successfully sent message to Kafka topic %s: %s", topic, string(payload))
	prometh.KafkaProducedMessages.WithLabelValues(topic).Inc()
	return nil
}

func (s *DriverService) GetDriverProfile(ctx context.Context, driverID string) (*drivers_pb.Driver, error) {
	driver, err := s.Repo.GetDriverByID(ctx, driverID)
	if err != nil {
		return nil, fmt.Errorf("driver not found: %w", err)
	}
	return modelToProtoDriver(driver), nil
}

func (s *DriverService) UpdateDriverProfile(ctx context.Context, req *drivers_pb.Driver) (*drivers_pb.Driver, error) {
	err := s.Repo.UpdateDriverProfile(ctx, &models.Driver{
		ID:         req.Id,
		UserName:   req.Username,
		Email:      req.Email,
		Phone:      req.Phone,
		Car_color:  req.CarColor,
		Car_number: req.CarNumber,
		Car_model:  req.CarModel,
		Car_marks:  req.CarMark,
	})
	if err != nil {
		return nil, err
	}
	return modelToProtoDriver(&models.Driver{
		ID:         req.Id,
		UserName:   req.Username,
		Email:      req.Email,
		Phone:      req.Phone,
		Car_color:  req.CarColor,
		Car_number: req.CarNumber,
		Car_model:  req.CarModel,
		Car_marks:  req.CarMark,
	}), nil
}

func (s *DriverService) GetPassengerInfo(ctx context.Context, userID string) (*users_pb.User, error) {
	correlationID := uuid.New().String()
	replyTo := s.driverNotificationsTopic
	event := GetPassengerInfoRequest{
		BaseEvent:   NewBaseEvent("GetPasssengerInfoRequest", correlationID, replyTo),
		PassengerID: userID,
	}

	respChan := make(chan interface{}, 1)
	s.responseChannels.Store(correlationID, respChan) // Changed from userID to correlationID

	timeoutCtx, cancel := context.WithTimeout(ctx, time.Second*5) // Reduced timeout
	defer cancel()

	if err := s.sendKafkaMessage(timeoutCtx, s.userRequestsTopic, event); err != nil {
		s.responseChannels.Delete(correlationID)
		prometh.KafkaProduceErrors.Inc()
		return nil, fmt.Errorf("failed to send GetPassengerInfoRequest: %w", err)
	}
	select {
	case <-timeoutCtx.Done():
		s.responseChannels.Delete(correlationID)
		prometh.KafkaRequestTimeouts.WithLabelValues("UserProfileResponse").Inc()
		return nil, fmt.Errorf("timeout waiting for UserProfileResponse for PassengerID %s (CorrelationID: %s)", userID, correlationID)
	case resp := <-respChan:
		profileResp, ok := resp.(UserProfileResponse)
		if !ok {
			return nil, fmt.Errorf("received unexpected response type for PassengerID %s (CorrelationID: %s)", userID, correlationID)
		}
		if profileResp.Error != "" {
			return nil, fmt.Errorf("error in UserProfileResponse for PassengerID %s: %s", userID, profileResp.Error)
		}
		return profileResp.User, nil
	}
}

func (s *DriverService) AcceptRide(ctx context.Context, rideID string, driverID string) (*drivers_pb.StatusResponse, error) {
	if s.RedisRides == nil {
		return &drivers_pb.StatusResponse{Status: false, Message: "Redis unavailable"}, nil
	}

	rideKey := "ride:" + rideID
	rideData, err := s.RedisRides.Get(ctx, rideKey).Result()
	if err == redis.Nil {
		return &drivers_pb.StatusResponse{Status: false, Message: "Ride not found"}, nil
	} else if err != nil {
		return &drivers_pb.StatusResponse{Status: false, Message: "Redis error"}, nil
	}

	var ride drivers_pb.Ride
	if err := json.Unmarshal([]byte(rideData), &ride); err != nil {
		return &drivers_pb.StatusResponse{Status: false, Message: "Invalid ride data"}, nil
	}
	if ride.Status != "pending" {
		return &drivers_pb.StatusResponse{Status: false, Message: "Ride already accepted or completed"}, nil
	}
	ride.Status = "accepted"
	ride.DriverId = driverID

	updatedData, err := json.Marshal(ride)
	if err != nil {
		return &drivers_pb.StatusResponse{Status: false, Message: "Failed to serialize ride"}, nil
	}
	if err := s.RedisRides.Set(ctx, rideKey, updatedData, time.Hour).Err(); err != nil {
		return &drivers_pb.StatusResponse{Status: false, Message: "Failed to update ride in Redis"}, nil
	}

	driverKey := "driver:" + driverID + ":current_ride"
	if err := s.RedisRides.Set(ctx, driverKey, rideID, time.Hour).Err(); err != nil {
		return &drivers_pb.StatusResponse{Status: false, Message: "Failed to set current ride for driver"}, nil
	}

	acceptedEvent := RideAcceptedEvent{
		BaseEvent:     NewBaseEvent("RideAcceptedEvent", uuid.New().String(), ""),
		RideID:        rideID,
		PassengerID:   ride.UserId,
		DriverID:      driverID,
		StartLocation: fmt.Sprintf("%f,%f", ride.StartLocation.Latitude, ride.StartLocation.Longitude),
		EndLocation:   fmt.Sprintf("%f,%f", ride.EndLocation.Latitude, ride.EndLocation.Longitude),
		Status:        "accepted",
	}
	if err := s.sendKafkaMessage(ctx, s.ridesTopic, acceptedEvent); err != nil {
		log.Printf("Error publishing RideAcceptedEvent: %v", err)
		return &drivers_pb.StatusResponse{Status: false, Message: "Failed to publish ride accepted event"}, err
	}

	prometh.RideAcceptedCounter.Inc()
	return &drivers_pb.StatusResponse{
		Status:  true,
		Message: "Ride accepted successfully",
	}, nil
}

func modelToProtoDriver(m *models.Driver) *drivers_pb.Driver {
	return &drivers_pb.Driver{
		Id:        m.ID,
		Username:  m.UserName,
		Email:     m.Email,
		Phone:     m.Phone,
		CarNumber: m.Car_number,
		CarModel:  m.Car_model,
		CarMark:   m.Car_marks,
		CarColor:  m.Car_color,
	}
}

func (s *DriverService) GetCurrentRide(ctx context.Context, driverID string) (*drivers_pb.Ride, error) {
	if s.RedisRides == nil {
		return nil, fmt.Errorf("Redis unavailable")
	}
	driverKey := "driver:" + driverID + ":current_ride"
	rideID, err := s.RedisRides.Get(ctx, driverKey).Result()
	if err == redis.Nil {
		return nil, fmt.Errorf("current ride not found for driver")
	} else if err != nil {
		return nil, fmt.Errorf("Redis error: %w", err)
	}
	rideKey := "ride:" + rideID
	rideData, err := s.RedisRides.Get(ctx, rideKey).Result()
	if err == redis.Nil {
		return nil, fmt.Errorf("ride not found in Redis")
	} else if err != nil {
		return nil, fmt.Errorf("Redis error: %w", err)
	}
	var ride drivers_pb.Ride
	if err := json.Unmarshal([]byte(rideData), &ride); err != nil {
		return nil, fmt.Errorf("invalid ride data: %w", err)
	}
	return &ride, nil
}

type GetUserProfileEvent struct {
	BaseEvent
	UserID string `json:"user_id"`
}

func (s *DriverService) CompleteRide(ctx context.Context, rideID string, driverID string) (*drivers_pb.StatusResponse, error) {
	if s.RedisRides == nil {
		return &drivers_pb.StatusResponse{Status: false, Message: "Redis unavailable"}, nil
	}

	rideKey := "ride:" + rideID
	rideData, err := s.RedisRides.Get(ctx, rideKey).Result()
	if err == redis.Nil {
		return &drivers_pb.StatusResponse{Status: false, Message: "Ride not found"}, nil
	} else if err != nil {
		return &drivers_pb.StatusResponse{Status: false, Message: "Redis error"}, nil
	}

	var ride drivers_pb.Ride
	if err := json.Unmarshal([]byte(rideData), &ride); err != nil {
		return &drivers_pb.StatusResponse{Status: false, Message: "Invalid ride data"}, nil
	}

	if ride.DriverId != driverID {
		return &drivers_pb.StatusResponse{Status: false, Message: "Ride not assigned to this driver"}, nil
	}
	if ride.Status != "accepted" && ride.Status != "in_progress" {
		return &drivers_pb.StatusResponse{Status: false, Message: "Ride not in progress"}, nil
	}

	ride.Status = "completed"
	updatedData, err := json.Marshal(ride)
	if err != nil {
		return &drivers_pb.StatusResponse{Status: false, Message: "Failed to serialize ride"}, nil
	}
	if err := s.RedisRides.Set(ctx, rideKey, updatedData, time.Hour).Err(); err != nil {
		return &drivers_pb.StatusResponse{Status: false, Message: "Failed to update ride in Redis"}, nil
	}

	driverKey := "driver:" + driverID + ":current_ride"
	_ = s.RedisRides.Del(ctx, driverKey).Err()

	completedEvent := RideCompletedEvent{
		BaseEvent: NewBaseEvent("RideCompletedEvent", uuid.New().String(), ""),
		RideID:    rideID,
		DriverID:  driverID,
	}
	if err := s.sendKafkaMessage(ctx, s.ridesTopic, completedEvent); err != nil {
		log.Printf("Error publishing RideCompletedEvent: %v", err)
		return &drivers_pb.StatusResponse{Status: false, Message: "Failed to publish ride completed event"}, err
	}
	prometh.RideCompletedCounter.Inc()

	return &drivers_pb.StatusResponse{
		Status:  true,
		Message: "Ride completed successfully",
	}, nil
}

func (s *DriverService) CancelRide(ctx context.Context, rideID string, driverID string) (*drivers_pb.StatusResponse, error) {
	if s.RedisRides == nil {
		return &drivers_pb.StatusResponse{Status: false, Message: "Redis unavailable"}, nil
	}

	rideKey := "ride:" + rideID
	rideData, err := s.RedisRides.Get(ctx, rideKey).Result()
	if err == redis.Nil {
		return &drivers_pb.StatusResponse{Status: false, Message: "Ride not found"}, nil
	} else if err != nil {
		return &drivers_pb.StatusResponse{Status: false, Message: "Redis error"}, nil
	}

	var ride drivers_pb.Ride
	if err := json.Unmarshal([]byte(rideData), &ride); err != nil {
		return &drivers_pb.StatusResponse{Status: false, Message: "Invalid ride data"}, nil
	}

	if ride.DriverId != driverID {
		return &drivers_pb.StatusResponse{Status: false, Message: "Ride not assigned to this driver"}, nil
	}
	if ride.Status != "accepted" && ride.Status != "in_progress" {
		return &drivers_pb.StatusResponse{Status: false, Message: "Ride not in progress"}, nil
	}

	ride.Status = "canceled"
	updatedData, err := json.Marshal(ride)
	if err != nil {
		return &drivers_pb.StatusResponse{Status: false, Message: "Failed to serialize ride"}, nil
	}
	if err := s.RedisRides.Set(ctx, rideKey, updatedData, time.Hour).Err(); err != nil {
		return &drivers_pb.StatusResponse{Status: false, Message: "Failed to update ride in Redis"}, nil
	}

	driverKey := "driver:" + driverID + ":current_ride"
	_ = s.RedisRides.Del(ctx, driverKey).Err()

	canceledEvent := RideCanceledEvent{
		BaseEvent: NewBaseEvent("RideCanceledEvent", uuid.New().String(), ""),
		RideID:    rideID,
		DriverID:  driverID,
		Reason:    "driver_cancelled",
	}
	if err := s.sendKafkaMessage(ctx, s.ridesTopic, canceledEvent); err != nil {
		log.Printf("Error publishing RideCanceledEvent: %v", err)
		return &drivers_pb.StatusResponse{Status: false, Message: "Failed to publish ride canceled event"}, err
	}
	prometh.RideCanceledCounter.Inc()

	return &drivers_pb.StatusResponse{
		Status:  true,
		Message: "Ride canceled successfully",
	}, nil
}

func (s *DriverService) GetRideHistory(ctx context.Context, driverID string) (*drivers_pb.RideHistoryResponse, error) {
	if s.RedisRides == nil {
		return nil, fmt.Errorf("Redis unavailable")
	}

	var rides []*drivers_pb.Ride

	iter := s.RedisRides.(*redis.Client).Scan(ctx, 0, "ride:*", 0).Iterator()
	for iter.Next(ctx) {
		rideKey := iter.Val()
		rideData, err := s.RedisRides.Get(ctx, rideKey).Result()
		if err != nil {
			continue
		}
		var ride drivers_pb.Ride
		if err := json.Unmarshal([]byte(rideData), &ride); err != nil {
			continue
		}
		if ride.DriverId == driverID {
			rides = append(rides, &ride)
		}
	}
	if err := iter.Err(); err != nil {
		return nil, fmt.Errorf("error scanning rides: %w", err)
	}

	return &drivers_pb.RideHistoryResponse{Rides: rides}, nil
}

func haversine(lat1, lon1, lat2, lon2 float64) float64 {
	const R = 6371
	dLat := (lat2 - lat1) * math.Pi / 180.0
	dLon := (lon2 - lon1) * math.Pi / 180.0
	lat1 = lat1 * math.Pi / 180.0
	lat2 = lat2 * math.Pi / 180.0

	a := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Sin(dLon/2)*math.Sin(dLon/2)*math.Cos(lat1)*math.Cos(lat2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return R * c
}

func (s *DriverService) GetNearbyRequests(ctx context.Context, loc *drivers_pb.Location) (*drivers_pb.RideRequestsResponse, error) {
	if s.RedisRides == nil {
		return nil, fmt.Errorf("Redis unavailable")
	}
	const radiusKm = 3.0

	var requests []*drivers_pb.RideRequest

	iter := s.RedisRides.(*redis.Client).Scan(ctx, 0, "ride:*", 0).Iterator()
	for iter.Next(ctx) {
		rideKey := iter.Val()
		rideData, err := s.RedisRides.Get(ctx, rideKey).Result()
		if err != nil {
			continue
		}
		var ride drivers_pb.Ride
		if err := json.Unmarshal([]byte(rideData), &ride); err != nil {
			continue
		}
		if ride.Status != "pending" {
			continue
		}
		dist := haversine(loc.Latitude, loc.Longitude, ride.StartLocation.Latitude, ride.StartLocation.Longitude)
		if dist <= radiusKm {
			requests = append(requests, &drivers_pb.RideRequest{
				UserId:        ride.UserId,
				StartLocation: ride.StartLocation,
				EndLocation:   ride.EndLocation,
			})
		}
	}
	if err := iter.Err(); err != nil {
		return nil, fmt.Errorf("error scanning rides: %w", err)
	}

	return &drivers_pb.RideRequestsResponse{RideRequests: requests}, nil
}

func (s *DriverService) UpdateLocation(ctx context.Context, update *drivers_pb.LocationUpdateRequest) (*drivers_pb.StatusResponse, error) {
	if s.RedisDrivers == nil {
		return &drivers_pb.StatusResponse{Status: false, Message: "Redis unavailable"}, nil
	}
	key := "driver_location:" + update.DriverId
	locData, err := json.Marshal(update.Location)
	if err != nil {
		return &drivers_pb.StatusResponse{Status: false, Message: "Invalid location data"}, fmt.Errorf("failed to marshal location: %w", err)
	}
	err = s.RedisDrivers.Set(ctx, key, locData, time.Second*30).Err() // Added 30 second expiry
	if err != nil {
		return &drivers_pb.StatusResponse{Status: false, Message: "Failed to save location"}, err
	}
	return &drivers_pb.StatusResponse{
		Status:  true,
		Message: "Location updated successfully",
	}, nil
}

func (s *DriverService) GetDriverLocation(ctx context.Context, driverID string) (*drivers_pb.Location, error) {
	if s.RedisDrivers == nil {
		return nil, fmt.Errorf("Redis unavailable")
	}
	key := "driver_location:" + driverID
	locData, err := s.RedisDrivers.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	var location drivers_pb.Location
	if err := json.Unmarshal([]byte(locData), &location); err != nil {
		return nil, err
	}
	return &location, nil
}

func (s *DriverService) HandleUserProfileResponse(ctx context.Context, response UserProfileResponse) {
	log.Printf("DRIVERS: Received UserProfileResponse for CorrelationID: %s. Error: %s", response.CorrelationID, response.Error)

	if ch, ok := s.responseChannels.Load(response.CorrelationID); ok {
		ch.(chan interface{}) <- response
		s.responseChannels.Delete(response.CorrelationID)
		log.Printf("DRIVERS: Dispatched UserProfileResponse for CorrelationID: %s", response.CorrelationID)
	} else {
		log.Printf("DRIVERS: No waiting channel found for CorrelationID: %s. Response might be unsolicited or timed out.", response.CorrelationID)
	}
}

func (s *DriverService) HandleGetDriverLocationEvent(ctx context.Context, event GetDriverLocationEvent) {
	log.Printf("DRIVERS: Received GetDriverLocationEvent for DriverID: %s, RideID: %s", event.DriverID, event.RideID)

	location, err := s.GetDriverLocation(ctx, event.DriverID)
	var errStr string
	var lat, lon float64
	if err != nil {
		errStr = fmt.Sprintf("Failed to get driver location: %v", err)
		log.Printf("DRIVERS: %s", errStr)
	} else {
		lat = location.Latitude
		lon = location.Longitude
	}

	response := DriverLocationResponse{
		BaseEvent: NewBaseEvent("DriverLocationResponse", event.CorrelationID, ""),
		DriverID:  event.DriverID,
		Latitude:  lat,
		Longitude: lon,
		Error:     errStr,
	}

	if event.ReplyTo == "" {
		log.Printf("Warning: GetDriverLocationEvent has no ReplyTo topic. Cannot send response for DriverID: %s", event.DriverID)
		return
	}

	if err := s.sendKafkaMessage(ctx, event.ReplyTo, response); err != nil {
		log.Printf("Error sending DriverLocationResponse to topic %s: %v", event.ReplyTo, err)
	} else {
		log.Printf("Sent DriverLocationResponse for DriverID: %s to topic %s", response.DriverID, event.ReplyTo)
	}
}

func (s *DriverService) HandleGetDriverInfoEvent(ctx context.Context, event GetDriverInfoEvent) {
	log.Printf("DRIVERS: Received GetDriverInfoEvent for DriverID: %s", event.DriverID)

	driver, err := s.GetDriverProfile(ctx, event.DriverID)
	var errStr string
	if err != nil {
		errStr = fmt.Sprintf("Failed to get driver profile: %v", err)
		log.Printf("DRIVERS: %s", errStr)
	}

	type DriverInfoResponse struct {
		BaseEvent
		Driver *drivers_pb.Driver `json:"driver,omitempty"`
		Error  string             `json:"error,omitempty"`
	}

	response := DriverInfoResponse{
		BaseEvent: NewBaseEvent("DriverInfoResponse", event.CorrelationID, ""),
		Driver:    driver,
		Error:     errStr,
	}

	if event.ReplyTo == "" {
		log.Printf("Warning: GetDriverInfoEvent has no ReplyTo topic. Cannot send response for DriverID: %s", event.DriverID)
		return
	}

	if err := s.sendKafkaMessage(ctx, event.ReplyTo, response); err != nil {
		log.Printf("Error sending DriverInfoResponse to topic %s: %v", event.ReplyTo, err)
	} else {
		log.Printf("Sent DriverInfoResponse for DriverID: %s to topic %s", response.Driver.Id, event.ReplyTo)
	}
}
