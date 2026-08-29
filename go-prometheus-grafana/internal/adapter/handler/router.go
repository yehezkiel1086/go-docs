package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/yehezkiel1086/go-prometheus-grafana/internal/adapter/config"
	"github.com/yehezkiel1086/go-prometheus-grafana/internal/core/port"
)

type Router struct {
	r *gin.Engine
}

func NewRouter(
	userHandler *UserHandler,
	productHandler *ProductHandler,
	stockHandler *StockHandler,
	orderHandler *OrderHandler,
	mtx port.InventoryMetrics,
) *Router {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(PrometheusMiddleware(mtx)) // records http_requests_total + http_request_duration_seconds

	// Prometheus metrics endpoint — must come AFTER middleware registration
	// but the middleware itself skips nothing; /metrics will also be measured,
	// which is fine (it's a real HTTP call).
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// rbac
	pb := r.Group("/api/v1")

	// public user routes
	pb.POST("/register", userHandler.Register)
	pb.GET("/users", userHandler.ListUsers)

	// product routes
	products := pb.Group("/products")
	products.POST("", productHandler.CreateProduct)
	products.GET("", productHandler.ListProducts)
	products.GET("/sku/:sku", productHandler.GetProductBySKU)
	products.GET("/category/:category", productHandler.GetProductsByCategory)
	products.GET("/:id", productHandler.GetProduct)
	products.PUT("/:id", productHandler.UpdateProduct)
	products.DELETE("/:id", productHandler.DeleteProduct)

	// stock routes
	stocks := pb.Group("/stocks")
	stocks.POST("", stockHandler.CreateStock)
	stocks.GET("/:id", stockHandler.GetStock)
	stocks.GET("/product/:product_id", stockHandler.GetStockByProductID)
	stocks.PUT("/:id", stockHandler.UpdateStock)
	stocks.PATCH("/:id/quantity", stockHandler.UpdateStockQuantity)
	stocks.POST("/product/:product_id/reserve", stockHandler.ReserveStock)
	stocks.POST("/product/:product_id/release", stockHandler.ReleaseStock)
	stocks.DELETE("/:id", stockHandler.DeleteStock)

	// order routes
	orders := pb.Group("/orders")
	orders.POST("", orderHandler.CreateOrder)
	orders.GET("", orderHandler.ListOrders)
	orders.GET("/user/:user_id", orderHandler.GetOrdersByUser)
	orders.GET("/:id", orderHandler.GetOrder)
	orders.PUT("/:id/status", orderHandler.UpdateOrderStatus)
	orders.PUT("/:id", orderHandler.UpdateOrder)
	orders.DELETE("/:id", orderHandler.DeleteOrder)

	return &Router{r: r}
}

func (r *Router) Run(conf *config.HTTP) error {
	return r.r.Run(conf.Host + ":" + conf.Port)
}
