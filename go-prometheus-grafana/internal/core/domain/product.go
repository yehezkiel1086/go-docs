package domain

import "gorm.io/gorm"

type Product struct {
	gorm.Model

	Name        string  `json:"name" gorm:"size:255;not null"`
	Description string  `json:"description" gorm:"type:text"`
	Price       float64 `json:"price" gorm:"type:decimal(10,2);not null"`
	SKU         string  `json:"sku" gorm:"size:100;not null;unique"`
	Category    string  `json:"category" gorm:"size:100"`
	Active      bool    `json:"active" gorm:"default:true"`

	Stock Stock       `json:"stock,omitempty" gorm:"foreignKey:ProductID"`
	Items []OrderItem `json:"items,omitempty" gorm:"foreignKey:ProductID"`
}
