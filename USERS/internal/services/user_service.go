package services

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"time"

	pb "github.com/GameXost/YandexGo_proj/USERS/API/generated/clients"
	"github.com/GameXost/YandexGo_proj/USERS/internal/models"
	"github.com/GameXost/YandexGo_proj/USERS/internal/repository"
	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
)

type UserService struct {
	Repo                   repository.UserRepositoryInterface
	Kafka                  *kafka.Writer
	OrderWaiters           map[string]*RideWaiter // rideID -> waiter
	WaitersMutex           sync.Mutex
	RidesTopic             string
	UserRequestsTopic      string
	DriverRequestTopic     string
	UserNotificationsTopic string
	ResponseWaiters        map[string]chan []byte
	ResponseWaitersMutex   sync.Mutex
	KafkaBrokers           []string
}

func NewUserService(repo repository.UserRepositoryInterface,
	kafkaWriter *kafka.Writer,
	ridesTopic,
	userRequestsTopic,
	userNotificationsTopic string,
	driverRequestTopic string,
	kafkaBrokers []string,
) *UserService {
	return &UserService{
		Repo:                   repo,
		Kafka:                  kafkaWriter,
		OrderWaiters:           make(map[string]*RideWaiter),
		RidesTopic:             ridesTopic,
		UserRequestsTopic:      userRequestsTopic,
		DriverRequestTopic:     driverRequestTopic,
		UserNotificationsTopic: userNotificationsTopic,
		ResponseWaiters:        make(map[string]chan []byte),
		ResponseWaitersMutex:   sync.Mutex{},
		KafkaBrokers:           kafkaBrokers, // Сохраняем список брокеров
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
	userID := req.UserId
	correlationId := uuid.New().String()
	replyTo := CreateReplyTopicName("user-ride-responses", correlationId)

	// Формируем событие
	event := RideCreatedEvent{
		BaseEvent:     NewBaseEvent("ride_created", correlationId, replyTo),
		RideID:        uuid.New().String(),
		PassengerID:   req.UserId,
		StartLocation: req.StartLocation.String(), // если нужно сериализовать в строку
		EndLocation:   req.EndLocation.String(),
		Status:        "pending",
	}
	responseChan := make(chan []byte, 1)
	s.ResponseWaitersMutex.Lock()
	s.ResponseWaiters[correlationId] = responseChan
	s.ResponseWaitersMutex.Unlock()

	defer func() {
		s.ResponseWaitersMutex.Lock()
		delete(s.ResponseWaiters, correlationId)
		close(responseChan)
		s.ResponseWaitersMutex.Unlock()
	}()
	err := s.sendKafkaMessage(ctx, s.RidesTopic, event)
	if err != nil {
		return nil, fmt.Errorf("failed to send kafka message: %w", err)
	}
	log.Printf("Sent ride_created for UserID: %s,CorrelationID,  waiting for response in : %s ", userID, correlationId, replyTo)
	select {
	case rawResponse := <-responseChan:
		var resp RideCreatedEvent
		if err := json.Unmarshal(rawResponse, &resp); err != nil {
			return nil, fmt.Errorf("failed to unmarshal response: %w", err)
		}
		return &pb.Ride{
			Id:            resp.RideID,
			UserId:        resp.PassengerID,
			DriverId:      "", // DriverID нет в RideCreatedEvent, он появится в RideAcceptedEvent
			StartLocation: req.StartLocation,
			EndLocation:   req.EndLocation,
			Status:        resp.Status,
			Timestamp:     resp.Timestamp,
		}, nil
	case <-time.After(10 * time.Second):
		return nil, fmt.Errorf("failed to send kafka message: timeout")
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

func (s *UserService) CancelRide(ctx context.Context, req *pb.RideIdRequest) (*pb.StatusResponse, error) {
	rideID := req.Id
	correlationId := uuid.New().String()
	replyTo := CreateReplyTopicName("ride-cancel-responses", correlationId)

	event := RideCanceledEvent{
		BaseEvent: NewBaseEvent("ride_cancelled", correlationId, replyTo),
		RideID:    rideID,
		Reason:    "user_cancelled",
	}
	responseChan := make(chan []byte, 1)
	s.ResponseWaitersMutex.Lock()
	s.ResponseWaiters[correlationId] = responseChan
	s.ResponseWaitersMutex.Unlock()

	defer func() {
		s.ResponseWaitersMutex.Lock()
		delete(s.ResponseWaiters, correlationId)
		close(responseChan)
		s.ResponseWaitersMutex.Unlock()
	}()
	err := s.sendKafkaMessage(ctx, s.RidesTopic, event)
	if err != nil {
		return nil, fmt.Errorf("failed to send kafka message: %w", err)
	}
	log.Printf("sent ride_canceled for RideID : %s, CorrelationID %s waiting for person in %s", rideID, correlationId, replyTo)

	select {
	case rawResponse := <-responseChan:
		var resp RideCanceledResponse
		if err := json.Unmarshal(rawResponse, &resp); err != nil {
			return nil, fmt.Errorf("failed to unmarshal response: %w", err)
		}
		return &pb.StatusResponse{
			Status:  resp.Status == "canceled" || resp.Status == "success",
			Message: resp.Reason,
		}, nil
	case <-time.After(10 * time.Second):
		return nil, fmt.Errorf("timeout waiting for ride canceled response for RideID: %s (CorrelationID: %s)", rideID, correlationId)
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

func (s *UserService) GetRideStatus(ctx context.Context, req *pb.UserIdRequest) (*pb.Ride, error) {
	userID := req.Id
	correlationId := uuid.New().String()
	replyTo := CreateReplyTopicName("user-ride-status-response", correlationId)

	statusRequest := GetRideStatusRequest{
		BaseEvent: NewBaseEvent("get_ride_status_request", correlationId, replyTo),
		UserID:    userID,
	}
	responseChan := make(chan []byte, 1)
	s.ResponseWaitersMutex.Lock()
	s.ResponseWaiters[correlationId] = responseChan
	s.ResponseWaitersMutex.Unlock()

	defer func() {
		s.ResponseWaitersMutex.Lock()
		delete(s.ResponseWaiters, correlationId)
		close(responseChan)
		s.ResponseWaitersMutex.Unlock()
	}()

	err := s.sendKafkaMessage(ctx, s.RidesTopic, statusRequest)
	if err != nil {
		return nil, fmt.Errorf("failed to send get_ride_status request: %w", err)
	}
	log.Printf("sent get_ride_status_request for userid: %s, correlationID: %s, waiting for response in %s", userID, correlationId, replyTo)

	select {
	case rawResponse := <-responseChan:
		var resp pb.Ride
		if err := json.Unmarshal(rawResponse, &resp); err != nil {
			return nil, fmt.Errorf("failed to unmarshal ride status response: %w", err)
		}
		return &resp, nil
	case <-time.After(10 * time.Second):
		return nil, fmt.Errorf("timeout waitning for ride status response for userid: %s, correlationID: %s", userID, correlationId)
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

func (s *UserService) GetRideHistory(ctx context.Context, req *pb.UserIdRequest) (*pb.RideHistoryResponse, error) {
	userID := req.Id
	correlationId := uuid.New().String()
	replyTo := CreateReplyTopicName("user-ride-history-responses", correlationId)

	historyRequest := GetRideHistoryRequest{
		BaseEvent: NewBaseEvent("get_ride_history_request", correlationId, replyTo),
		UserID:    userID,
	}
	responseChan := make(chan []byte, 1)
	s.ResponseWaitersMutex.Lock()
	s.ResponseWaiters[correlationId] = responseChan
	s.ResponseWaitersMutex.Unlock()

	defer func() {
		s.ResponseWaitersMutex.Lock()
		delete(s.ResponseWaiters, correlationId)
		close(responseChan)
		s.ResponseWaitersMutex.Unlock()
	}()

	err := s.sendKafkaMessage(ctx, s.RidesTopic, historyRequest)
	if err != nil {
		return nil, fmt.Errorf("failed to send get_ride_history request: %w", err)
	}
	log.Printf("Sent get_ride_history_request for UserID: %s, CorrelationID: %s, waiting for response in %s", userID, correlationId, replyTo)

	select {
	case rawResponse := <-responseChan:
		var resp pb.RideHistoryResponse
		if err := json.Unmarshal(rawResponse, &resp); err != nil {
			return nil, fmt.Errorf("failed to unmarshal ride history response: %w", err)
		}
		return &resp, nil
	case <-time.After(10 * time.Second):
		return nil, fmt.Errorf("timeout waiting for ride history response for UserID: %s (CorrelationID: %s)", userID, correlationId)
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

func (s *UserService) GetDriverLocation(ctx context.Context, req *pb.DriverIdRequest) (*pb.Location, error) {
	correlationID := uuid.New().String()
	replyToTopic := CreateReplyTopicName("reply-to-users", correlationID)
	request := GetDriverLocationRequest{
		BaseEvent: NewBaseEvent("get_driver_location_request", correlationID, replyToTopic),
		DriverID:  req.Id,
	}
	responseChan := make(chan []byte, 1)
	s.ResponseWaitersMutex.Lock()
	s.ResponseWaiters[correlationID] = responseChan
	s.ResponseWaitersMutex.Unlock()

	defer func() {
		s.ResponseWaitersMutex.Lock()
		delete(s.ResponseWaiters, correlationID)
		close(responseChan)
		s.ResponseWaitersMutex.Unlock()
	}()
	if err := s.sendKafkaMessage(ctx, s.DriverRequestTopic, request); err != nil {
		return nil, fmt.Errorf("failed to publish get_driver_location_request: %w", err)
	}

	log.Printf("Sent get_driver_location_request for DriverID: %s, CorrelationID: %s, waiting for response in %s", req.Id, correlationID, replyToTopic)

	select {
	case rawResponse := <-responseChan:
		var response DriverLocationResponse
		if err := json.Unmarshal(rawResponse, &response); err != nil {
			return nil, fmt.Errorf("failed to unmarshal driver location response: %w", err)
		}
		if response.Error != "" {
			return nil, fmt.Errorf("driver service error: %s", response.Error)
		}
		return &pb.Location{
			Latitude:  response.Latitude,
			Longitude: response.Longitude,
		}, nil
	case <-time.After(10 * time.Second):
		return nil, fmt.Errorf("timeout waiting for driver location response for DriverID: %s (CorrelationID: %s)", req.Id, correlationID)
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

func (s *UserService) GetDriverInfo(ctx context.Context, req *pb.DriverIdRequest) (*pb.Driver, error) {
	correlationID := uuid.New().String()
	replyToTopic := CreateReplyTopicName("reply-to-users", correlationID)

	driverInfoRequest := GetDriverInfoRequest{ // <--- ИСПОЛЬЗУЕМ GetDriverInfoRequest
		BaseEvent: NewBaseEvent("get_driver_info_request", correlationID, replyToTopic), // <--- ИСПОЛЬЗУЕМ NewBaseEvent
		DriverID:  req.Id,
	}

	responseChan := make(chan []byte, 1)
	s.ResponseWaitersMutex.Lock()
	s.ResponseWaiters[correlationID] = responseChan
	s.ResponseWaitersMutex.Unlock()
	defer func() {
		s.ResponseWaitersMutex.Lock()
		delete(s.ResponseWaiters, correlationID)
		close(responseChan)
		s.ResponseWaitersMutex.Unlock()
	}()

	if err := s.sendKafkaMessage(ctx, s.DriverRequestTopic, driverInfoRequest); err != nil {
		return nil, fmt.Errorf("failed to send driver info request: %w", err)
	}
	log.Printf("Sent get_driver_info_request for DriverID %s, CorrelationID %s, waitong for response in %s", req.Id, correlationID, replyToTopic)

	select {
	case rawResponse := <-responseChan:
		var response pb.Driver
		if err := json.Unmarshal(rawResponse, &response); err != nil {
			return nil, fmt.Errorf("failed to unmarshal response: %w", err)
		}
		return &response, nil
	case <-time.After(10 * time.Second):
		return nil, fmt.Errorf("timeout waitng for driver info response fior DriverID: %s (CorrelationID %s)", req.Id, correlationID)
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

func (s *UserService) sendKafkaMessage(ctx context.Context, topic string, event interface{}) error {
	data, err := json.Marshal(event)
	if err != nil {
		return fmt.Errorf("failed to marshal event: %w", err)
	}
	msg := &kafka.Message{
		Topic: topic,
		Value: data,
	}
	return s.Kafka.WriteMessages(ctx, *msg)
}
