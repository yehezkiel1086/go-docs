package domain

import "gorm.io/gorm"

type OrderStatus string

const (
	OrderStatusPending   OrderStatus = "pending"
	OrderStatusConfirmed OrderStatus = "confirmed"
	OrderStatusShipped   OrderStatus = "shipped"
	OrderStatusDelivered OrderStatus = "delivered"
	OrderStatusCancelled OrderStatus = "cancelled"
)

type Order struct {
	gorm.Model

	UserID     uint        `json:"user_id" gorm:"not null"`
	ProductID  uint        `json:"product_id" gorm:"not null"`
	Quantity   int         `json:"quantity" gorm:"not null"`
	TotalPrice float64     `json:"total_price" gorm:"not null"`
	Status     OrderStatus `json:"status" gorm:"size:20;default:'pending'"`
}

type CreateOrderReq struct {
	ProductID uint `json:"product_id" binding:"required"`
	Quantity  int  `json:"quantity" binding:"required,gt=0"`
}

type CreateOrderRes struct {
	ID         uint        `json:"id"`
	UserID     uint        `json:"user_id"`
	ProductID  uint        `json:"product_id"`
	Quantity   int         `json:"quantity"`
	TotalPrice float64     `json:"total_price"`
	Status     OrderStatus `json:"status"`
}

type GetOrderRes struct {
	ID         uint        `json:"id"`
	UserID     uint        `json:"user_id"`
	ProductID  uint        `json:"product_id"`
	Quantity   int         `json:"quantity"`
	TotalPrice float64     `json:"total_price"`
	Status     OrderStatus `json:"status"`
}

type UpdateOrderReq struct {
	Status *OrderStatus `json:"status"`
}
