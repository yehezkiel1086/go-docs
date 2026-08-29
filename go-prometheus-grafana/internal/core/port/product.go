package port

import (
	"context"

	"github.com/yehezkiel1086/go-prometheus-grafana/internal/core/domain"
)

type ProductRepository interface {
	CreateNewProduct(ctx context.Context, product *domain.Product) (*domain.Product, error)
	GetProductByID(ctx context.Context, id uint) (*domain.Product, error)
	GetProductBySKU(ctx context.Context, sku string) (*domain.Product, error)
	GetProductsByCategory(ctx context.Context, category string, limit, offset int) ([]domain.Product, error)
	ListProducts(ctx context.Context, limit, offset int) ([]domain.Product, error)
	UpdateProduct(ctx context.Context, product *domain.Product) (*domain.Product, error)
	DeleteProduct(ctx context.Context, id uint) error
}

type ProductService interface {
	CreateNewProduct(ctx context.Context, product *domain.Product) (*domain.Product, error)
	GetProductByID(ctx context.Context, id uint) (*domain.Product, error)
	GetProductBySKU(ctx context.Context, sku string) (*domain.Product, error)
	GetProductsByCategory(ctx context.Context, category string, limit, offset int) ([]domain.Product, error)
	ListProducts(ctx context.Context, limit, offset int) ([]domain.Product, error)
	UpdateProduct(ctx context.Context, product *domain.Product) (*domain.Product, error)
	DeleteProduct(ctx context.Context, id uint) error
}
