package service

import (
	"context"

	inventory "github.com/yehezkiel1086/go-grpc-inventory-microservices/services/common/protobuf/inventory/protobuf"
	"github.com/yehezkiel1086/go-grpc-inventory-microservices/services/product-service/internal/core/domain"
	"github.com/yehezkiel1086/go-grpc-inventory-microservices/services/product-service/internal/core/port"
)

type ProductService struct {
	productRepo     port.ProductRepository
	inventoryClient inventory.InventoryServiceClient
}

func NewProductService(productRepo port.ProductRepository, inventoryClient inventory.InventoryServiceClient) *ProductService {
	return &ProductService{
		productRepo:     productRepo,
		inventoryClient: inventoryClient,
	}
}

func (s *ProductService) CreateProduct(ctx context.Context, req *domain.CreateProductReq) (*domain.CreateProductRes, error) {
	// create product
	product, err := s.productRepo.CreateProduct(ctx, &domain.Product{
		Name:        req.Name,
		Price:       req.Price,
		Description: req.Description,
	})
	if err != nil {
		return nil, err
	}

	// create inventory
	_, err = s.inventoryClient.CreateInventory(ctx, &inventory.CreateInventoryRequest{
		ProductId: uint64(product.ID),
		Quantity:  int32(req.Quantity),
	})
	if err != nil {
		return nil, err
	}

	return &domain.CreateProductRes{
		ID:          product.ID,
		Name:        product.Name,
		Price:       product.Price,
		Description: product.Description,
		Quantity:    req.Quantity,
	}, nil
}

func (s *ProductService) GetProductByID(ctx context.Context, id uint) (*domain.GetProductRes, error) {
	// get inventory
	inventory, err := s.inventoryClient.GetInventoryByProductID(ctx, &inventory.GetInventoryByProductIDRequest{
		ProductId: uint64(id),
	})
	if err != nil {
		return nil, err
	}

	product, err := s.productRepo.GetProductByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &domain.GetProductRes{
		ID:          product.ID,
		Name:        product.Name,
		Price:       product.Price,
		Description: product.Description,
		Quantity:    int(inventory.Inventory.Quantity),
	}, nil
}

func (s *ProductService) UpdateProduct(ctx context.Context, id uint, req *domain.UpdateProductReq) (*domain.GetProductRes, error) {
	// get product
	product, err := s.productRepo.GetProductByID(ctx, id)
	if err != nil {
		return nil, err
	}

	var inventoryRes *inventory.UpdateInventoryResponse

	// update product
	if req.Name != nil {
		product.Name = *req.Name
	}
	if req.Price != nil {
		product.Price = *req.Price
	}
	if req.Description != nil {
		product.Description = *req.Description
	}
	if req.Quantity != nil {
		// get inventory by product id first
		inv, err := s.inventoryClient.GetInventoryByProductID(ctx, &inventory.GetInventoryByProductIDRequest{
			ProductId: uint64(product.ID),
		})
		if err != nil {
			return nil, err
		}

		// update inventory using the inventory ID
		inventoryRes, err = s.inventoryClient.UpdateInventory(ctx, &inventory.UpdateInventoryRequest{
			Id:       inv.Inventory.Id,
			Quantity: int32(*req.Quantity),
		})
		if err != nil {
			return nil, err
		}
	}

	// update product
	_, err = s.productRepo.UpdateProduct(ctx, product)
	if err != nil {
		return nil, err
	}

	return &domain.GetProductRes{
		ID:          product.ID,
		Name:        product.Name,
		Price:       product.Price,
		Description: product.Description,
		Quantity:    int(inventoryRes.Inventory.Quantity),
	}, nil
}

func (s *ProductService) GetAllProducts(ctx context.Context) ([]domain.GetProductRes, error) {
	// get all products
	products, err := s.productRepo.GetAllProducts(ctx)
	if err != nil {
		return nil, err
	}

	var result []domain.GetProductRes
	for _, product := range products {
		// get inventory
		inventory, err := s.inventoryClient.GetInventoryByProductID(ctx, &inventory.GetInventoryByProductIDRequest{
			ProductId: uint64(product.ID),
		})
		if err != nil {
			return nil, err
		}

		result = append(result, domain.GetProductRes{
			ID:          product.ID,
			Name:        product.Name,
			Price:       product.Price,
			Description: product.Description,
			Quantity:    int(inventory.Inventory.Quantity),
		})
	}
	return result, nil
}

func (s *ProductService) DeleteProduct(ctx context.Context, id uint) error {
	// get inventory by product id first
	inv, err := s.inventoryClient.GetInventoryByProductID(ctx, &inventory.GetInventoryByProductIDRequest{
		ProductId: uint64(id),
	})
	if err != nil {
		return err
	}

	// delete inventory using inventory ID
	if _, err := s.inventoryClient.DeleteInventory(ctx, &inventory.DeleteInventoryRequest{
		Id: inv.Inventory.Id,
	}); err != nil {
		return err
	}

	return s.productRepo.DeleteProduct(ctx, id)
}
