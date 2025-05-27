package models

import "time"

type User struct {
	ID        string    `json:"id"`
	UserName  string    `json:"name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	CreatedAt time.Time `json:"created"`
	UpdatedAt time.Time `json:"updated"`
}
