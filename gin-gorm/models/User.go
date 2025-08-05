package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name string `json:"name" binding:"required"`
	Age uint8 `json:"age" binding:"required"`
}
