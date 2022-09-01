package models

import "gorm.io/gorm"

type APIError struct {
	gorm.Model
	ErrorCode    uint
	ErrorMessage string
}
