package repository

import (
	"context"

	"github.com/yehezkiel1086/go-grpc-inventory-microservices/services/order-service/internal/adapter/storage/postgres"
	"github.com/yehezkiel1086/go-grpc-inventory-microservices/services/order-service/internal/core/domain"
)

type OrderRepository struct {
	db *postgres.DB
}

func NewOrderRepository(db *postgres.DB) *OrderRepository {
	return &OrderRepository{
		db: db,
	}
}

func (r *OrderRepository) CreateOrder(ctx context.Context, order *domain.Order) (*domain.Order, error) {
	db := r.db.GetDB()

	if err := db.WithContext(ctx).Create(order).Error; err != nil {
		return nil, err
	}

	return order, nil
}

func (r *OrderRepository) GetOrderByID(ctx context.Context, id uint) (*domain.Order, error) {
	db := r.db.GetDB()

	var order domain.Order
	if err := db.WithContext(ctx).First(&order, id).Error; err != nil {
		return nil, err
	}

	return &order, nil
}

func (r *OrderRepository) GetOrdersByUserID(ctx context.Context, userID uint) ([]domain.Order, error) {
	db := r.db.GetDB()

	var orders []domain.Order
	if err := db.WithContext(ctx).Where("user_id = ?", userID).Find(&orders).Error; err != nil {
		return nil, err
	}

	return orders, nil
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
