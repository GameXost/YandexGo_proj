package services

import (
	"context"
	"fmt"

	pb "github.com/GameXost/YandexGo_proj/USERS/API/generated/clients"
	"github.com/GameXost/YandexGo_proj/USERS/internal/models"
	"github.com/GameXost/YandexGo_proj/USERS/internal/repository"
	"github.com/segmentio/kafka-go"
)

type UserService struct {
	Repo  repository.UserRepository
	Kafka *kafka.Writer
}

func NewUserService(repo repository.UserRepository, kafka *kafka.Writer) *UserService {
	return &UserService{
		Repo:  repo,
		Kafka: kafka,
	}
}

func (s *UserService) GetUserProfile(ctx context.Context, userID string) (*pb.User, error) {
	user, err := s.Repo.GetUserByID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("invalid user id: %w", err)
	}
	return modelToProtoUser(user), nil
}

func (s *UserService) UpdateUserProfile(ctx context.Context, req *pb.UpdateProfileRequest) (*pb.User, error) {
	err := s.Repo.UpdateUserProfile(ctx, &models.User{
		ID:       req.Id,
		UserName: req.Username,
		Email:    req.Email,
		Phone:    req.Phone,
	})
	if err != nil {
		return nil, err
	}
	return modelToProtoUser(&models.User{
		ID:       req.Id,
		UserName: req.Username,
		Email:    req.Email,
		Phone:    req.Phone,
	}), nil
}

func modelToProtoUser(m *models.User) *pb.User {
	return &pb.User{
		Id:       m.ID,
		Username: m.UserName,
		Email:    m.Email,
		Phone:    m.Phone,
	}
}
