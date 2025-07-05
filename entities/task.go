package entities

import (
	"time"
)

type Task struct {
	ID        int    `gorm:"primaryKey"`
	Title     string `gorm:"not null"`
	Done      bool   `gorm:"not null"`
	UserEmail string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
