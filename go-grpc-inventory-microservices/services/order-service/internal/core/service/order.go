package service

import (
	"context"
	"errors"
	"fmt"

	inventory "github.com/yehezkiel1086/go-grpc-inventory-microservices/services/common/protobuf/inventory/protobuf"
	product "github.com/yehezkiel1086/go-grpc-inventory-microservices/services/common/protobuf/product/protobuf"
	"github.com/yehezkiel1086/go-grpc-inventory-microservices/services/order-service/internal/adapter/storage/rabbitmq"
	"github.com/yehezkiel1086/go-grpc-inventory-microservices/services/order-service/internal/core/domain"
	"github.com/yehezkiel1086/go-grpc-inventory-microservices/services/order-service/internal/core/port"
)

type OrderService struct {
	mq              *rabbitmq.RabbitMQ
	repo            port.OrderRepository
	productClient   product.ProductServiceClient
	inventoryClient inventory.InventoryServiceClient
}

func NewOrderService(repo port.OrderRepository, productClient product.ProductServiceClient, inventoryClient inventory.InventoryServiceClient, mq *rabbitmq.RabbitMQ) *OrderService {
	return &OrderService{
		repo:            repo,
		productClient:   productClient,
		inventoryClient: inventoryClient,
		mq:              mq,
	}
}

func (s *OrderService) CreateOrder(ctx context.Context, userID uint, req *domain.CreateOrderReq) (*domain.CreateOrderRes, error) {
	// Get product details from product service
	productResp, err := s.productClient.GetProduct(ctx, &product.GetProductRequest{
		Id: uint64(req.ProductID),
	})
	if err != nil {
		return nil, err
	}

	// Get inventory by product ID
	inventoryResp, err := s.inventoryClient.GetInventoryByProductID(ctx, &inventory.GetInventoryByProductIDRequest{
		ProductId: uint64(req.ProductID),
	})
	if err != nil {
		return nil, err
	}

	// Check if inventory is sufficient
	if inventoryResp.Inventory.Quantity < int32(req.Quantity) {
		return nil, errors.New("insufficient stock")
	}

	// Deduct inventory
	newQuantity := inventoryResp.Inventory.Quantity - int32(req.Quantity)
	_, err = s.inventoryClient.UpdateInventory(ctx, &inventory.UpdateInventoryRequest{
		Id:       inventoryResp.Inventory.Id,
		Quantity: newQuantity,
	})
	if err != nil {
		return nil, err
	}

	// Calculate total price
	totalPrice := productResp.Product.Price * float64(req.Quantity)

	order, err := s.repo.CreateOrder(ctx, &domain.Order{
		UserID:     userID,
		ProductID:  req.ProductID,
		Quantity:   req.Quantity,
		TotalPrice: totalPrice,
		Status:     domain.OrderStatusPending,
	})
	if err != nil {
		return nil, err
	}

	// create order notification
	msg := fmt.Sprintf("%d %s orders have been placed", order.Quantity, productResp.Product.Name)
	if err := s.mq.Publish(ctx, []byte(msg)); err != nil {
		return nil, err
	}

	return &domain.CreateOrderRes{
		ID:         order.ID,
		UserID:     order.UserID,
		ProductID:  order.ProductID,
		Quantity:   order.Quantity,
		TotalPrice: order.TotalPrice,
		Status:     order.Status,
	}, nil
}

func (s *OrderService) GetOrderByID(ctx context.Context, id uint) (*domain.GetOrderRes, error) {
	order, err := s.repo.GetOrderByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &domain.GetOrderRes{
		ID:         order.ID,
		UserID:     order.UserID,
		ProductID:  order.ProductID,
		Quantity:   order.Quantity,
		TotalPrice: order.TotalPrice,
		Status:     order.Status,
	}, nil
}

func (s *OrderService) GetOrdersByUserID(ctx context.Context, userID uint) ([]domain.GetOrderRes, error) {
	orders, err := s.repo.GetOrdersByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	var result []domain.GetOrderRes
	for _, order := range orders {
		result = append(result, domain.GetOrderRes{
			ID:         order.ID,
			UserID:     order.UserID,
			ProductID:  order.ProductID,
			Quantity:   order.Quantity,
			TotalPrice: order.TotalPrice,
			Status:     order.Status,
		})
	}

	return result, nil
}

func (s *OrderService) UpdateOrder(ctx context.Context, id uint, req *domain.UpdateOrderReq) (*domain.GetOrderRes, error) {
	order, err := s.repo.GetOrderByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if req.Status != nil {
		order.Status = *req.Status
	}

	updated, err := s.repo.UpdateOrder(ctx, order)
	if err != nil {
		return nil, err
	}

	return &domain.GetOrderRes{
		ID:         updated.ID,
		UserID:     updated.UserID,
		ProductID:  updated.ProductID,
		Quantity:   updated.Quantity,
		TotalPrice: updated.TotalPrice,
		Status:     updated.Status,
	}, nil
}

func (s *OrderService) DeleteOrder(ctx context.Context, id uint) error {
	return s.repo.DeleteOrder(ctx, id)
}
