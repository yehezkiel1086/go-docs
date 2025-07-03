package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID `json:"id" binding:"required" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Username string `json:"username" binding:"required" gorm:"size:255;not null;unique"`
	Password string `json:"password" binding:"required" gorm:"size:255;not null"`

	RoleID uuid.UUID `json:"role_id" binding:"required" gorm:"type:uuid"`
	UserRole Role	`json:"-" binding:"-" gorm:"foreignKey:RoleID"`
}

type Role struct {
	ID        uuid.UUID `json:"id" binding:"required" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Role string `json:"role" binding:"required" gorm:"size:255;not null;unique"`
}

type UserInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	RoleID uuid.UUID `json:"role_id" binding:"required" gorm:"type:uuid"`
}
