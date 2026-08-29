package port

import "context"

type InventoryMetrics interface {
	// Business Metrics
	IncStockUpdate(ctx context.Context, action string)      // e.g., "restock", "sale", "return"
	UpdateTotalProducts(ctx context.Context, count float64) // Current unique SKUs
	TrackLowStockAlert(ctx context.Context, category string)

	// Performance
	RecordTransactionDuration(ctx context.Context, seconds float64)

	// HTTP Metrics
	RecordHTTPRequest(ctx context.Context, method, path string, status int, duration float64)
	IncHTTPRequest(ctx context.Context, method, path string, status int)
}
