package models

import (
	"time"

	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model

	ID           string    `gorm:"primaryKey; not null; unique_index"`
	PaymentDate  time.Time `gorm:"not null"`
	OrderID      string    `gorm:"not null; unique_index"`
	CreditCardID string    `gorm:"not null; unique_index"`
}
