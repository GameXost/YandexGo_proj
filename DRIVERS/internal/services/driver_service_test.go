package services

import (
	"context"
	"errors"
	"testing"

	"github.com/GameXost/YandexGo_proj/DRIVERS/internal/models"
)

// Мок-репозиторий
type mockRepo struct {
	GetDriverByIDFunc       func(ctx context.Context, id string) (*models.Driver, error)
	UpdateDriverProfileFunc func(ctx context.Context, driver *models.Driver) error
}

func (m *mockRepo) GetDriverByID(ctx context.Context, id string) (*models.Driver, error) {
	return m.GetDriverByIDFunc(ctx, id)
}

func (m *mockRepo) UpdateDriverProfile(ctx context.Context, driver *models.Driver) error {
	return m.UpdateDriverProfileFunc(ctx, driver)
}

func TestGetDriverProfile_Success(t *testing.T) {
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
		Redis: nil,
		Kafka: nil,
	}

	driver, err := service.GetDriverProfile(context.Background(), "1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if driver.Id != "1" || driver.Username != "Ivan" {
		t.Errorf("unexpected driver: %+v", driver)
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
