package service

import (
	"context"
	"fmt"
	"time"

	"github.com/yehezkiel1086/go-prometheus-grafana/internal/core/domain"
	"github.com/yehezkiel1086/go-prometheus-grafana/internal/core/port"
)

type StockService struct {
	repo port.StockRepository
	mtx  port.InventoryMetrics
}

func NewStockService(repo port.StockRepository, mtx port.InventoryMetrics) *StockService {
	return &StockService{
		repo,
		mtx,
	}
}

func (s *StockService) CreateStock(ctx context.Context, stock *domain.Stock) (*domain.Stock, error) {
	start := time.Now()
	result, err := s.repo.CreateStock(ctx, stock)
	s.mtx.RecordTransactionDuration(ctx, time.Since(start).Seconds())
	if err == nil {
		s.mtx.IncStockUpdate(ctx, "create")
		// Update the total products gauge after a new SKU is created.
		// Requires CountStocks on the repository — see port.StockRepository.
		if count, err := s.repo.CountStocks(ctx); err == nil {
			s.mtx.UpdateTotalProducts(ctx, float64(count))
		}
	}
	return result, err
}

func (s *StockService) GetStockByID(ctx context.Context, id uint) (*domain.Stock, error) {
	start := time.Now()
	result, err := s.repo.GetStockByID(ctx, id)
	s.mtx.RecordTransactionDuration(ctx, time.Since(start).Seconds())
	return result, err
}

func (s *StockService) GetStockByProductID(ctx context.Context, productID uint) (*domain.Stock, error) {
	start := time.Now()
	result, err := s.repo.GetStockByProductID(ctx, productID)
	s.mtx.RecordTransactionDuration(ctx, time.Since(start).Seconds())
	return result, err
}

func (s *StockService) UpdateStock(ctx context.Context, stock *domain.Stock) (*domain.Stock, error) {
	start := time.Now()
	result, err := s.repo.UpdateStock(ctx, stock)
	s.mtx.RecordTransactionDuration(ctx, time.Since(start).Seconds())
	if err == nil {
		s.mtx.IncStockUpdate(ctx, "update")
		if result.Available() < 10 {
			// Stock has no Category field — use ProductID as the label so alerts
			// are at least scoped to a specific product rather than all lumped
			// under "general". If a Category field is added to domain.Stock or
			// domain.Product in future, swap this out.
			s.mtx.TrackLowStockAlert(ctx, fmt.Sprintf("product_%d", result.ProductID))
		}
	}
	return result, err
}

func (s *StockService) UpdateStockQuantity(ctx context.Context, id uint, quantity int) (*domain.Stock, error) {
	start := time.Now()
	result, err := s.repo.UpdateStockQuantity(ctx, id, quantity)
	s.mtx.RecordTransactionDuration(ctx, time.Since(start).Seconds())
	if err == nil {
		s.mtx.IncStockUpdate(ctx, "restock")
		if result.Available() < 10 {
			s.mtx.TrackLowStockAlert(ctx, fmt.Sprintf("product_%d", result.ProductID))
		}
	}
	return result, err
}

func (s *StockService) ReserveStock(ctx context.Context, productID uint, quantity int) error {
	start := time.Now()
	err := s.repo.ReserveStock(ctx, productID, quantity)
	s.mtx.RecordTransactionDuration(ctx, time.Since(start).Seconds())
	if err == nil {
		s.mtx.IncStockUpdate(ctx, "reserve")
	}
	return err
}

func (s *StockService) ReleaseStock(ctx context.Context, productID uint, quantity int) error {
	start := time.Now()
	err := s.repo.ReleaseStock(ctx, productID, quantity)
	s.mtx.RecordTransactionDuration(ctx, time.Since(start).Seconds())
	if err == nil {
		s.mtx.IncStockUpdate(ctx, "release")
	}
	return err
}

func (s *StockService) DeleteStock(ctx context.Context, id uint) error {
	start := time.Now()
	err := s.repo.DeleteStock(ctx, id)
	s.mtx.RecordTransactionDuration(ctx, time.Since(start).Seconds())
	if err == nil {
		s.mtx.IncStockUpdate(ctx, "delete")
		// Update the total products gauge after a SKU is removed.
		if count, err := s.repo.CountStocks(ctx); err == nil {
			s.mtx.UpdateTotalProducts(ctx, float64(count))
		}
	}
	return err
}
