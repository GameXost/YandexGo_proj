package services

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	pb "github.com/GameXost/YandexGo_proj/USERS/API/generated/clients"
	"github.com/GameXost/YandexGo_proj/USERS/internal/models"
	"github.com/GameXost/YandexGo_proj/USERS/internal/repository"
	"github.com/google/uuid"
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
	// Генерируем correlationId и replyTo
	userID := req.UserId
	correlationId := uuid.New().String()
	replyTo := "user-" + userID + "-responses"

	// Формируем событие
	event := RideCreatedEvent{
		BaseEvent: BaseEvent{
			Event:         "ride_created",
			CorrelationID: correlationId,
			ReplyTo:       replyTo,
			Timestamp:     time.Now().Unix(),
		},
		RideID:        uuid.New().String(),
		PassengerID:   req.UserId,
		StartLocation: req.StartLocation.String(), // если нужно сериализовать в строку
		EndLocation:   req.EndLocation.String(),
		Status:        "pending",
	}

	// Отправляем событие в Kafka
	err := PublishEvent(ctx, s.Kafka, "ride-requests", event, event.RideID)
	if err != nil {
		return nil, fmt.Errorf("failed to send ride request: %w", err)
	}

	// Ждём ответ (например, ride_accepted или ride_created)
	brokers := []string{"localhost:9092"} // или из конфига
	resp, err := s.WaitForRideCreatedResponse(ctx, brokers, replyTo, correlationId, 10*time.Second)
	if err != nil {
		return nil, fmt.Errorf("timeout or error waiting for response: %w", err)
	}

	// Собираем pb.Ride из ответа
	return &pb.Ride{
		Id:            resp.RideID,
		UserId:        resp.PassengerID,
		DriverId:      resp.DriverID,
		StartLocation: req.StartLocation,
		EndLocation:   req.EndLocation,
		Status:        resp.Status,
		Timestamp:     resp.Timestamp,
	}, nil
}

func (s *UserService) CancelRide(ctx context.Context, req *pb.RideIdRequest) (*pb.StatusResponse, error) {
	rideID := req.Id
	correlationId := uuid.New().String()
	replyTo := "ride-cancel-responses-" + rideID // replyTo теперь уникален по rideID

	event := RideCanceledEvent{
		BaseEvent: BaseEvent{
			Event:         "ride_canceled",
			CorrelationID: correlationId,
			ReplyTo:       replyTo,
			Timestamp:     time.Now().Unix(),
		},
		RideID: rideID,
		Reason: "user_cancelled",
	}

	err := PublishEvent(ctx, s.Kafka, "ride-requests", event, event.RideID)
	if err != nil {
		return nil, fmt.Errorf("failed to send cancel ride request: %w", err)
	}

	brokers := []string{"localhost:9092"} // или из конфига
	resp, err := s.WaitForRideCanceledResponse(ctx, brokers, replyTo, correlationId, 10*time.Second)
	if err != nil {
		return nil, fmt.Errorf("timeout or error waiting for response: %w", err)
	}

	return &pb.StatusResponse{
		Status:  resp.Status == "canceled",
		Message: resp.Reason,
	}, nil
}

func (s *UserService) GetRideStatus(ctx context.Context, req *pb.UserIdRequest) (*pb.Ride, error) {
	userID := req.Id
	correlationId := uuid.New().String()
	replyTo := "user-" + userID + "-responses"

	// Формируем событие-запрос
	event := BaseEvent{
		Event:         "get_ride_status",
		CorrelationID: correlationId,
		ReplyTo:       replyTo,
		Timestamp:     time.Now().Unix(),
	}

	err := PublishEvent(ctx, s.Kafka, "ride-requests", event, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to send get_ride_status request: %w", err)
	}

	brokers := []string{"localhost:9092"}
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: brokers,
		Topic:   replyTo,
	})
	defer reader.Close()

	timeoutCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	for {
		m, err := reader.ReadMessage(timeoutCtx)
		if err != nil {
			return nil, err
		}
		var resp pb.Ride
		if err := json.Unmarshal(m.Value, &resp); err != nil {
			continue
		}
		// Можно добавить проверку correlationId, если оно есть в ответе
		return &resp, nil
	}
}

func (s *UserService) GetRideHistory(ctx context.Context, req *pb.UserIdRequest) (*pb.RideHistoryResponse, error) {
	userID := req.Id
	correlationId := uuid.New().String()
	replyTo := "user-" + userID + "-responses"

	event := BaseEvent{
		Event:         "get_ride_history",
		CorrelationID: correlationId,
		ReplyTo:       replyTo,
		Timestamp:     time.Now().Unix(),
	}

	err := PublishEvent(ctx, s.Kafka, "ride-requests", event, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to send get_ride_history request: %w", err)
	}

	brokers := []string{"localhost:9092"}
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: brokers,
		Topic:   replyTo,
	})
	defer reader.Close()

	timeoutCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	for {
		m, err := reader.ReadMessage(timeoutCtx)
		if err != nil {
			return nil, err
		}
		var resp pb.RideHistoryResponse
		if err := json.Unmarshal(m.Value, &resp); err != nil {
			continue
		}
		return &resp, nil
	}
}

func (s *UserService) GetDriverLocation(ctx context.Context, req *pb.DriverIdRequest) (*pb.Location, error) {
	correlationId := uuid.New().String()
	replyTo := "user-driver-location-" + req.Id

	event := GetDriverLocationEvent{
		BaseEvent: BaseEvent{
			Event:         "get_driver_location",
			CorrelationID: correlationId,
			ReplyTo:       replyTo,
			Timestamp:     time.Now().Unix(),
		},
		RideID:   "", // если есть rideID, подставь его
		DriverID: req.Id,
	}

	err := PublishEvent(ctx, s.Kafka, "ride-requests", event, req.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to send get_driver_location request: %w", err)
	}

	brokers := []string{"localhost:9092"}
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: brokers,
		Topic:   replyTo,
	})
	defer reader.Close()

	timeoutCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	for {
		m, err := reader.ReadMessage(timeoutCtx)
		if err != nil {
			return nil, err
		}
		var resp DriverLocationResponse
		if err := json.Unmarshal(m.Value, &resp); err != nil {
			continue
		}
		return &pb.Location{
			Latitude:  resp.Latitude,
			Longitude: resp.Longitude,
		}, nil
	}
}

func (s *UserService) GetDriverInfo(ctx context.Context, req *pb.DriverIdRequest) (*pb.Driver, error) {
	correlationId := uuid.New().String()
	replyTo := "user-driver-info-" + req.Id

	event := BaseEvent{
		Event:         "get_driver_info",
		CorrelationID: correlationId,
		ReplyTo:       replyTo,
		Timestamp:     time.Now().Unix(),
	}

	err := PublishEvent(ctx, s.Kafka, "ride-requests", event, req.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to send get_driver_info request: %w", err)
	}

	brokers := []string{"localhost:9092"}
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: brokers,
		Topic:   replyTo,
	})
	defer reader.Close()

	timeoutCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	for {
		m, err := reader.ReadMessage(timeoutCtx)
		if err != nil {
			return nil, err
		}
		var resp pb.Driver
		if err := json.Unmarshal(m.Value, &resp); err != nil {
			continue
		}
		return &resp, nil
	}
}

// Отправка запроса на получение координат водителя с correlationId и replyTo
func (s *UserService) RequestDriverLocation(ctx context.Context, userID, rideID, driverID string) (correlationId, replyTo string, err error) {
	correlationId = uuid.New().String()
	replyTo = "user-" + userID + "-responses"

	event := GetDriverLocationEvent{
		BaseEvent: BaseEvent{
			Event:         "get_driver_location",
			CorrelationID: correlationId,
			ReplyTo:       replyTo,
			Timestamp:     time.Now().Unix(),
		},
		RideID:   rideID,
		DriverID: driverID,
	}
	err = PublishEvent(ctx, s.Kafka, "ride-requests", event, event.RideID)
	return
}

// Ожидание ответа по replyTo-топику и correlationId
func (s *UserService) WaitForDriverLocationResponse(ctx context.Context, brokers []string, replyTo, correlationId string, timeout time.Duration) (*DriverLocationResponse, error) {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: brokers,
		Topic:   replyTo,
	})
	defer reader.Close()

	timeoutCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	for {
		m, err := reader.ReadMessage(timeoutCtx)
		if err != nil {
			return nil, err
		}
		var resp DriverLocationResponse
		if err := json.Unmarshal(m.Value, &resp); err != nil {
			continue
		}
		if resp.CorrelationID == correlationId {
			return &resp, nil
		}
	}
}

func (s *UserService) WaitForRideCanceledResponse(ctx context.Context, brokers []string, replyTo, correlationId string, timeout time.Duration) (*RideCanceledResponse, error) {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: brokers,
		Topic:   replyTo,
	})
	defer reader.Close()

	timeoutCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	for {
		m, err := reader.ReadMessage(timeoutCtx)
		if err != nil {
			return nil, err
		}
		var resp RideCanceledResponse
		if err := json.Unmarshal(m.Value, &resp); err != nil {
			continue
		}
		if resp.CorrelationID == correlationId {
			return &resp, nil
		}
	}
}

func (s *UserService) WaitForRideCreatedResponse(ctx context.Context, brokers []string, replyTo, correlationId string, timeout time.Duration) (*RideCreatedEvent, error) {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: brokers,
		Topic:   replyTo,
	})
	defer reader.Close()

	timeoutCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	for {
		m, err := reader.ReadMessage(timeoutCtx)
		if err != nil {
			return nil, err
		}
		var resp RideCreatedEvent
		if err := json.Unmarshal(m.Value, &resp); err != nil {
			continue
		}
		if resp.CorrelationID == correlationId {
			return &resp, nil
		}
	}
}
