package port

import (
	"context"

	"github.com/yehezkiel1086/go-grpc-inventory-microservices/services/order-service/internal/core/domain"
)

type OrderRepository interface {
	CreateOrder(ctx context.Context, order *domain.Order) (*domain.Order, error)
	GetOrderByID(ctx context.Context, id uint) (*domain.Order, error)
	GetOrdersByUserID(ctx context.Context, userID uint) ([]domain.Order, error)
	UpdateOrder(ctx context.Context, order *domain.Order) (*domain.Order, error)
	DeleteOrder(ctx context.Context, id uint) error
}

type OrderService interface {
	CreateOrder(ctx context.Context, userID uint, req *domain.CreateOrderReq) (*domain.CreateOrderRes, error)
	GetOrderByID(ctx context.Context, id uint) (*domain.GetOrderRes, error)
	GetOrdersByUserID(ctx context.Context, userID uint) ([]domain.GetOrderRes, error)
	UpdateOrder(ctx context.Context, id uint, req *domain.UpdateOrderReq) (*domain.GetOrderRes, error)
	DeleteOrder(ctx context.Context, id uint) error
}
