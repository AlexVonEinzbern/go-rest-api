package models

import "gorm.io/gorm"

type CategoryBase struct {
	CategoryName string `gorm:"not null; unique_index"`
	Description  string `gorm:"not null"`
	Active       bool   `gorm:"not null; default:true"`
}

type Category struct {
	gorm.Model

	ID string `gorm:"primaryKey; not null; unique_index"`
	CategoryBase
	Subcategories []Subcategory
	Products      []Product
}

type CategoryCreate struct {
	CategoryBase
}

// TableName overrides the table name used by Category to `category`
func (Category) TableName() string {
	return "category"
}
