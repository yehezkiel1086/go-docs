package domain

import (
	"gorm.io/gorm"
)

type Notification struct {
	gorm.Model

	UserID  uint   `json:"user_id" gorm:"not null;index"`
	OrderID *uint  `json:"order_id,omitempty" gorm:"index"`
	Message string `json:"message" gorm:"not null"`
	Type    string `json:"type" gorm:"size:50;not null"` // e.g., "order_created", "order_shipped", "payment_received"
	IsRead  bool   `json:"is_read" gorm:"default:false"`
}
