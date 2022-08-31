package models

import "gorm.io/gorm"

type CreditCard struct {
	gorm.Model

	ID         string `gorm:"primaryKey; not null; unique_index"`
	Brand      string `gorm:"not null"`
	Number     string `gorm:"not null; unique_index"`
	CVV        string `gorm:"not null; unique_index"`
	MM         string `gorm:"not null"`
	YYYY       string `gorm:"not null"`
	CustomerID string `gorm:"not null; unique_index"`
	Payments   []Payment
}

// TableName overrides the table name used by CredirCard to `credit_card`
func (CreditCard) TableName() string {
	return "credit_card"
}