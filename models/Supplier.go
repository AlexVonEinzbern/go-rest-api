package models

import "gorm.io/gorm"

type Supplier struct {
	gorm.Model

	ID          string `gorm:"primaryKey; not null; unique_index"`
	CompanyName string `gorm:"not null"`
	Address     string `gorm:"not null"`
	City        string `gorm:"not null"`
	Country     string `gorm:"not null"`
	PostalCode  string `gorm:"not null"`
	Phone       string `gorm:"not null"`
	HomePage    string `gorm:"not null; unique_index"`
	Products    []Product
}

// TableName overrides the table name used by Supplier to `supplier`
func (Supplier) TableName() string {
	return "supplier"
}