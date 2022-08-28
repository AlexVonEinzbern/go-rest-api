package models

import "gorm.io/gorm"

type CreditCard struct {
	gorm.Model

	CreditCardID string `gorm:"primaryKey; not null; unique_index"`
	Brand        string `gorm:"not null"`
	Number       string `gorm:"not null; unique_index"`
	CVV          string `gorm:"not null; unique_index"`
	MM           string `gorm:"not null"`
	YYYY         string `gorm:"not null"`
	CustomersID  string `gorm:"not null; unique_index"`
	Payments     []Payment
}
