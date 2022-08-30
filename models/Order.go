package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model

	ID             string    `gorm:"primaryKey; not null; unique_index"`
	OrderDate      time.Time `gorm:"not null"`
	RequiredDate   time.Time `gorm:"not null"`
	ShippedDate    time.Time `gorm:"not null"`
	ShipName       string    `gorm:"not null"`
	ShipAddress    string    `gorm:"not null"`
	ShipCity       string    `gorm:"not null"`
	ShipPostalCode string    `gorm:"not null"`
	ShipCountry    string    `gorm:"not null"`
	CustomerID     string    `gorm:"not null; unique_index"`
	ShipperID      string    `gorm:"not null; unique_index"`
	Payments       []Payment
}
