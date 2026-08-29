package service

import (
	"context"

	"github.com/yehezkiel1086/go-prometheus-grafana/internal/core/domain"
	"github.com/yehezkiel1086/go-prometheus-grafana/internal/core/port"
)

type OrderService struct {
	repo port.OrderRepository
}

func NewOrderService(repo port.OrderRepository) *OrderService {
	return &OrderService{
		repo,
	}
}

func (s *OrderService) CreateOrder(ctx context.Context, order *domain.Order) (*domain.Order, error) {
	return s.repo.CreateOrder(ctx, order)
}

func (s *OrderService) GetOrderByID(ctx context.Context, id uint) (*domain.Order, error) {
	return s.repo.GetOrderByID(ctx, id)
}

func (s *OrderService) GetOrdersByUserID(ctx context.Context, userID uint, limit, offset int) ([]domain.Order, error) {
	return s.repo.GetOrdersByUserID(ctx, userID, limit, offset)
}

func (s *OrderService) ListOrders(ctx context.Context, limit, offset int) ([]domain.Order, error) {
	return s.repo.ListOrders(ctx, limit, offset)
}

func (s *OrderService) UpdateOrderStatus(ctx context.Context, id uint, status domain.OrderStatus) (*domain.Order, error) {
	return s.repo.UpdateOrderStatus(ctx, id, status)
}

func (s *OrderService) UpdateOrder(ctx context.Context, order *domain.Order) (*domain.Order, error) {
	return s.repo.UpdateOrder(ctx, order)
}

func (s *OrderService) DeleteOrder(ctx context.Context, id uint) error {
	return s.repo.DeleteOrder(ctx, id)
}
