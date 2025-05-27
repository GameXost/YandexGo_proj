package services

import (
	"USERS/internal/models"
	"USERS/internal/repository"
	"context"
	"encoding/json"
	"fmt"
	"github.com/segmentio/kafka-go"
)

type UserService struct {
	Repo repository.UserRepository
	Kafka *kafka.Writer
}

func NewUserService(repo *repository.UserRepository, kafka *kafka.Writer) *UserService {
	return &UserService{
		Repo: repo,
		Kafka: kafka,
	}
}

func (s *UserService) GetUserProfile(ctx context.Context, userID string) (*pb.User, error) {
	cacheKey := "user_profile:" + userID
	if err == nil && cached != "" {
		var user pb.User
		if err := json.Unmarshal([]byte(cached), &user); err == nil {
			return &user, nil
		}
	}

	user, err := s.Repo.GetUserByID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("get user by id not found: %w", err)
	}
	if s.RedisRides

}
	return s.repo.GetByID(ctx, id)
}

func (s *UserService) UpdateUserProfile(ctx context.Context, user *models.User) error {
	return s.repo.Update(ctx, user)
}
