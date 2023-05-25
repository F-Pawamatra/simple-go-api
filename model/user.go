package model

import "gorm.io/gorm"

type (
	// User represents the structure of our resource
	User struct {
		gorm.Model
		Name  string `json:"name"`
		Email string `json:"email" gorm:"unique" validate:"required,email"`
	}
)
