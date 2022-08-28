package models

import (
	"time"

	"gorm.io/gorm"
)

type Customers struct {
	gorm.Model

	CustomersID string    `gorm:"primaryKey; not null; unique_index"`
	Name        string    `gorm:"not null"`
	Email       string    `gorm:"not null; unique_index"`
	DNI         string    `gorm:"not null; unique_index"`
	Address     string    `gorm:"not null"`
	City        string    `gorm:"not null"`
	Birthday    time.Time `gorm:"not null"`
	PostalCode  string    `gorm:"not null"`
	Country     string    `gorm:"not null"`
	Phone       string    `gorm:"not null"`
	UserName    string    `gorm:"not null"`
	Password    string    `gorm:"not null; unique_index"`
	CreditCards []CreditCard
	Logins      []Login
	Orders      []Orders
}
