package models

type Task struct {
	ID   int64 `gorm:"primaryKey"`
	Task string
}
