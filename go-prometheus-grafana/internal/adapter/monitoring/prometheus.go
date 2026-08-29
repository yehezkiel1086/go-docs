package monitoring

import (
	"context"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// PrometheusMetrics implements port.InventoryMetrics
type PrometheusMetrics struct {
	stockUpdates        *prometheus.CounterVec
	totalProducts       prometheus.Gauge
	lowStockAlerts      *prometheus.CounterVec
	transactionDuration prometheus.Histogram
	httpRequests        *prometheus.CounterVec
	httpDuration        *prometheus.HistogramVec
}

func NewPrometheusMetrics() *PrometheusMetrics {
	return &PrometheusMetrics{
		stockUpdates: promauto.NewCounterVec(
			prometheus.CounterOpts{
				Name: "inventory_stock_updates_total",
				Help: "Total number of stock updates by action type",
			},
			[]string{"action"}, // restock, sale, return
		),
		totalProducts: promauto.NewGauge(
			prometheus.GaugeOpts{
				Name: "inventory_products_total",
				Help: "Current number of unique products (SKUs)",
			},
		),
		lowStockAlerts: promauto.NewCounterVec(
			prometheus.CounterOpts{
				Name: "inventory_low_stock_alerts_total",
				Help: "Total number of low stock alerts by category",
			},
			[]string{"category"},
		),
		transactionDuration: promauto.NewHistogram(
			prometheus.HistogramOpts{
				Name:    "inventory_transaction_duration_seconds",
				Help:    "Transaction duration distribution",
				Buckets: prometheus.DefBuckets,
			},
		),
		httpRequests: promauto.NewCounterVec(
			prometheus.CounterOpts{
				Name: "http_requests_total",
				Help: "Total HTTP requests by method, path, and status",
			},
			[]string{"method", "path", "status"},
		),
		httpDuration: promauto.NewHistogramVec(
			prometheus.HistogramOpts{
				Name:    "http_request_duration_seconds",
				Help:    "HTTP request duration by method and path",
				Buckets: prometheus.DefBuckets,
			},
			[]string{"method", "path"},
		),
	}
}

func (m *PrometheusMetrics) IncStockUpdate(ctx context.Context, action string) {
	m.stockUpdates.WithLabelValues(action).Inc()
}

func (m *PrometheusMetrics) UpdateTotalProducts(ctx context.Context, count float64) {
	m.totalProducts.Set(count)
}

func (m *PrometheusMetrics) TrackLowStockAlert(ctx context.Context, category string) {
	m.lowStockAlerts.WithLabelValues(category).Inc()
}

func (m *PrometheusMetrics) RecordTransactionDuration(ctx context.Context, seconds float64) {
	m.transactionDuration.Observe(seconds)
}

func (m *PrometheusMetrics) RecordHTTPRequest(ctx context.Context, method, path string, status int, duration float64) {
	m.httpRequests.WithLabelValues(method, path, strconv.Itoa(status)).Inc()
	m.httpDuration.WithLabelValues(method, path).Observe(duration)
}

func (m *PrometheusMetrics) IncHTTPRequest(ctx context.Context, method, path string, status int) {
	m.httpRequests.WithLabelValues(method, path, strconv.Itoa(status)).Inc()
}

// Timer helper for measuring operation duration
func (m *PrometheusMetrics) StartTimer() *Timer {
	return &Timer{start: time.Now(), metrics: m}
}

type Timer struct {
	start   time.Time
	metrics *PrometheusMetrics
}

func (t *Timer) ObserveTransaction(ctx context.Context) {
	duration := time.Since(t.start).Seconds()
	t.metrics.RecordTransactionDuration(ctx, duration)
}

func (t *Timer) ObserveHTTP(ctx context.Context, method, path string, status int) {
	duration := time.Since(t.start).Seconds()
	t.metrics.RecordHTTPRequest(ctx, method, path, status, duration)
}
