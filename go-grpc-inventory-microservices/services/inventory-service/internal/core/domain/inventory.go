package domain

import "gorm.io/gorm"

type Inventory struct {
	gorm.Model

	ProductID uint `json:"product_id" gorm:"not null"`
	Quantity  int  `json:"quantity" gorm:"not null"`
}
