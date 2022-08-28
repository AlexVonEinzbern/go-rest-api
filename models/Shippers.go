package models

import "gorm.io/gorm"

type Shippers struct {
	gorm.Model

	ShipperID   string `gorm:"primaryKey; not null; unique_index"`
	CompanyName string `gorm:"not null; unique_index"`
	Phone       string `gorm:"not null"`
	Orders      []Orders
}
