package repository

import (
	"context"
	"os"
	"testing"

	// Импортируй pgx и свою тестовую БД
	"github.com/jackc/pgx/v5/pgxpool"
)

// ВНИМАНИЕ: Этот тест работает с реальной БД и только читает данные.
// Не используйте методы, которые что-то меняют!
// Перед запуском укажите существующий ID водителя в переменной realDriverID.

func setupTestDB(t *testing.T) *pgxpool.Pool {
	dbURL := os.Getenv("TEST_DB_URL") // например, "postgres://user:pass@host:port/dbname?sslmode=disable"
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
	realDriverID := "PUT_REAL_DRIVER_ID_HERE"

	driver, err := repo.GetDriverByID(context.Background(), realDriverID)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if driver.ID != realDriverID {
		t.Errorf("unexpected driver: %+v", driver)
	}
}
