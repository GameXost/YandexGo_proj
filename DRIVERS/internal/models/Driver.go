package models

import "time"

type Driver struct {
	ID         string    `json:"id"`
	UserName   string    `json:"username"`
	Email      string    `json:"email"`
	Phone      string    `json:"phone"`
	Car        string    `json:"car"`
	Car_number string    `json:"car_number"`
	Car_model  string    `json:"car_model"`
	Car_marks  string    `json:"car_marks"`
	Car_color  string    `json:"car_color"`
	Online     bool      `json:"online"`
	Location   string    `json:"location"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
