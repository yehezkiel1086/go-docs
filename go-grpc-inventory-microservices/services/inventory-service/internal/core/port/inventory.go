package port

import (
	"context"

	"github.com/yehezkiel1086/go-grpc-inventory-microservices/services/inventory-service/internal/core/domain"
)

type InventoryRepository interface {
	CreateInventory(ctx context.Context, inventory *domain.Inventory) (*domain.Inventory, error)
	GetInventoryByID(ctx context.Context, id uint) (*domain.Inventory, error)
	GetInventoryByProductID(ctx context.Context, productID uint) (*domain.Inventory, error)
	UpdateInventory(ctx context.Context, inventory *domain.Inventory) (*domain.Inventory, error)
	DeleteInventory(ctx context.Context, id uint) error
}

type InventoryService interface {
	CreateInventory(ctx context.Context, inventory *domain.Inventory) (*domain.Inventory, error)
	GetInventoryByID(ctx context.Context, id uint) (*domain.Inventory, error)
	GetInventoryByProductID(ctx context.Context, productID uint) (*domain.Inventory, error)
	UpdateInventory(ctx context.Context, id uint, req *domain.Inventory) (*domain.Inventory, error)
	DeleteInventory(ctx context.Context, id uint) error
}
