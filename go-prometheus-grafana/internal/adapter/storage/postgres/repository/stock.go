package repository

import (
	"context"
	"errors"

	"github.com/yehezkiel1086/go-prometheus-grafana/internal/adapter/storage/postgres"
	"github.com/yehezkiel1086/go-prometheus-grafana/internal/core/domain"
	"gorm.io/gorm"
)

var (
	ErrInsufficientStock = errors.New("insufficient stock available")
	ErrInvalidRelease    = errors.New("cannot release more than reserved quantity")
)

type StockRepository struct {
	db *postgres.DB
}

func NewStockRepository(db *postgres.DB) *StockRepository {
	return &StockRepository{db}
}

func (r *StockRepository) CreateStock(ctx context.Context, stock *domain.Stock) (*domain.Stock, error) {
	db := r.db.GetDB()

	if err := db.WithContext(ctx).Create(stock).Error; err != nil {
		return nil, err
	}

	return stock, nil
}

func (r *StockRepository) GetStockByID(ctx context.Context, id uint) (*domain.Stock, error) {
	var stock domain.Stock

	db := r.db.GetDB()

	if err := db.WithContext(ctx).Preload("Product").First(&stock, id).Error; err != nil {
		return nil, err
	}

	return &stock, nil
}

func (r *StockRepository) GetStockByProductID(ctx context.Context, productID uint) (*domain.Stock, error) {
	var stock domain.Stock

	db := r.db.GetDB()

	if err := db.WithContext(ctx).Preload("Product").Where("product_id = ?", productID).First(&stock).Error; err != nil {
		return nil, err
	}

	return &stock, nil
}

func (r *StockRepository) UpdateStock(ctx context.Context, stock *domain.Stock) (*domain.Stock, error) {
	db := r.db.GetDB()

	// Only update mutable fields (not product_id which is immutable)
	if err := db.WithContext(ctx).Model(&domain.Stock{}).
		Where("id = ?", stock.ID).
		Updates(map[string]interface{}{
			"quantity": stock.Quantity,
			"reserved": stock.Reserved,
		}).Error; err != nil {
		return nil, err
	}

	return r.GetStockByID(ctx, stock.ID)
}

func (r *StockRepository) UpdateStockQuantity(ctx context.Context, id uint, quantity int) (*domain.Stock, error) {
	db := r.db.GetDB()

	if err := db.WithContext(ctx).Model(&domain.Stock{}).Where("id = ?", id).Update("quantity", quantity).Error; err != nil {
		return nil, err
	}

	return r.GetStockByID(ctx, id)
}

func (r *StockRepository) ReserveStock(ctx context.Context, productID uint, quantity int) error {
	db := r.db.GetDB()

	result := db.WithContext(ctx).Model(&domain.Stock{}).
		Where("product_id = ? AND (quantity - reserved) >= ?", productID, quantity).
		UpdateColumn("reserved", gorm.Expr("reserved + ?", quantity))

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return ErrInsufficientStock
	}

	return nil
}

func (r *StockRepository) ReleaseStock(ctx context.Context, productID uint, quantity int) error {
	db := r.db.GetDB()

	result := db.WithContext(ctx).Model(&domain.Stock{}).
		Where("product_id = ? AND reserved >= ?", productID, quantity).
		UpdateColumn("reserved", gorm.Expr("reserved - ?", quantity))

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return ErrInvalidRelease
	}

	return nil
}

func (r *StockRepository) DeleteStock(ctx context.Context, id uint) error {
	db := r.db.GetDB()

	if err := db.WithContext(ctx).Delete(&domain.Stock{}, id).Error; err != nil {
		return err
	}

	return nil
}

// CountStocks returns the total number of stock records in the database.
// Used by the service layer to keep the inventory_products_total Prometheus
// gauge accurate after create and delete operations.
func (r *StockRepository) CountStocks(ctx context.Context) (int64, error) {
	var count int64
	db := r.db.GetDB()

	if err := db.WithContext(ctx).Model(&domain.Stock{}).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}
