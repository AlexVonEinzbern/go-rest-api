package models

import "gorm.io/gorm"

type Suppliers struct {
	gorm.Model

	SupplierID  string `gorm:"primaryKey; not null; unique_index"`
	CompanyName string `gorm:"not null"`
	Address     string `gorm:"not null"`
	City        string `gorm:"not null"`
	Country     string `gorm:"not null"`
	PostalCode  string `gorm:"not null"`
	Phone       string `gorm:"not null"`
	HomePage    string `gorm:"not null; unique_index"`
	Products    []Products
}
