package repository

import (
	"context"

	"github.com/yehezkiel1086/go-prometheus-grafana/internal/adapter/storage/postgres"
	"github.com/yehezkiel1086/go-prometheus-grafana/internal/core/domain"
)

type ProductRepository struct {
	db *postgres.DB
}

func NewProductRepository(db *postgres.DB) *ProductRepository {
	return &ProductRepository{db}
}

func (r *ProductRepository) CreateNewProduct(ctx context.Context, product *domain.Product) (*domain.Product, error) {
	db := r.db.GetDB()

	if err := db.WithContext(ctx).Create(product).Error; err != nil {
		return nil, err
	}

	return product, nil
}

func (r *ProductRepository) GetProductByID(ctx context.Context, id uint) (*domain.Product, error) {
	var product domain.Product

	db := r.db.GetDB()

	if err := db.WithContext(ctx).Preload("Stock").First(&product, id).Error; err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *ProductRepository) GetProductBySKU(ctx context.Context, sku string) (*domain.Product, error) {
	var product domain.Product

	db := r.db.GetDB()

	if err := db.WithContext(ctx).Preload("Stock").Where("sku = ?", sku).First(&product).Error; err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *ProductRepository) GetProductsByCategory(ctx context.Context, category string, limit, offset int) ([]domain.Product, error) {
	var products []domain.Product

	db := r.db.GetDB()

	query := db.WithContext(ctx).Where("category = ?", category).Limit(limit).Offset(offset)

	if err := query.Preload("Stock").Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

func (r *ProductRepository) ListProducts(ctx context.Context, limit, offset int) ([]domain.Product, error) {
	var products []domain.Product

	db := r.db.GetDB()

	if err := db.WithContext(ctx).Limit(limit).Offset(offset).Preload("Stock").Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

func (r *ProductRepository) UpdateProduct(ctx context.Context, product *domain.Product) (*domain.Product, error) {
	db := r.db.GetDB()

	if err := db.WithContext(ctx).Save(product).Error; err != nil {
		return nil, err
	}

	return product, nil
}

func (r *ProductRepository) DeleteProduct(ctx context.Context, id uint) error {
	db := r.db.GetDB()

	if err := db.WithContext(ctx).Delete(&domain.Product{}, id).Error; err != nil {
		return err
	}

	return nil
}
