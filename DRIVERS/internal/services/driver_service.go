package services

import (
	"context"
	"encoding/json"
	"fmt"
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
	Repo  repository.DriverRepositoryInterface
	Redis RedisClient
	Kafka *kafka.Writer
}

func NewDriverService(repo *repository.DriverRepository, redis *redis.Client, kafka *kafka.Writer) *DriverService {
	return &DriverService{
		Repo:  repo,
		Redis: redis,
		Kafka: kafka,
	}
}

func (s *DriverService) GetDriverProfile(ctx context.Context, driverID string) (*pb.Driver, error) {
	cacheKey := "driver_profile:" + driverID

	if s.Redis != nil {
		cached, err := s.Redis.Get(ctx, cacheKey).Result()
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

	if s.Redis != nil {
		if data, err := json.Marshal(driver); err == nil {
			_ = s.Redis.Set(ctx, cacheKey, data, time.Hour).Err()
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
	if s.Redis != nil {
		_ = s.Redis.Del(ctx, cacheKey).Err()
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
	if s.Redis == nil {
		return &pb.StatusResponse{Status: false, Message: "Redis unavailable"}, nil
	}

	// 1. Получить ride из Redis
	rideKey := "ride:" + rideID
	rideData, err := s.Redis.Get(ctx, rideKey).Result()
	if err == redis.Nil {
		return &pb.StatusResponse{Status: false, Message: "Ride not found"}, nil
	} else if err != nil {
		return &pb.StatusResponse{Status: false, Message: "Redis error"}, nil
	}

	var ride pb.Ride
	if err := json.Unmarshal([]byte(rideData), &ride); err != nil {
		return &pb.StatusResponse{Status: false, Message: "Invalid ride data"}, nil
	}
	//если хуйня в обработке то игнорим
	if ride.Status != "pending" {
		return &pb.StatusResponse{Status: false, Message: "Ride already accepted or completed"}, nil
	}
	ride.Status = "accepted"
	ride.DriverId = driverID

	// сохраняем в редиску ебучую
	updatedData, err := json.Marshal(ride)
	if err != nil {
		return &pb.StatusResponse{Status: false, Message: "Failed to serialize ride"}, nil
	}
	if err := s.Redis.Set(ctx, rideKey, updatedData, time.Hour).Err(); err != nil {
		return &pb.StatusResponse{Status: false, Message: "Failed to update ride in Redis"}, nil
	}

	// 6. Сохранить rideID как текущий для водителя
	driverKey := "driver:" + driverID + ":current_ride"
	if err := s.Redis.Set(ctx, driverKey, rideID, time.Hour).Err(); err != nil {
		return &pb.StatusResponse{Status: false, Message: "Failed to set current ride for driver"}, nil
	}

	// пиздует в кафку
	if s.Kafka != nil {
		msg := kafka.Message{
			Key:   []byte(driverID),
			Value: []byte(fmt.Sprintf("accepted ride: %s", rideID)),
		}
		_ = s.Kafka.WriteMessages(ctx, msg)
	}

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
