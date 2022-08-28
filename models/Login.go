package models

import (
	"time"

	"gorm.io/gorm"
)

type Login struct {
	gorm.Model

	LoginID     string    `gorm:"primaryKey; not null; unique_index"`
	DateLogin   time.Time `gorm:"not null"`
	CustomersID string    `gorm:"not null; unique_index"`
}
