package models

import "gorm.io/gorm"

type Shipper struct {
	gorm.Model

	ID          string `gorm:"primaryKey; not null; unique_index"`
	CompanyName string `gorm:"not null; unique_index"`
	Phone       string `gorm:"not null"`
	Orders      []Order
}
