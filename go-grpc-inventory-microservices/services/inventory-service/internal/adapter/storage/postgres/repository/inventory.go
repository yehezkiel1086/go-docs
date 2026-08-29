package repository

import (
	"context"

	"github.com/yehezkiel1086/go-grpc-inventory-microservices/services/inventory-service/internal/adapter/storage/postgres"
	"github.com/yehezkiel1086/go-grpc-inventory-microservices/services/inventory-service/internal/core/domain"
)

type InventoryRepository struct {
	db *postgres.DB
}

func NewInventoryRepository(db *postgres.DB) *InventoryRepository {
	return &InventoryRepository{
		db: db,
	}
}

func (r *InventoryRepository) CreateInventory(ctx context.Context, inventory *domain.Inventory) (*domain.Inventory, error) {
	db := r.db.GetDB()

	if err := db.WithContext(ctx).Create(inventory).Error; err != nil {
		return nil, err
	}

	return inventory, nil
}

func (r *InventoryRepository) GetInventoryByID(ctx context.Context, id uint) (*domain.Inventory, error) {
	db := r.db.GetDB()

	var inventory domain.Inventory
	if err := db.WithContext(ctx).First(&inventory, id).Error; err != nil {
		return nil, err
	}

	return &inventory, nil
}

func (r *InventoryRepository) GetInventoryByProductID(ctx context.Context, productID uint) (*domain.Inventory, error) {
	db := r.db.GetDB()

	var inventory domain.Inventory
	if err := db.WithContext(ctx).Where("product_id = ?", productID).First(&inventory).Error; err != nil {
		return nil, err
	}

	return &inventory, nil
}

func (r *InventoryRepository) UpdateInventory(ctx context.Context, inventory *domain.Inventory) (*domain.Inventory, error) {
	db := r.db.GetDB()

	if err := db.WithContext(ctx).Save(inventory).Error; err != nil {
		return nil, err
	}

	return inventory, nil
}

func (r *InventoryRepository) DeleteInventory(ctx context.Context, id uint) error {
	db := r.db.GetDB()

	if err := db.WithContext(ctx).Delete(&domain.Inventory{}, id).Error; err != nil {
		return err
	}

	return nil
}
