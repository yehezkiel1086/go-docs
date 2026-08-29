package repository

import (
	"context"

	"github.com/yehezkiel1086/go-prometheus-grafana/internal/adapter/storage/postgres"
	"github.com/yehezkiel1086/go-prometheus-grafana/internal/core/domain"
)

type OrderRepository struct {
	db *postgres.DB
}

func NewOrderRepository(db *postgres.DB) *OrderRepository {
	return &OrderRepository{db}
}

func (r *OrderRepository) CreateOrder(ctx context.Context, order *domain.Order) (*domain.Order, error) {
	db := r.db.GetDB()

	if err := db.WithContext(ctx).Create(order).Error; err != nil {
		return nil, err
	}

	return order, nil
}

func (r *OrderRepository) GetOrderByID(ctx context.Context, id uint) (*domain.Order, error) {
	var order domain.Order

	db := r.db.GetDB()

	if err := db.WithContext(ctx).Preload("Items").Preload("Items.Product").Preload("User").First(&order, id).Error; err != nil {
		return nil, err
	}

	return &order, nil
}

func (r *OrderRepository) GetOrdersByUserID(ctx context.Context, userID uint, limit, offset int) ([]domain.Order, error) {
	var orders []domain.Order

	db := r.db.GetDB()

	if err := db.WithContext(ctx).Where("user_id = ?", userID).Limit(limit).Offset(offset).Preload("Items").Find(&orders).Error; err != nil {
		return nil, err
	}

	return orders, nil
}

func (r *OrderRepository) ListOrders(ctx context.Context, limit, offset int) ([]domain.Order, error) {
	var orders []domain.Order

	db := r.db.GetDB()

	if err := db.WithContext(ctx).Limit(limit).Offset(offset).Preload("Items").Preload("User").Find(&orders).Error; err != nil {
		return nil, err
	}

	return orders, nil
}

func (r *OrderRepository) UpdateOrderStatus(ctx context.Context, id uint, status domain.OrderStatus) (*domain.Order, error) {
	db := r.db.GetDB()

	if err := db.WithContext(ctx).Model(&domain.Order{}).Where("id = ?", id).Update("status", status).Error; err != nil {
		return nil, err
	}

	return r.GetOrderByID(ctx, id)
}

func (r *OrderRepository) UpdateOrder(ctx context.Context, order *domain.Order) (*domain.Order, error) {
	db := r.db.GetDB()

	if err := db.WithContext(ctx).Save(order).Error; err != nil {
		return nil, err
	}

	return order, nil
}

func (r *OrderRepository) DeleteOrder(ctx context.Context, id uint) error {
	db := r.db.GetDB()

	if err := db.WithContext(ctx).Delete(&domain.Order{}, id).Error; err != nil {
		return err
	}

	return nil
}
