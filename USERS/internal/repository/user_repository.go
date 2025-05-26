package repository

import (
	"USERS/internal/models"
	"context"
)

type UserRepository interface {
	GetByID(ctx context.Context, id int64) (*models.User, error)
	GetByEmail(ctx context.Context, email string) (*models.User, error)
	Update(ctx context.Context, user *models.User) error
	// Добавьте другие методы по необходимости
}
