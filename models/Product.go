package models

import "gorm.io/gorm"

type ProductBase struct {
	ProductName     string `gorm:"not null"`
	QuantityPerUnit string `gorm:"not null"`
	UnitsInStock    uint   `gorm:"not null"`
	UnitsOnOrder    uint   `gorm:"not null"`
	ReorderLevel    uint   `gorm:"not null"`
	Discontinued    bool   `gorm:"default:false"`
	Quantity        uint   `gorm:"not null"`
	SupplierID      string `gorm:"not null; unique_index"`
	CategoryID      string `gorm:"not null; unique_index"`
}

type Product struct {
	gorm.Model
	ID string `gorm:"primaryKey; not null; unique_index"`
	ProductBase
	Orders []Order `gorm:"many2many:order_product;"`
}

type ProductCreate struct {
	ProductBase
}

// TableName overrides the table name used by Product to `product`
func (Product) TableName() string {
	return "product"
}
