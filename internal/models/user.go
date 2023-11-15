package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uuid.UUID `gorm:"type:uuid"`
	Name     string    `gorm:"type:varchar(100);not null" validate:"required,min=5,max=20"`
	Email    string    `gorm:"type:varchar(100);uniqueIndex;not null" validate:"required,email"`
	Password string    `gorm:"type:varchar(100);not null" validate:"required,min=6"`
	Verified *bool     `gorm:"not null;default:false"`
}
