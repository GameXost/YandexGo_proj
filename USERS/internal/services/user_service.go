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
	Repo  repository.UserRepositoryInterface
	Kafka *kafka.Writer
}

func NewUserService(repo repository.UserRepositoryInterface, kafka *kafka.Writer) *UserService {
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

func (s *UserService) RequestRide(ctx context.Context, req *pb.RideRequest) (*pb.Ride, error) {
	// TODO: реализовать создание поездки работает с кафка только
	return nil, fmt.Errorf("not implemented")
}

func (s *UserService) CancelRide(ctx context.Context, req *pb.RideIdRequest) (*pb.StatusResponse, error) {
	// TODO: реализовать отмену поездки только с кафка
	return nil, fmt.Errorf("not implemented")
}

func (s *UserService) GetRideStatus(ctx context.Context, req *pb.UserIdRequest) (*pb.Ride, error) {
	// TODO: реализовать получение статуса поездки только с кафка
	return nil, fmt.Errorf("not implemented")
}

func (s *UserService) GetRideHistory(ctx context.Context, req *pb.UserIdRequest) (*pb.RideHistoryResponse, error) {
	// TODO: реализовать получение истории поездок только с кафка
	return nil, fmt.Errorf("not implemented")
}

func (s *UserService) GetDriverLocation(ctx context.Context, req *pb.DriverIdRequest) (*pb.Location, error) {
	// TODO: реализовать получение локации водителя только с кафка
	return nil, fmt.Errorf("not implemented")
}

func (s *UserService) GetDriverInfo(ctx context.Context, req *pb.DriverIdRequest) (*pb.Driver, error) {
	// TODO: реализовать получение информации о водителе тоже только кафка
	return nil, fmt.Errorf("not implemented")
}
