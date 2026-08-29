package service

import (
	"context"

	"github.com/yehezkiel1086/go-prometheus-grafana/internal/core/domain"
	"github.com/yehezkiel1086/go-prometheus-grafana/internal/core/port"
)

type ProductService struct {
	repo port.ProductRepository
}

func NewProductService(repo port.ProductRepository) *ProductService {
	return &ProductService{
		repo,
	}
}

func (s *ProductService) CreateNewProduct(ctx context.Context, product *domain.Product) (*domain.Product, error) {
	return s.repo.CreateNewProduct(ctx, product)
}

func (s *ProductService) GetProductByID(ctx context.Context, id uint) (*domain.Product, error) {
	return s.repo.GetProductByID(ctx, id)
}

func (s *ProductService) GetProductBySKU(ctx context.Context, sku string) (*domain.Product, error) {
	return s.repo.GetProductBySKU(ctx, sku)
}

func (s *ProductService) GetProductsByCategory(ctx context.Context, category string, limit, offset int) ([]domain.Product, error) {
	return s.repo.GetProductsByCategory(ctx, category, limit, offset)
}

func (s *ProductService) ListProducts(ctx context.Context, limit, offset int) ([]domain.Product, error) {
	return s.repo.ListProducts(ctx, limit, offset)
}

func (s *ProductService) UpdateProduct(ctx context.Context, product *domain.Product) (*domain.Product, error) {
	return s.repo.UpdateProduct(ctx, product)
}

func (s *ProductService) DeleteProduct(ctx context.Context, id uint) error {
	return s.repo.DeleteProduct(ctx, id)
}
