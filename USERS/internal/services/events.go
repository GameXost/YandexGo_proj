package services

type RideCreatedEvent struct {
	Event         string `json:"event"`
	RideID        string `json:"ride_id"`
	PassengerID   string `json:"passenger_id"`
	StartLocation string `json:"pickup_location"`
	EndLocation   string `json:"dropoff_location"`
	Timestamp     int64  `json:"timestamp"`
	Status        string `json:"status"`
}

type RideAcceptedEvent struct {
	Event         string `json:"event"`
	RideID        string `json:"ride_id"`
	PassengerID   string `json:"passenger_id"`
	DriverID      string `json:"driver_id"`
	StartLocation string `json:"pickup_location"`
	EndLocation   string `json:"dropoff_location"`
	Timestamp     int64  `json:"timestamp"`
	Status        string `json:"status"`
}

type RideCompletedEvent struct {
	Event     string `json:"event"`
	RideID    string `json:"ride_id"`
	DriverID  string `json:"driver_id"`
	Timestamp int64  `json:"timestamp"`
	// Можно добавить duration, итоговую стоимость, координаты завершения и т.д.

}

type RideCanceledEvent struct {
	Event     string `json:"event"`
	RideID    string `json:"ride_id"`
	DriverID  string `json:"driver_id,omitempty"`
	Reason    string `json:"reason,omitempty"`
	Timestamp int64  `json:"timestamp"`
}
