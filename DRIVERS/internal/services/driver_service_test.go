package services

import (
	"context"
	"errors"
	"testing"
	"time"

	"your_project/DRIVERS/internal/prometh"

	pb "github.com/GameXost/YandexGo_proj/DRIVERS/API/generated/drivers"
	"github.com/GameXost/YandexGo_proj/DRIVERS/internal/models"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/testutil"
	"github.com/redis/go-redis/v9"
)

// Мок-репозиторий
type mockRepo struct {
	GetDriverByIDFunc       func(ctx context.Context, id string) (*models.Driver, error)
	UpdateDriverProfileFunc func(ctx context.Context, driver *models.Driver) error
	GetUserByIDFunc         func(ctx context.Context, id string) (*models.User, error)
}

func (m *mockRepo) GetDriverByID(ctx context.Context, id string) (*models.Driver, error) {
	return m.GetDriverByIDFunc(ctx, id)
}

func (m *mockRepo) UpdateDriverProfile(ctx context.Context, driver *models.Driver) error {
	return m.UpdateDriverProfileFunc(ctx, driver)
}

func (m *mockRepo) GetUserByID(ctx context.Context, id string) (*models.User, error) {
	if m.GetUserByIDFunc != nil {
		return m.GetUserByIDFunc(ctx, id)
	}
	return nil, errors.New("not implemented")
}

// Мок-структура для Redis
type mockRedisClient struct {
	GetFunc func(ctx context.Context, key string) *redis.StringCmd
	SetFunc func(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd
	DelFunc func(ctx context.Context, keys ...string) *redis.IntCmd
}

func (m *mockRedisClient) Get(ctx context.Context, key string) *redis.StringCmd {
	return m.GetFunc(ctx, key)
}
func (m *mockRedisClient) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	return m.SetFunc(ctx, key, value, expiration)
}
func (m *mockRedisClient) Del(ctx context.Context, keys ...string) *redis.IntCmd {
	return m.DelFunc(ctx, keys...)
}

func TestGetDriverProfile_Success(t *testing.T) {
	// Новый реестр для изоляции теста
	reg := prometheus.NewRegistry()
	reg.MustRegister(prometh.RideAcceptedCounter)

	repo := &mockRepo{
		GetDriverByIDFunc: func(ctx context.Context, id string) (*models.Driver, error) {
			return &models.Driver{ID: "1", UserName: "Ivan"}, nil
		},
		UpdateDriverProfileFunc: func(ctx context.Context, driver *models.Driver) error {
			return nil
		},
	}
	service := &DriverService{
		Repo:  repo,
		Kafka: nil,
	}

	driver, err := service.GetDriverProfile(context.Background(), "1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if driver.Id != "1" || driver.Username != "Ivan" {
		t.Errorf("unexpected driver: %+v", driver)
	}

	// Проверяем, что метрика увеличилась
	value := testutil.ToFloat64(prometh.RideAcceptedCounter)
	if value != 1 {
		t.Errorf("expected RideAcceptedCounter to be 1, got %v", value)
	}
}

func TestGetDriverProfile_NotFound(t *testing.T) {
	repo := &mockRepo{
		GetDriverByIDFunc: func(ctx context.Context, id string) (*models.Driver, error) {
			return nil, errors.New("not found")
		},
	}
	service := &DriverService{Repo: repo}

	_, err := service.GetDriverProfile(context.Background(), "2")
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestUpdateDriverProfile_Success(t *testing.T) {
	mockRepo := &mockRepo{
		UpdateDriverProfileFunc: func(ctx context.Context, driver *models.Driver) error {
			return nil
		},
	}
	service := &DriverService{
		Repo:  mockRepo,
		Kafka: nil,
	}
	req := &pb.Driver{
		Id:        "1",
		Username:  "Ivan",
		Email:     "ivan@example.com",
		Phone:     "1234567890",
		CarNumber: "A123BC",
		CarModel:  "Toyota",
		CarMark:   "Corolla",
		CarColor:  "White",
	}
	updated, err := service.UpdateDriverProfile(context.Background(), req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if updated.Username != "Ivan" {
		t.Errorf("unexpected updated driver: %+v", updated)
	}
}

// func TestAcceptRide_Success(t *testing.T) {
// 	ride := pb.Ride{
// 		Id:     "ride123",
// 		Status: "pending",
// 	}
// 	rideData, _ := json.Marshal(ride)

// 	mockRedis := &mockRedisClient{
// 		GetFunc: func(ctx context.Context, key string) *redis.StringCmd {
// 			return redis.NewStringResult(string(rideData), nil)
// 		},
// 		SetFunc: func(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
// 			return redis.NewStatusResult("OK", nil)
// 		},
// 	}

// 	service := &DriverService{
// 		Repo:  nil,
// 		Kafka: nil,
// 	}

// 	resp, err := service.AcceptRide(context.Background(), "ride123", "driver1")
// 	if err != nil {
// 		t.Fatalf("unexpected error: %v", err)
// 	}
// 	if !resp.Status {
// 		t.Errorf("expected status true, got false")
// 	}
// 	if resp.Message != "Ride accepted successfully" {
// 		t.Errorf("unexpected message: %s", resp.Message)
// 	}
// }

// func TestAcceptRide_RideNotFound(t *testing.T) {
// 	mockRedis := &mockRedisClient{
// 		GetFunc: func(ctx context.Context, key string) *redis.StringCmd {
// 			return redis.NewStringResult("", redis.Nil)
// 		},
// 		SetFunc: func(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
// 			return redis.NewStatusResult("OK", nil)
// 		},
// 	}

// 	service := &DriverService{
// 		Repo:  nil,
// 		Redis: mockRedis,
// 		Kafka: nil,
// 	}

// 	resp, err := service.AcceptRide(context.Background(), "ride404", "driver1")
// 	if err != nil {
// 		t.Fatalf("unexpected error: %v", err)
// 	}
// 	if resp.Status {
// 		t.Errorf("expected status false, got true")
// 	}
// 	if resp.Message != "Ride not found" {
// 		t.Errorf("unexpected message: %s", resp.Message)
// 	}
// }

func TestMyMetric(t *testing.T) {
	// Сбросить метрику, если нужно (или создать новый реестр)
	reg := prometheus.NewRegistry()
	reg.MustRegister(prometh.RideAcceptedCounter)

	// Выполнить действие, которое должно инкрементировать метрику
	prometh.RideAcceptedCounter.WithLabelValues("foo").Inc()

	// Проверить значение метрики
	value := testutil.ToFloat64(prometh.RideAcceptedCounter.WithLabelValues("foo"))
	if value != 1 {
		t.Errorf("expected counter to be 1, got %v", value)
	}
}

func TestRideAcceptedMetric(t *testing.T) {
	// Новый реестр для изоляции теста
	reg := prometheus.NewRegistry()
	reg.MustRegister(prometh.RideAcceptedCounter)

	// Вызовите метод, который инкрементирует метрику
	prometh.RideAcceptedCounter.WithLabelValues("driver1").Inc()

	// Проверьте значение
	value := testutil.ToFloat64(prometh.RideAcceptedCounter.WithLabelValues("driver1"))
	if value != 1 {
		t.Errorf("expected counter to be 1, got %v", value)
	}
}
