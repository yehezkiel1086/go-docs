package handler

import (
	"context"

	product "github.com/yehezkiel1086/go-grpc-inventory-microservices/services/common/protobuf/product/protobuf"
	"github.com/yehezkiel1086/go-grpc-inventory-microservices/services/product-service/internal/core/domain"
	"github.com/yehezkiel1086/go-grpc-inventory-microservices/services/product-service/internal/core/port"
	"google.golang.org/grpc"
)

type ProductGRPCHandler struct {
	svc port.ProductService
	product.UnimplementedProductServiceServer
}

func NewProductGRPCHandler(server *grpc.Server, svc port.ProductService) {
	productGRPCHandler := &ProductGRPCHandler{
		svc: svc,
	}
	product.RegisterProductServiceServer(server, productGRPCHandler)
}

func (h *ProductGRPCHandler) CreateProduct(ctx context.Context, req *product.CreateProductRequest) (*product.CreateProductResponse, error) {
	createdProduct, err := h.svc.CreateProduct(ctx, &domain.CreateProductReq{
		Name:        req.Name,
		Price:       req.Price,
		Description: req.Description,
		Quantity:    int(req.Quantity),
	})
	if err != nil {
		return nil, err
	}

	return &product.CreateProductResponse{
		Product: &product.Product{
			Id:          uint64(createdProduct.ID),
			Name:        createdProduct.Name,
			Price:       createdProduct.Price,
			Description: createdProduct.Description,
			Quantity:    int32(createdProduct.Quantity),
		},
	}, nil
}

func (h *ProductGRPCHandler) GetProduct(ctx context.Context, req *product.GetProductRequest) (*product.GetProductResponse, error) {
	productRes, err := h.svc.GetProductByID(ctx, uint(req.Id))
	if err != nil {
		return nil, err
	}

	return &product.GetProductResponse{
		Product: &product.Product{
			Id:          uint64(productRes.ID),
			Name:        productRes.Name,
			Price:       productRes.Price,
			Description: productRes.Description,
			Quantity:    int32(productRes.Quantity),
		},
	}, nil
}

func (h *ProductGRPCHandler) UpdateProduct(ctx context.Context, req *product.UpdateProductRequest) (*product.UpdateProductResponse, error) {
	var name *string
	var price *float64
	var description *string
	var quantity *int

	if req.Name != "" {
		name = &req.Name
	}
	if req.Price != 0 {
		price = &req.Price
	}
	if req.Description != "" {
		description = &req.Description
	}
	if req.Quantity != 0 {
		q := int(req.Quantity)
		quantity = &q
	}

	updatedProduct, err := h.svc.UpdateProduct(ctx, uint(req.Id), &domain.UpdateProductReq{
		Name:        name,
		Price:       price,
		Description: description,
		Quantity:    quantity,
	})
	if err != nil {
		return nil, err
	}

	return &product.UpdateProductResponse{
		Product: &product.Product{
			Id:          uint64(updatedProduct.ID),
			Name:        updatedProduct.Name,
			Price:       updatedProduct.Price,
			Description: updatedProduct.Description,
			Quantity:    int32(updatedProduct.Quantity),
		},
	}, nil
}

func (h *ProductGRPCHandler) DeleteProduct(ctx context.Context, req *product.DeleteProductRequest) (*product.DeleteProductResponse, error) {
	err := h.svc.DeleteProduct(ctx, uint(req.Id))
	if err != nil {
		return nil, err
	}

	return &product.DeleteProductResponse{
		Success: true,
	}, nil
}
