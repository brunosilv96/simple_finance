package infra

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializePostgresDB() (*gorm.DB, error) {
	dsn := "postgresql://postgres:123456@localhost:5432/simple_finance"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	println("[DB] Postgres initialize successfully")

	return db, nil
}
