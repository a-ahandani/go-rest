package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Note struct {
	gorm.Model
	ID       uuid.UUID `gorm:"type:uuid"`
	Title    string    `gorm:"type:varchar(100);not null" validate:"required,min=5,max=20"`
	Subtitle string    `gorm:"type:varchar(100);not null" validate:"required,min=5,max=20"`
	Text     string    `gorm:"type:varchar(100);not null" validate:"required,min=5,max=20"`
}
