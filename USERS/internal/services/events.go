package services

type BaseEvent struct {
	Event         string `json:"event"`
	Timestamp     int64  `json:"timestamp"`
	CorrelationID string `json:"correlation_id,omitempty"`
	ReplyTo       string `json:"reply_to,omitempty"`
}

type RideCreatedEvent struct {
	BaseEvent
	RideID        string `json:"ride_id"`
	PassengerID   string `json:"passenger_id"`
	StartLocation string `json:"pickup_location"`
	EndLocation   string `json:"dropoff_location"`
	Status        string `json:"status"`
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

type RideCompletedEvent struct {
	BaseEvent
	RideID   string `json:"ride_id"`
	DriverID string `json:"driver_id"`
}

type RideCanceledEvent struct {
	BaseEvent
	RideID   string `json:"ride_id"`
	DriverID string `json:"driver_id,omitempty"`
	Reason   string `json:"reason,omitempty"`
}

type RideCanceledResponse struct {
	BaseEvent
	RideID   string `json:"ride_id"`
	DriverID string `json:"driver_id,omitempty"`
	Reason   string `json:"reason,omitempty"`
	Status   string `json:"status,omitempty"`
}

type GetDriverLocationEvent struct {
	BaseEvent
	RideID   string `json:"ride_id"`
	DriverID string `json:"driver_id"`
}

type DriverLocationResponse struct {
	BaseEvent
	DriverID  string  `json:"driver_id"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type GetDriverInfoEvent struct {
	BaseEvent
	DriverID string `json:"driver_id"`
}
