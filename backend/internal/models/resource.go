package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Resource struct {
	gorm.Model
	ID    uuid.UUID `gorm:"type:uuid"`
	Label string    `gorm:"type:varchar(100);not null" validate:"required,min=5,max=20"`
	Path  string    `gorm:"type:varchar(100);not null" validate:"required,min=5,max=20"`
}
