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

// TableName overrides the table name used by Login to `login`
func (Login) TableName() string {
	return "login"
}