package services

import (
	"fmt"
	"time"

	pb "github.com/GameXost/YandexGo_proj/USERS/API/generated/clients"
)

type BaseEvent struct {
	Event         string `json:"event"`
	Timestamp     int64  `json:"timestamp"`
	CorrelationID string `json:"correlation_id,omitempty"`
	ReplyTo       string `json:"reply_to,omitempty"`
}

func (b BaseEvent) GetCorrelationID() string {
	return b.CorrelationID
}

type RideCreatedEvent struct {
	BaseEvent
	RideID        string `json:"ride_id"`
	PassengerID   string `json:"passenger_id"`
	StartLocation string `json:"pickup_location"`
	EndLocation   string `json:"dropoff_location"`
	Status        string `json:"status"`
}

func (r RideCreatedEvent) GetRideID() string {
	return r.RideID
}

type RideAcceptedEvent struct {
	BaseEvent
	RideID        string `json:"ride_id"`
	DriverID      string `json:"driver_id"`
	PassengerID   string `json:"passenger_id"`
	StartLocation string `json:"pickup_location"`
	EndLocation   string `json:"dropoff_location"`
	Status        string `json:"status"`
}

func (r RideAcceptedEvent) GetRideID() string {
	return r.RideID
}

type RideCompletedEvent struct {
	BaseEvent
	RideID   string `json:"ride_id"`
	DriverID string `json:"driver_id"`
}

func (r RideCompletedEvent) GetRideID() string {
	return r.RideID
}

type RideCanceledEvent struct {
	BaseEvent
	RideID   string `json:"ride_id"`
	DriverID string `json:"driver_id,omitempty"`
	Reason   string `json:"reason,omitempty"`
}

func (r RideCanceledEvent) GetRideID() string {
	return r.RideID
}

type RideCanceledResponse struct {
	BaseEvent
	RideID   string `json:"ride_id"`
	DriverID string `json:"driver_id,omitempty"`
	Reason   string `json:"reason,omitempty"`
	Status   string `json:"status,omitempty"`
}

type GetUserProfileRequest struct {
	BaseEvent
	UserID string `json:"user_id"`
}

type UserProfileResponse struct {
	BaseEvent
	User  *pb.User `json:"user,omitempty"`
	Error string   `json:"error,omitempty"`
}

type GetDriverLocationRequest struct {
	BaseEvent
	DriverID string `json:"driver_id"`
}

type DriverLocationResponse struct {
	BaseEvent
	DriverID  string  `json:"driver_id"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Error     string  `json:"error,omitempty"`
}

type GetDriverInfoRequest struct {
	BaseEvent
	DriverID string `json:"driver_id"`
}

type DriverInfoResponse struct {
	BaseEvent
	Driver *pb.Driver `json:"driver,omitempty"`
	Error  string     `json:"error,omitempty"`
}

type GetRideStatusRequest struct {
	BaseEvent
	UserID string `json:"user_id"`
}

type GetRideHistoryRequest struct {
	BaseEvent
	UserID string `json:"user_id"`
}

func NewBaseEvent(eventType, correlationID, replyTo string) BaseEvent {
	return BaseEvent{
		Event:         eventType,
		CorrelationID: correlationID,
		ReplyTo:       replyTo,
		Timestamp:     time.Now().Unix(),
	}
}

func CreateReplyTopicName(baseTopic, correlationID string) string {
	return fmt.Sprintf("%s-%s", baseTopic, correlationID)
}
