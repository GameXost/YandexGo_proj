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

func (r *DriverRepository) UpdateDriverProfile(ctx context.Context, driver *models.Driver) error {
	_, err := r.DB.Exec(ctx, `
	UPDATE drivers
	SET
	username=$1, email=$2, phone=$3,
	car=$4, car_number=$5, car_model=$6,
	car_marks=$7, car_color=$8,	WHERE id=$9	`,
		driver.UserName, driver.Email, driver.Phone,
		driver.Car, driver.Car_number, driver.Car_model,
		driver.Car_marks, driver.Car_color, driver.ID,
	)
	if err != nil {
		return fmt.Errorf("failed to update driver profile: %w", err)
	}
	return nil
}
