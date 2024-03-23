package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	FirstName string `gorm:"not null" json:"first_name" validate:"required"`
	LastName  string `gorm:"not null" json:"last_name" validate:"required"`
	Email     string `gorm:"not null" json:"email" validate:"required,email"`
	Password  string `gorm:"not null" json:"password" validate:"required,min=6,max=20"`
}
