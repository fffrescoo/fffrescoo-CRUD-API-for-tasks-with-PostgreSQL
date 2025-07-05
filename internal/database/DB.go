package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func InitDB() (*gorm.DB, error) {
	dsn := "host=localhost user=appuser password=12345 dbname=taskdb port=5432 sslmode=disable"
	log.Println("Connecting to database...")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %w", err)
	}
	log.Println("Connected to database successfully.")
	DB = db
	return db, nil
}
