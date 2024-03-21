package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	FirstName string `gorm:"not null"`
	LastName  string `gorm:"not null"`
	Email     string `gorm:"uniqueIndex;not null"`
	Password  string
	Tasks     []Task
}
