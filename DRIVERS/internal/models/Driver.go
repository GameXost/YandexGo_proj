package models

import "time"

type Driver struct {
	ID        int64     `json:"id"`
	UserName  string    `json:"username"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Car       string    `json:"car"`
	Plate     string    `json:"plate"`
	Online    bool      `json:"online"`
	Location  string    `json:"location"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
