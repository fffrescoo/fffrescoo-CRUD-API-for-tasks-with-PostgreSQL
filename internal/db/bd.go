package db

import (
	"fmt"
	"log"
	"pedprojectFinal/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	dsn := "host=localhost user=appuser password=12345 dbname=taskdb port=5432 sslmode=disable"
	log.Println("Connecting to database...")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %w", err)
	}
	log.Println("Connected to database successfully.")

	log.Println("Running auto-migration...")
	err = db.AutoMigrate(&models.Task{})
	if err != nil {
		log.Printf("AutoMigrate error: %v", err)
		return nil, fmt.Errorf("failed to migrate database: %w", err)
	}
	log.Println("Migration completed.")

	return db, nil
}
