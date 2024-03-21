package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model

	Title       string `gorm:"uniqueIndex;not null"`
	Description string
	Done        bool `gorm:"default:false"`
	UserID      uint
}
