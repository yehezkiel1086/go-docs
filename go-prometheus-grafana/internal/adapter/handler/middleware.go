package handler

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yehezkiel1086/go-prometheus-grafana/internal/core/port"
)

// PrometheusMiddleware records http_requests_total and http_request_duration_seconds
// for every incoming request. Attach it globally via r.Use() before any route groups.
func PrometheusMiddleware(mtx port.InventoryMetrics) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// Run the actual handler
		c.Next()

		duration := time.Since(start).Seconds()
		method := c.Request.Method
		// Use the matched route pattern (e.g. /api/v1/stocks/:id) instead of
		// the raw URL so high-cardinality paths like IDs don't explode the metric.
		path := c.FullPath()
		if path == "" {
			path = "unmatched"
		}
		status := c.Writer.Status()

		mtx.RecordHTTPRequest(c.Request.Context(), method, path, status, duration)
	}
}
