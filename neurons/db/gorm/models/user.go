package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name string
	Age  int
}

type ReadUsers struct {
	ID uint
	Name string
	Age int
}