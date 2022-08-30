package models

import (
	"time"

	"gorm.io/gorm"
)

type Login struct {
	gorm.Model

	ID         string    `gorm:"primaryKey; not null; unique_index"`
	DateLogin  time.Time `gorm:"not null"`
	CustomerID string    `gorm:"not null; unique_index"`
}
