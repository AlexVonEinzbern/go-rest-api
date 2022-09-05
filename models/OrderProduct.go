package models

import "gorm.io/gorm"

type OrderProductBase struct {
	UnitPrice float32 `gorm:"not null"`
	Quantity  uint    `gorm:"not null"`
	Discount  float32 `gorm:"not null"`
	ProductID string  `gorm:"primaryKey; not null"`
	OrderID   string  `gorm:"primaryKey; not null"`
}

type OrderProduct struct {
	gorm.Model

	OrderProductBase
}

type OrderProductCreate struct {
	OrderProductBase
}

type OrderProductResponse struct {
	OrderProductBase
	ProductID string
	OrderID   string
}

// TableName overrides the table name used by OrderProduct to `order_product`
func (OrderProduct) TableName() string {
	return "order_product"
}
