package infra

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializePostgresDB() (*gorm.DB, error) {
	dsn := os.Getenv("DATABASE_URL")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	println("[DB] Postgres initialize successfully")

	return db, nil
}
