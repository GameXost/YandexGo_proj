package repository

import (
	"context"
	"fmt"

	"github.com/GameXost/YandexGo_proj/DRIVERS/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DriverRepositoryInterface interface {
	GetDriverByID(ctx context.Context, driverID string) (*models.Driver, error)
	UpdateDriverProfile(ctx context.Context, driver *models.Driver) error
	GetUserByID(ctx context.Context, userID string) (*models.User, error)
}

type DriverRepository struct {
	DB *pgxpool.Pool
}

func NewDriverRepository(db *pgxpool.Pool) *DriverRepository {
	return &DriverRepository{
		DB: db,
	}
}

func (r *DriverRepository) GetDriverByID(ctx context.Context, driverID string) (*models.Driver, error) {
	query := `SELECT id, first_name, email, phone_number,car_number, car_model, car_marks, car_color FROM drivers WHERE id = $1`
	row := r.DB.QueryRow(ctx, query, driverID)

	var driver models.Driver
	err := row.Scan(
		&driver.ID, &driver.UserName, &driver.Email,
		&driver.Phone,
		&driver.Car_number, &driver.Car_model, &driver.Car_marks, &driver.Car_color,
	)
	if err != nil {
		return nil, fmt.Errorf("driver not found: %w", err)
	}
	return &driver, nil
}

// sqlbuilder
func (r *DriverRepository) UpdateDriverProfile(ctx context.Context, driver *models.Driver) error {
	_, err := r.DB.Exec(ctx, `
	UPDATE drivers
	SET
	first_name=$1, email=$2, phone_number=$3,
	car_number=$4, car_model=$5,
	car_marks=$6, car_color=$7 WHERE id=$8	`,
		driver.UserName, driver.Email, driver.Phone,
		driver.Car_number, driver.Car_model,
		driver.Car_marks, driver.Car_color, driver.ID,
	)
	if err != nil {
		return fmt.Errorf("failed to update driver profile: %w", err)
	}
	return nil
}

// под вопросом, т.к кафка это всё передает
func (r *DriverRepository) GetUserByID(ctx context.Context, userID string) (*models.User, error) {
	query := `SELECT id, username, phone FROM users WHERE id = $1`
	row := r.DB.QueryRow(ctx, query, userID)

	var user models.User
	err := row.Scan(&user.ID, &user.Username, &user.Phone)
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}
	return &user, nil
}
