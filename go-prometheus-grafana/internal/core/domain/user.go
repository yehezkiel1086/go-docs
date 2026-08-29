package domain

import "gorm.io/gorm"

type Role uint32

const (
	UserRole Role = 2001
	AdminRole Role = 5150
)

type User struct {
	gorm.Model

	Name string `json:"name" gorm:"size:255;not null"`
	Email string `json:"email" gorm:"size:255;not null;unique"`
	Password string `json:"password" gorm:"size:255;not null"`
	Role Role `json:"role" gorm:"default:2001"`
}
