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

type DriverService struct {
	Repo  repository.DriverRepositoryInterface
	Redis *redis.Client
	Kafka *kafka.Writer
}

func NewDriverService(repo *repository.DriverRepository, redis *redis.Client, kafka *kafka.Writer) *DriverService {
	return &DriverService{
		Repo:  repo,
		Redis: redis,
		Kafka: kafka,
	}
}

// GetDriverProfile возвращает профиль водителя по driverID с кешированием в Redis и логированием в Kafka
func (s *DriverService) GetDriverProfile(ctx context.Context, driverID string) (*pb.Driver, error) {
	cacheKey := "driver_profile:" + driverID

	if s.Redis != nil {
		cached, err := s.Redis.Get(ctx, cacheKey).Result()
		if err == nil && cached != "" {
			var driver pb.Driver
			if err := json.Unmarshal([]byte(cached), &driver); err == nil {
				return &driver, nil
			}
			// Если не удалось распарсить — логируем, но идём дальше
		}
	}

	// Получаем из БД
	driver, err := s.Repo.GetDriverByID(ctx, driverID)
	if err != nil {
		return nil, fmt.Errorf("driver not found: %w", err)
	}

	// Кладём в кеш
	if s.Redis != nil {
		if data, err := json.Marshal(driver); err == nil {
			_ = s.Redis.Set(ctx, cacheKey, data, time.Hour).Err()
		}
	}

	// (Опционально) Логируем в Kafka

	if s.Kafka != nil {
		msg := kafka.Message{
			Key:   []byte(driverID),
			Value: []byte(fmt.Sprintf("profile viewed: %s", driverID)),
		}
		_ = s.Kafka.WriteMessages(ctx, msg) // не делаем ошибку критичной
	}

	return modelToProtoDriver(driver), nil
}

func (s *DriverService) UpdateDriverProfile(ctx context.Context, req *pb.Driver) (*pb.Driver, error) {
	driver := &models.Driver{
		ID:         req.Id,
		UserName:   req.Username,
		Email:      req.Email,
		Phone:      req.Phone,
		Car_color:  req.CarColor,
		Car_number: req.CarNumber,
		Car_model:  req.CarModel,
		Car_marks:  req.CarMark,
	}
	err := s.Repo.UpdateDriverProfile(ctx, driver)
	if err != nil {
		return nil, err
	}
	cacheKey := "driver_profile:" + driver.ID
	_ = s.Redis.Del(ctx, cacheKey).Err()
	if s.Kafka != nil {
		msg := kafka.Message{
			Key:   []byte(driver.ID),
			Value: []byte("profile updated: " + driver.ID),
		}
		_ = s.Kafka.WriteMessages(ctx, msg)
	}
	return modelToProtoDriver(driver), nil
}

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
