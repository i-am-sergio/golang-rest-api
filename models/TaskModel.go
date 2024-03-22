package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model

	ID          uint   `gorm:"primaryKey;autoIncrement;not null;unique"`
	Title       string `gorm:"not null"`
	Description string
	Done        bool `gorm:"default:false"`
	UserID      uint
}
