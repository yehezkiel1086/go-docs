package port

import (
	"context"

	"github.com/yehezkiel1086/go-grpc-inventory-microservices/services/product-service/internal/core/domain"
)

type ProductRepository interface {
	CreateProduct(ctx context.Context, product *domain.Product) (*domain.Product, error)
	GetProductByID(ctx context.Context, id uint) (*domain.Product, error)
	UpdateProduct(ctx context.Context, product *domain.Product) (*domain.Product, error)
	GetAllProducts(ctx context.Context) ([]domain.Product, error)
	DeleteProduct(ctx context.Context, id uint) error
}

type ProductService interface {
	CreateProduct(ctx context.Context, req *domain.CreateProductReq) (*domain.CreateProductRes, error)
	GetProductByID(ctx context.Context, id uint) (*domain.GetProductRes, error)
	UpdateProduct(ctx context.Context, id uint, req *domain.UpdateProductReq) (*domain.GetProductRes, error)
	GetAllProducts(ctx context.Context) ([]domain.GetProductRes, error)
	DeleteProduct(ctx context.Context, id uint) error
}
