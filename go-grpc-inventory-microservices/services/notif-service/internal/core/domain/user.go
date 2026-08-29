package domain

import "gorm.io/gorm"

type Role uint32

const (
	RoleUser  Role = 2001
	RoleAdmin Role = 5150
)

type User struct {
	gorm.Model

	Email    string `json:"email" gorm:"size:255;unique;not null"`
	Role     Role   `json:"role" gorm:"default:2001"`
	Name     string `json:"name" gorm:"size:255;not null"`
	Password string `json:"password" gorm:"size:255;not null"`
}
