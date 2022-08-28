package models

import "gorm.io/gorm"

type Subcategory struct {
	gorm.Model

	SubcategoryID string `gorm:"primaryKey; not null; unique_index"`
	CategoryName  string `gorm:"not null; unique_index"`
	Description   string `gorm:"not null"`
	Active        bool   `gorm:"not null; default:true"`
	CategoryID    string `gorm:"not null: unique_index"`
}
