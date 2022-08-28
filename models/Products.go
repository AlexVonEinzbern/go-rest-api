package models

import "gorm.io/gorm"

type Products struct {
	gorm.Model

	ProductID       string   `gorm:"primaryKey; not null; unique_index"`
	ProductName     string   `gorm:"not null"`
	QuantityPerUnit string   `gorm:"not null"`
	UnitsInStock    uint     `gorm:"not null"`
	UnitsOnOrder    uint     `gorm:"not null"`
	ReorderLevel    uint     `gorm:"not null"`
	Discontinued    bool     `gorm:"default:false"`
	Quantity        uint     `gorm:"not null"`
	SupplierID      string   `gorm:"not null; unique_index"`
	CategoryID      string   `gorm:"not null; unique_index"`
	Orders          []Orders `gorm:"many2many:order_product;"`
}
