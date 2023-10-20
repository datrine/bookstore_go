package models

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	//gorm.Model
	ID            uint `gorm:"primaryKey"`
	Title         string
	ISBN          string `gorm:"unique"`
	AdderId       uint
	AuthorId      uint
	YearOfPublish int16
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}
