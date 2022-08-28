package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model

	CategoryID    string `gorm:"primaryKey; not null; unique_index"`
	CategoryName  string `gorm:"not null; unique_index"`
	Description   string `gorm:"not null"`
	Active        bool   `gorm:"not null; default:true"`
	Subcategories []Subcategory
	Products      []Products
}
