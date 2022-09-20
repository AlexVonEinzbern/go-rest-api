package models

import (
	"gorm.io/gorm"
)

type PaymentBase struct {
	OrderID      string `gorm:"not null; unique_index"`
	CreditCardID string `gorm:"not null; unique_index"`
}

type Payment struct {
	gorm.Model
	ID string `gorm:"primaryKey; not null; unique_index"`
	PaymentBase
}

type PaymentCreate struct {
	PaymentBase
}

// TableName overrides the table name used by Payment to `payment`
func (Payment) TableName() string {
	return "payment"
}
