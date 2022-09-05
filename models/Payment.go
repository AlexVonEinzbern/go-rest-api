package models

import (
	"time"

	"gorm.io/gorm"
)

type PaymentBase struct {
	PaymentDate  time.Time `gorm:"not null"`
	OrderID      string    `gorm:"not null; unique_index"`
	CreditCardID string    `gorm:"not null; unique_index"`
}

type Payment struct {
	gorm.Model
	ID string `gorm:"primaryKey; not null; unique_index"`
	PaymentBase
}

type PaymentCreate struct {
	ID string `gorm:"primaryKey; not null; unique_index"`
	PaymentBase
}

type PaymentResponse struct {
	PaymentBase
	ID           string
	OrderID      string
	CreditCardID string
}

// TableName overrides the table name used by Payment to `payment`
func (Payment) TableName() string {
	return "payment"
}
