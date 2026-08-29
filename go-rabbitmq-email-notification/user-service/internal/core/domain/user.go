package domain

import (
	"gorm.io/gorm"
)

type Role uint16

const (
	AdminRole Role = 5150
	UserRole  Role = 2001
)

type User struct {
	gorm.Model

	Email             string `json:"email" gorm:"size:255;not null;unique"`
	Password          string `json:"password" gorm:"size:255;not null"`
	Name              string `json:"name" gorm:"size:255;not null"`
	Role              Role   `json:"role" gorm:"default:2001;not null"`
	ConfirmationToken string `json:"confirmation_token" gorm:"size:255"`
	IsVerified        bool   `json:"is_verified" gorm:"default:false"`
}

type UserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
	Name     string `json:"name" binding:"required"`
}

// UserResponse is a DTO — no gorm.Model embedding to avoid exposing
// internal fields (DeletedAt, etc.) and GORM model confusion.
type UserResponse struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
	Role  Role   `json:"role"`
}
