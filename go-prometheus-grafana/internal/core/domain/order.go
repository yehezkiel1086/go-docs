package domain

import "gorm.io/gorm"

type OrderStatus string

const (
	OrderStatusPending   OrderStatus = "pending"
	OrderStatusPaid      OrderStatus = "paid"
	OrderStatusShipped   OrderStatus = "shipped"
	OrderStatusDelivered OrderStatus = "delivered"
	OrderStatusCancelled OrderStatus = "cancelled"
)

type Order struct {
	gorm.Model

	UserID uint        `json:"user_id" gorm:"not null;index"`
	Status OrderStatus `json:"status" gorm:"size:50;not null;default:'pending'"`
	Total  float64     `json:"total" gorm:"type:decimal(10,2);not null;default:0"`

	User  User        `json:"user,omitempty" gorm:"foreignKey:UserID"`
	Items []OrderItem `json:"items,omitempty" gorm:"foreignKey:OrderID"`
}

type OrderItem struct {
	ID        uint    `json:"id" gorm:"primaryKey"`
	OrderID   uint    `json:"order_id" gorm:"not null;index"`
	ProductID uint    `json:"product_id" gorm:"not null;index"`
	Quantity  int     `json:"quantity" gorm:"not null;default:1"`
	Price     float64 `json:"price" gorm:"type:decimal(10,2);not null"`

	Order   *Order   `json:"order,omitempty" gorm:"foreignKey:OrderID"`
	Product *Product `json:"product,omitempty" gorm:"foreignKey:ProductID"`
}

func (o *Order) CalculateTotal() float64 {
	total := 0.0
	for _, item := range o.Items {
		total += item.Price * float64(item.Quantity)
	}
	o.Total = total
	return total
}
