package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	//gorm.Model
	ID          uint   `gorm:"primaryKey"`
	Email       string `gorm:"unique;not null"`
	PhoneNumber string `gorm:"unique;not null"`
	UserName    string
	PassHash    string `gorm:"not null"`
	AccountID   uint   `gorm:"unique;not null"`
	FirstName   string
	LastName    string
	DOB         time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
