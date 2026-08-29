package port

import (
	"context"

	"github.com/yehezkiel1086/go-prometheus-grafana/internal/core/domain"
)

type OrderRepository interface {
	CreateOrder(ctx context.Context, order *domain.Order) (*domain.Order, error)
	GetOrderByID(ctx context.Context, id uint) (*domain.Order, error)
	GetOrdersByUserID(ctx context.Context, userID uint, limit, offset int) ([]domain.Order, error)
	ListOrders(ctx context.Context, limit, offset int) ([]domain.Order, error)
	UpdateOrderStatus(ctx context.Context, id uint, status domain.OrderStatus) (*domain.Order, error)
	UpdateOrder(ctx context.Context, order *domain.Order) (*domain.Order, error)
	DeleteOrder(ctx context.Context, id uint) error
}

type OrderService interface {
	CreateOrder(ctx context.Context, order *domain.Order) (*domain.Order, error)
	GetOrderByID(ctx context.Context, id uint) (*domain.Order, error)
	GetOrdersByUserID(ctx context.Context, userID uint, limit, offset int) ([]domain.Order, error)
	ListOrders(ctx context.Context, limit, offset int) ([]domain.Order, error)
	UpdateOrderStatus(ctx context.Context, id uint, status domain.OrderStatus) (*domain.Order, error)
	UpdateOrder(ctx context.Context, order *domain.Order) (*domain.Order, error)
	DeleteOrder(ctx context.Context, id uint) error
}
