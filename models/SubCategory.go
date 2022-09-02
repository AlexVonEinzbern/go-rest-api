package models

import "gorm.io/gorm"

type SubcategoryBase struct {
	SubCategoryName string `gorm:"not null; unique_index"`
	Description     string `gorm:"not null"`
	Active          bool   `gorm:"not null; default:true"`
	CategoryID      string `gorm:"not null: unique_index"`
}

type Subcategory struct {
	gorm.Model

	ID string `gorm:"primaryKey; not null; unique_index"`
	SubcategoryBase
}

type SubcategoryCreate struct {
	ID string `gorm:"primaryKey; not null; unique_index"`
	SubcategoryBase
}

type SubcategoryResponse struct {
	SubcategoryBase
	ID string //`gorm:"primaryKey"`
}

// TableName overrides the table name used by Subcategory to `subcategory`
func (Subcategory) TableName() string {
	return "subcategory"
}
