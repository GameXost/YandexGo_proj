package services

import (
	"fmt"
	//pb "github.com/GameXost/YandexGo_proj/DRIVERS/API/generated/drivers"
	users_pb "github.com/GameXost/YandexGo_proj/USERS/API/generated/clients"

	"time"
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

func (r RideCreatedEvent) GetPassengerID() string {
	return r.PassengerID
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

func (r RideAcceptedEvent) GetDriverID() string {
	return r.DriverID
}

type RideCompletedEvent struct {
	BaseEvent
	RideID   string `json:"ride_id"`
	DriverID string `json:"driver_id"`
}

func (r RideCompletedEvent) GetRideID() string {
	return r.RideID
}
func (r RideCompletedEvent) GetDriverID() string {
	return r.DriverID
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
func (r RideCanceledEvent) GetDriverID() string {
	return r.DriverID
}

type RideCanceledResponse struct {
	BaseEvent
	RideID   string `json:"ride_id"`
	DriverID string `json:"driver_id,omitempty"`
	Reason   string `json:"reason,omitempty"`
	Status   string `json:"status,omitempty"`
}

func (r RideCanceledResponse) GetRideID() string {
	return r.RideID
}
func (r RideCanceledResponse) GetDriverID() string {
	return r.DriverID
}

type GetDriverLocationEvent struct {
	BaseEvent
	RideID   string `json:"ride_id"`
	DriverID string `json:"driver_id"`
}

func (r GetDriverLocationEvent) GetRideID() string {
	return r.RideID
}

func (r GetDriverLocationEvent) GetDriverID() string {
	return r.DriverID
}

type DriverLocationResponse struct {
	BaseEvent
	DriverID  string  `json:"driver_id"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Error     string  `json:"error,omitempty"`
}

func (r DriverLocationResponse) GetDriverID() string {
	return r.DriverID
}

func (r DriverLocationResponse) GetError() string {
	return r.Error
}
func (r DriverLocationResponse) GetLatitude() float64 {
	return r.Latitude
}
func (r DriverLocationResponse) GetLongitude() float64 {
	return r.Longitude
}

type GetDriverInfoEvent struct {
	BaseEvent
	DriverID string `json:"driver_id"`
}

func (r GetDriverInfoEvent) GetDriverID() string {
	return r.DriverID
}

type GetPassengerInfoRequest struct {
	BaseEvent
	PassengerID string `json:"passenger_id"`
}

func (r GetPassengerInfoRequest) GetPassengerID() string {
	return r.PassengerID
}

type UserProfileResponse struct {
	BaseEvent
	User  *users_pb.User `json:"user,omitempty"`
	Error string         `json:"error,omitempty"`
}

func (r UserProfileResponse) GetUser() interface{} {
	return r.User
}

func (r UserProfileResponse) GetError() string {
	return r.Error
}
