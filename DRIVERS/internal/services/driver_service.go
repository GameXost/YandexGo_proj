package services

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/GameXost/YandexGo_proj/DRIVERS/internal/metrics"
	"math"
	"time"

	pb "github.com/GameXost/YandexGo_proj/DRIVERS/API/generated/drivers"
	"github.com/GameXost/YandexGo_proj/DRIVERS/internal/models"
	"github.com/GameXost/YandexGo_proj/DRIVERS/internal/repository"
	"github.com/redis/go-redis/v9"
	"github.com/segmentio/kafka-go"
)

// хунйя ля работы редиса вроде чет робит
type RedisClient interface {
	Get(ctx context.Context, key string) *redis.StringCmd
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd
	Del(ctx context.Context, keys ...string) *redis.IntCmd
}

type DriverService struct {
	Repo         repository.DriverRepositoryInterface
	RedisDrivers RedisClient // For online drivers
	RedisRides   RedisClient // For rides
	Kafka        *kafka.Writer
}

func NewDriverService(repo *repository.DriverRepository, redisDrivers *redis.Client, redisRides *redis.Client, kafka *kafka.Writer) *DriverService {
	return &DriverService{
		Repo:         repo,
		RedisDrivers: redisDrivers,
		RedisRides:   redisRides,
		Kafka:        kafka,
	}
}

func (s *DriverService) GetDriverProfile(ctx context.Context, driverID string) (*pb.Driver, error) {
	cacheKey := "driver_profile:" + driverID

	if s.RedisRides != nil {
		cached, err := s.RedisRides.Get(ctx, cacheKey).Result()
		if err == nil && cached != "" {
			var driver pb.Driver
			if err := json.Unmarshal([]byte(cached), &driver); err == nil {
				return &driver, nil
			}
		}
	}

	driver, err := s.Repo.GetDriverByID(ctx, driverID)
	if err != nil {
		return nil, fmt.Errorf("driver not found: %w", err)
	}

	if s.RedisRides != nil {
		if data, err := json.Marshal(driver); err == nil {
			_ = s.RedisRides.Set(ctx, cacheKey, data, time.Hour).Err()
		}
	}

	if s.Kafka != nil {
		msg := kafka.Message{
			Key:   []byte(driverID),
			Value: []byte(fmt.Sprintf("profile viewed: %s", driverID)),
		}
		_ = s.Kafka.WriteMessages(ctx, msg)
	}

	return modelToProtoDriver(driver), nil
}

func (s *DriverService) UpdateDriverProfile(ctx context.Context, req *pb.Driver) (*pb.Driver, error) {
	// req.Id уже гарантированно driverID
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
	cacheKey := "driver_profile:" + req.Id
	if s.RedisRides != nil {
		_ = s.RedisRides.Del(ctx, cacheKey).Err()
	}
	if s.Kafka != nil {
		msg := kafka.Message{
			Key:   []byte(req.Id),
			Value: []byte("profile updated: " + req.Id),
		}
		_ = s.Kafka.WriteMessages(ctx, msg)
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

func (s *DriverService) AcceptRide(ctx context.Context, rideID string, driverID string) (*pb.StatusResponse, error) {
	if s.RedisRides == nil {
		return &pb.StatusResponse{Status: false, Message: "Redis unavailable"}, nil
	}

	rideKey := "ride:" + rideID
	rideData, err := s.RedisRides.Get(ctx, rideKey).Result()
	if err == redis.Nil {
		return &pb.StatusResponse{Status: false, Message: "Ride not found"}, nil
	} else if err != nil {
		return &pb.StatusResponse{Status: false, Message: "Redis error"}, nil
	}

	var ride pb.Ride
	if err := json.Unmarshal([]byte(rideData), &ride); err != nil {
		return &pb.StatusResponse{Status: false, Message: "Invalid ride data"}, nil
	}
	if ride.Status != "pending" {
		return &pb.StatusResponse{Status: false, Message: "Ride already accepted or completed"}, nil
	}
	ride.Status = "accepted"
	ride.DriverId = driverID

	updatedData, err := json.Marshal(ride)
	if err != nil {
		return &pb.StatusResponse{Status: false, Message: "Failed to serialize ride"}, nil
	}
	if err := s.RedisRides.Set(ctx, rideKey, updatedData, time.Hour).Err(); err != nil {
		return &pb.StatusResponse{Status: false, Message: "Failed to update ride in Redis"}, nil
	}

	driverKey := "driver:" + driverID + ":current_ride"
	if err := s.RedisRides.Set(ctx, driverKey, rideID, time.Hour).Err(); err != nil {
		return &pb.StatusResponse{Status: false, Message: "Failed to set current ride for driver"}, nil
	}

	// Публикация события в Kafka
	if s.Kafka != nil {
		event := RideAcceptedEvent{
			Event:         "ride_accepted",
			RideID:        rideID,
			PassengerID:   ride.UserId,
			DriverID:      driverID,
			StartLocation: fmt.Sprintf("%f,%f", ride.StartLocation.Latitude, ride.StartLocation.Longitude),
			EndLocation:   fmt.Sprintf("%f,%f", ride.EndLocation.Latitude, ride.EndLocation.Longitude),
			Timestamp:     time.Now().Unix(),
			Status:        "accepted",
		}
		_ = PublishRideAccepted(ctx, s.Kafka, event)
	}
	metrics.RideAcceptedCounter.Inc()
	return &pb.StatusResponse{
		Status:  true,
		Message: "Ride accepted successfully",
	}, nil
}

// хуйня чтоб из протоформатика в модельку для норм работы
func modelToProtoDriver(m *models.Driver) *pb.Driver {
	return &pb.Driver{
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

func (s *DriverService) GetCurrentRide(ctx context.Context, driverID string) (*pb.Ride, error) {
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
	var ride pb.Ride
	if err := json.Unmarshal([]byte(rideData), &ride); err != nil {
		return nil, fmt.Errorf("invalid ride data: %w", err)
	}
	return &ride, nil
}

func (s *DriverService) GetPassengerInfo(ctx context.Context, userID string) (*pb.User, error) {
	if s.Repo == nil {
		return nil, fmt.Errorf("repository unavailable")
	}
	user, err := s.Repo.GetUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return &pb.User{
		Id:       user.ID,
		Username: user.Username,
		Phone:    user.Phone,
	}, nil
}

func (s *DriverService) CompleteRide(ctx context.Context, rideID string, driverID string) (*pb.StatusResponse, error) {
	if s.RedisRides == nil {
		return &pb.StatusResponse{Status: false, Message: "Redis unavailable"}, nil
	}

	// 1. Get ride from Redis
	rideKey := "ride:" + rideID
	rideData, err := s.RedisRides.Get(ctx, rideKey).Result()
	if err == redis.Nil {
		return &pb.StatusResponse{Status: false, Message: "Ride not found"}, nil
	} else if err != nil {
		return &pb.StatusResponse{Status: false, Message: "Redis error"}, nil
	}

	var ride pb.Ride
	if err := json.Unmarshal([]byte(rideData), &ride); err != nil {
		return &pb.StatusResponse{Status: false, Message: "Invalid ride data"}, nil
	}

	// 2. Check if the ride is assigned to this driver and is in progress/accepted
	if ride.DriverId != driverID {
		return &pb.StatusResponse{Status: false, Message: "Ride not assigned to this driver"}, nil
	}
	if ride.Status != "accepted" && ride.Status != "in_progress" {
		return &pb.StatusResponse{Status: false, Message: "Ride not in progress"}, nil
	}

	// 3. Update status to completed
	ride.Status = "completed"
	updatedData, err := json.Marshal(ride)
	if err != nil {
		return &pb.StatusResponse{Status: false, Message: "Failed to serialize ride"}, nil
	}
	if err := s.RedisRides.Set(ctx, rideKey, updatedData, time.Hour).Err(); err != nil {
		return &pb.StatusResponse{Status: false, Message: "Failed to update ride in Redis"}, nil
	}

	// 4. Remove current ride from driver
	driverKey := "driver:" + driverID + ":current_ride"
	_ = s.RedisRides.Del(ctx, driverKey).Err()

	// 5. Notify via Kafka (структурированное событие)
	if s.Kafka != nil {
		event := RideCompletedEvent{
			Event:     "ride_completed",
			RideID:    rideID,
			DriverID:  driverID,
			Timestamp: time.Now().Unix(),
			// Можно добавить duration, стоимость, координаты завершения и т.д.
		}
		_ = PublishRideCompleted(ctx, s.Kafka, event)
	}
	metrics.RideCompletedCounter.Inc()

	return &pb.StatusResponse{
		Status:  true,
		Message: "Ride completed successfully",
	}, nil
}

func (s *DriverService) CancelRide(ctx context.Context, rideID string, driverID string) (*pb.StatusResponse, error) {
	if s.RedisRides == nil {
		return &pb.StatusResponse{Status: false, Message: "Redis unavailable"}, nil
	}

	// 1. Get ride from Redis
	rideKey := "ride:" + rideID
	rideData, err := s.RedisRides.Get(ctx, rideKey).Result()
	if err == redis.Nil {
		return &pb.StatusResponse{Status: false, Message: "Ride not found"}, nil
	} else if err != nil {
		return &pb.StatusResponse{Status: false, Message: "Redis error"}, nil
	}

	var ride pb.Ride
	if err := json.Unmarshal([]byte(rideData), &ride); err != nil {
		return &pb.StatusResponse{Status: false, Message: "Invalid ride data"}, nil
	}

	// 2. Check if the ride is assigned to this driver and is in progress/accepted
	if ride.DriverId != driverID {
		return &pb.StatusResponse{Status: false, Message: "Ride not assigned to this driver"}, nil
	}
	if ride.Status != "accepted" && ride.Status != "in_progress" {
		return &pb.StatusResponse{Status: false, Message: "Ride not in progress"}, nil
	}

	// 3. Update status to canceled
	ride.Status = "canceled"
	updatedData, err := json.Marshal(ride)
	if err != nil {
		return &pb.StatusResponse{Status: false, Message: "Failed to serialize ride"}, nil
	}
	if err := s.RedisRides.Set(ctx, rideKey, updatedData, time.Hour).Err(); err != nil {
		return &pb.StatusResponse{Status: false, Message: "Failed to update ride in Redis"}, nil
	}

	// 4. Remove current ride from driver
	driverKey := "driver:" + driverID + ":current_ride"
	_ = s.RedisRides.Del(ctx, driverKey).Err()

	// 5. Notify via Kafka (структурированное событие)
	if s.Kafka != nil {
		event := RideCanceledEvent{
			Event:     "ride_canceled",
			RideID:    rideID,
			DriverID:  driverID,
			Reason:    "driver_cancelled",
			Timestamp: time.Now().Unix(),
		}
		_ = PublishRideCanceled(ctx, s.Kafka, event)
	}
	metrics.RideCanceledCounter.Inc()
	return &pb.StatusResponse{
		Status:  true,
		Message: "Ride canceled successfully",
	}, nil
}

// мб это у водятела вообще не нужно, надо фиксить мбмб
func (s *DriverService) GetRideHistory(ctx context.Context, driverID string) (*pb.RideHistoryResponse, error) {
	if s.RedisRides == nil {
		return nil, fmt.Errorf("Redis unavailable")
	}

	var rides []*pb.Ride

	// Scan all rides in RedisRides
	iter := s.RedisRides.(*redis.Client).Scan(ctx, 0, "ride:*", 0).Iterator()
	for iter.Next(ctx) {
		rideKey := iter.Val()
		rideData, err := s.RedisRides.Get(ctx, rideKey).Result()
		if err != nil {
			continue // skip if not found or error
		}
		var ride pb.Ride
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

	return &pb.RideHistoryResponse{Rides: rides}, nil
}

// Helper function to calculate distance between two coordinates (Haversine formula)
func haversine(lat1, lon1, lat2, lon2 float64) float64 {
	const R = 6371 // Earth radius in km LOL
	dLat := (lat2 - lat1) * math.Pi / 180.0
	dLon := (lon2 - lon1) * math.Pi / 180.0
	lat1 = lat1 * math.Pi / 180.0
	lat2 = lat2 * math.Pi / 180.0

	a := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Sin(dLon/2)*math.Sin(dLon/2)*math.Cos(lat1)*math.Cos(lat2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return R * c
}

func (s *DriverService) GetNearbyRequests(ctx context.Context, loc *pb.Location) (*pb.RideRequestsResponse, error) {
	if s.RedisRides == nil {
		return nil, fmt.Errorf("Redis unavailable")
	}
	// const variable for pending requests, maximum availability in radius km
	const radiusKm = 3.0

	var requests []*pb.RideRequest

	iter := s.RedisRides.(*redis.Client).Scan(ctx, 0, "ride:*", 0).Iterator()
	for iter.Next(ctx) {
		rideKey := iter.Val()
		rideData, err := s.RedisRides.Get(ctx, rideKey).Result()
		if err != nil {
			continue
		}
		var ride pb.Ride
		if err := json.Unmarshal([]byte(rideData), &ride); err != nil {
			continue
		}
		if ride.Status != "pending" {
			continue
		}
		dist := haversine(loc.Latitude, loc.Longitude, ride.StartLocation.Latitude, ride.StartLocation.Longitude)
		if dist <= radiusKm {
			requests = append(requests, &pb.RideRequest{
				UserId:        ride.UserId,
				StartLocation: ride.StartLocation,
				EndLocation:   ride.EndLocation,
			})
		}
	}
	if err := iter.Err(); err != nil {
		return nil, fmt.Errorf("error scanning rides: %w", err)
	}

	return &pb.RideRequestsResponse{RideRequests: requests}, nil
}

func (s *DriverService) UpdateLocation(ctx context.Context, updates <-chan *pb.LocationUpdateRequest) (*pb.StatusResponse, error) {
	if s.RedisDrivers == nil {
		return &pb.StatusResponse{Status: false, Message: "Redis unavailable"}, nil
	}
	for update := range updates {
		if update == nil {
			continue
		}
		key := "driver_location:" + update.DriverId
		locData, err := json.Marshal(update.Location)
		if err != nil {
			continue // skip invalid location
		}
		_ = s.RedisDrivers.Set(ctx, key, locData, time.Hour).Err()
	}
	return &pb.StatusResponse{
		Status:  true,
		Message: "Location updates received successfully",
	}, nil
}
