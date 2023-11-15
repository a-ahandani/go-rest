package models

import (
	"github.com/golang-jwt/jwt"
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

// TokenClaims represents the claims in the JWT token
type TokenClaims struct {
	UserID uuid.UUID `json:"user_id"`
	jwt.StandardClaims
}
