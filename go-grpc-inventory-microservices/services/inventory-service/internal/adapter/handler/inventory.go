package handler

import (
	"context"

	go_grpc_inventory_microservices "github.com/yehezkiel1086/go-grpc-inventory-microservices/services/common/protobuf/inventory/protobuf"
	"github.com/yehezkiel1086/go-grpc-inventory-microservices/services/inventory-service/internal/core/domain"
	"github.com/yehezkiel1086/go-grpc-inventory-microservices/services/inventory-service/internal/core/port"
	"google.golang.org/grpc"
)

type InventoryHandler struct {
	svc port.InventoryService
	go_grpc_inventory_microservices.UnimplementedInventoryServiceServer
}

func NewInventoryHandler(grpc *grpc.Server, svc port.InventoryService) {
	inventoryHandler := &InventoryHandler{
		svc: svc,
	}

	go_grpc_inventory_microservices.RegisterInventoryServiceServer(grpc, inventoryHandler)
}

func (h *InventoryHandler) CreateInventory(ctx context.Context, req *go_grpc_inventory_microservices.CreateInventoryRequest) (*go_grpc_inventory_microservices.CreateInventoryResponse, error) {
	inventory, err := h.svc.CreateInventory(ctx, &domain.Inventory{
		ProductID: uint(req.ProductId),
		Quantity:  int(req.Quantity),
	})
	if err != nil {
		return nil, err
	}

	return &go_grpc_inventory_microservices.CreateInventoryResponse{
		Inventory: &go_grpc_inventory_microservices.Inventory{
			Id:        uint64(inventory.ID),
			ProductId: uint64(inventory.ProductID),
			Quantity:  int32(inventory.Quantity),
			CreatedAt: inventory.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: inventory.UpdatedAt.Format("2006-01-02 15:04:05"),
		},
	}, nil
}

func (h *InventoryHandler) GetInventory(ctx context.Context, req *go_grpc_inventory_microservices.GetInventoryRequest) (*go_grpc_inventory_microservices.GetInventoryResponse, error) {
	inventory, err := h.svc.GetInventoryByID(ctx, uint(req.Id))
	if err != nil {
		return nil, err
	}

	return &go_grpc_inventory_microservices.GetInventoryResponse{
		Inventory: &go_grpc_inventory_microservices.Inventory{
			Id:        uint64(inventory.ID),
			ProductId: uint64(inventory.ProductID),
			Quantity:  int32(inventory.Quantity),
			CreatedAt: inventory.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: inventory.UpdatedAt.Format("2006-01-02 15:04:05"),
		},
	}, nil
}

func (h *InventoryHandler) GetInventoryByProductID(ctx context.Context, req *go_grpc_inventory_microservices.GetInventoryByProductIDRequest) (*go_grpc_inventory_microservices.GetInventoryByProductIDResponse, error) {
	inventory, err := h.svc.GetInventoryByProductID(ctx, uint(req.ProductId))
	if err != nil {
		return nil, err
	}

	return &go_grpc_inventory_microservices.GetInventoryByProductIDResponse{
		Inventory: &go_grpc_inventory_microservices.Inventory{
			Id:        uint64(inventory.ID),
			ProductId: uint64(inventory.ProductID),
			Quantity:  int32(inventory.Quantity),
			CreatedAt: inventory.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: inventory.UpdatedAt.Format("2006-01-02 15:04:05"),
		},
	}, nil
}

func (h *InventoryHandler) UpdateInventory(ctx context.Context, req *go_grpc_inventory_microservices.UpdateInventoryRequest) (*go_grpc_inventory_microservices.UpdateInventoryResponse, error) {
	inventory, err := h.svc.UpdateInventory(ctx, uint(req.Id), &domain.Inventory{
		ProductID: uint(req.ProductId),
		Quantity:  int(req.Quantity),
	})
	if err != nil {
		return nil, err
	}

	return &go_grpc_inventory_microservices.UpdateInventoryResponse{
		Inventory: &go_grpc_inventory_microservices.Inventory{
			Id:        uint64(inventory.ID),
			ProductId: uint64(inventory.ProductID),
			Quantity:  int32(inventory.Quantity),
			CreatedAt: inventory.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: inventory.UpdatedAt.Format("2006-01-02 15:04:05"),
		},
	}, nil
}

func (h *InventoryHandler) DeleteInventory(ctx context.Context, req *go_grpc_inventory_microservices.DeleteInventoryRequest) (*go_grpc_inventory_microservices.DeleteInventoryResponse, error) {
	err := h.svc.DeleteInventory(ctx, uint(req.Id))
	if err != nil {
		return &go_grpc_inventory_microservices.DeleteInventoryResponse{Success: false}, err
	}

	return &go_grpc_inventory_microservices.DeleteInventoryResponse{Success: true}, nil
}
