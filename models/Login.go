package models

import (
	"time"

	"gorm.io/gorm"
)

type LoginBase struct {
	DateLogin  time.Time `gorm:"not null"`
	CustomerID string    `gorm:"not null; unique_index"`
}

type Login struct {
	gorm.Model
	ID string `gorm:"primaryKey; not null; unique_index"`
	LoginBase
}

type LoginCreate struct {
	ID string `gorm:"primaryKey; not null; unique_index"`
	LoginBase
}

type LoginResponse struct {
	LoginBase
	ID string
}

// TableName overrides the table name used by Login to `login`
func (Login) TableName() string {
	return "login"
}
