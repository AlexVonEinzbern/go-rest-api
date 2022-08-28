package models

import "gorm.io/gorm"

type OrderProduct struct {
	gorm.Model

	ProductID string  `gorm:"primaryKey; not null"`
	OrderID   string  `gorm:"primaryKey; not null"`
	UnitPrice float32 `gorm:"not null"`
	Quantity  uint    `gorm:"not null"`
	Discount  float32 `gorm:"not null"`
}
