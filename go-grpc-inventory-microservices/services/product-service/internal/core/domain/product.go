package domain

import "gorm.io/gorm"

type Product struct {
	gorm.Model

	Name        string  `json:"name" gorm:"size:255;not null;unique"`
	Price       float64 `json:"price" gorm:"not null"`
	Description string  `json:"description" gorm:"type:text"`
}

type CreateProductReq struct {
	Name        string  `json:"name" binding:"required"`
	Price       float64 `json:"price" binding:"required,gt=0"`
	Description string  `json:"description"`
	Quantity    int     `json:"quantity" binding:"required,gt=0"`
}

type CreateProductRes struct {
	ID          uint    `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Quantity    int     `json:"quantity"`
}

type GetProductRes struct {
	ID          uint    `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Quantity    int     `json:"quantity"`
}

type UpdateProductReq struct {
	Name        *string  `json:"name"`
	Price       *float64 `json:"price"`
	Description *string  `json:"description"`
	Quantity    *int     `json:"quantity"`
}
