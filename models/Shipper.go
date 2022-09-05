package models

import "gorm.io/gorm"

type ShipperBase struct {
	CompanyName string `gorm:"not null; unique_index"`
	Phone       string `gorm:"not null"`
}

type Shipper struct {
	gorm.Model
	ID string `gorm:"primaryKey; not null; unique_index"`
	ShipperBase
	Orders []Order
}

type ShipperCreate struct {
	ID string `gorm:"primaryKey; not null; unique_index"`
	ShipperBase
}

type ShipperResponse struct {
	ShipperBase
	ID string
}

// TableName overrides the table name used by Shipper to `shipper`
func (Shipper) TableName() string {
	return "shipper"
}
