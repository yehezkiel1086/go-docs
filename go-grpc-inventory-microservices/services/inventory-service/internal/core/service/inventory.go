package service

import (
	"context"

	"github.com/yehezkiel1086/go-grpc-inventory-microservices/services/inventory-service/internal/core/domain"
	"github.com/yehezkiel1086/go-grpc-inventory-microservices/services/inventory-service/internal/core/port"
)

type InventoryService struct {
	repo port.InventoryRepository
}

func NewInventoryService(repo port.InventoryRepository) *InventoryService {
	return &InventoryService{
		repo,
	}
}

func (s *InventoryService) CreateInventory(ctx context.Context, inventory *domain.Inventory) (*domain.Inventory, error) {
	return s.repo.CreateInventory(ctx, inventory)
}

func (s *InventoryService) GetInventoryByID(ctx context.Context, id uint) (*domain.Inventory, error) {
	return s.repo.GetInventoryByID(ctx, id)
}

func (s *InventoryService) GetInventoryByProductID(ctx context.Context, productID uint) (*domain.Inventory, error) {
	return s.repo.GetInventoryByProductID(ctx, productID)
}

func (s *InventoryService) UpdateInventory(ctx context.Context, id uint, req *domain.Inventory) (*domain.Inventory, error) {
	// get inventory
	inventory, err := s.repo.GetInventoryByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// update fields if provided
	if req.ProductID != 0 {
		inventory.ProductID = req.ProductID
	}
	if req.Quantity != 0 {
		inventory.Quantity = req.Quantity
	}

	return s.repo.UpdateInventory(ctx, inventory)
}

func (s *InventoryService) DeleteInventory(ctx context.Context, id uint) error {
	return s.repo.DeleteInventory(ctx, id)
}
