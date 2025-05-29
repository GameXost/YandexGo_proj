package repository

import (
	"context"
	"fmt"

	"github.com/GameXost/YandexGo_proj/USERS/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepositoryInterface interface {
	GetUserByID(ctx context.Context, userID string) (*models.User, error)
	UpdateUserProfile(ctx context.Context, user *models.User) error
}

type UserRepository struct {
	DB *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (r *UserRepository) GetUserByID(ctx context.Context, userID string) (*models.User, error) {
	query := `SELECT id, first_name, phone FROM users WHERE id = $1`
	row := r.DB.QueryRow(ctx, query, userID)

	var user models.User
	err := row.Scan(&user.ID, &user.UserName, &user.Phone)
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}
	return &user, nil
}

func (r *UserRepository) UpdateUserProfile(ctx context.Context, user *models.User) error {
	_, err := r.DB.Exec(ctx, `
	UPDATE users
	SET first_name=$1, email=$2,
	phone_number=$3	WHERE id=$4
	`, user.UserName, user.Email,
		user.Phone, user.ID,
	)
	if err != nil {
		return fmt.Errorf("failed to update user profile: %w", err)
	}
	return nil
}
