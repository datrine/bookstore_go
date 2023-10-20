package models

import (
	"time"

	"gorm.io/gorm"
)

type Account struct {
	//gorm.Model
	ID        uint   `gorm:"primaryKey"`
	Role      string `gorm:"default:'USER'"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	User      []User         `gorm:"constraint:OnUpdate:CASCADE,OnDelete: CASCADE;"`
}
