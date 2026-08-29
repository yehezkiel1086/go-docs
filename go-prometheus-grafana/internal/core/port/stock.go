package port

import (
	"context"

	"github.com/yehezkiel1086/go-prometheus-grafana/internal/core/domain"
)

type StockRepository interface {
	CreateStock(ctx context.Context, stock *domain.Stock) (*domain.Stock, error)
	GetStockByID(ctx context.Context, id uint) (*domain.Stock, error)
	GetStockByProductID(ctx context.Context, productID uint) (*domain.Stock, error)
	UpdateStock(ctx context.Context, stock *domain.Stock) (*domain.Stock, error)
	UpdateStockQuantity(ctx context.Context, id uint, quantity int) (*domain.Stock, error)
	ReserveStock(ctx context.Context, productID uint, quantity int) error
	ReleaseStock(ctx context.Context, productID uint, quantity int) error
	DeleteStock(ctx context.Context, id uint) error

	// CountStocks returns the total number of unique stock records (SKUs).
	// Called by the service layer to keep the inventory_products_total gauge accurate
	// after create and delete operations.
	CountStocks(ctx context.Context) (int64, error)
}

type StockService interface {
	CreateStock(ctx context.Context, stock *domain.Stock) (*domain.Stock, error)
	GetStockByID(ctx context.Context, id uint) (*domain.Stock, error)
	GetStockByProductID(ctx context.Context, productID uint) (*domain.Stock, error)
	UpdateStock(ctx context.Context, stock *domain.Stock) (*domain.Stock, error)
	UpdateStockQuantity(ctx context.Context, id uint, quantity int) (*domain.Stock, error)
	ReserveStock(ctx context.Context, productID uint, quantity int) error
	ReleaseStock(ctx context.Context, productID uint, quantity int) error
	DeleteStock(ctx context.Context, id uint) error
}
