package repository

import (
	"context"
	"testing"

	// Импортируй pgx и свою тестовую БД
	"github.com/GameXost/YandexGo_proj/DRIVERS/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

// ВНИМАНИЕ: Этот тест работает с реальной БД и только читает данные.
// Не используйте методы, которые что-то меняют!
// Перед запуском укажите существующий ID водителя в переменной realDriverID.

func setupTestDB(t *testing.T) *pgxpool.Pool {
	dbURL := "postgres://gamexost:gopython@localhost:5432/postgres?sslmode=disable"
	pool, err := pgxpool.New(context.Background(), dbURL)
	if err != nil {
		t.Fatalf("failed to connect to db: %v", err)
	}
	return pool
}

func TestGetDriverByID_RealDB(t *testing.T) {
	db := setupTestDB(t)
	repo := NewDriverRepository(db)

	// Укажи здесь реально существующий ID водителя из своей БД
	realDriverID := "1"

	driver, err := repo.GetDriverByID(context.Background(), realDriverID)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if driver.ID != realDriverID {
		t.Errorf("unexpected driver: %+v", driver)
	}
}

func TestUpdateDriverProfile_RealDB(t *testing.T) {
	db := setupTestDB(t)
	repo := NewDriverRepository(db)

	// Подставь существующий ID и новые значения
	driver := &models.Driver{
		ID:         "1",
		UserName:   "newshit",
		Email:      "updated@example.com",
		Phone:      "12344134",
		Car_number: "B456CD",
		Car_model:  "Honda",
		Car_marks:  "Civic",
		Car_color:  "Black",
	}
	err := repo.UpdateDriverProfile(context.Background(), driver)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Проверь, что данные реально обновились
	got, err := repo.GetDriverByID(context.Background(), "1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got.UserName != "newshit" {
		t.Errorf("profile not updated: %+v", got)
	}
}
