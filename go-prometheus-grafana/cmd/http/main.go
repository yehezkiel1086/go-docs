package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/yehezkiel1086/go-prometheus-grafana/internal/adapter/config"
	"github.com/yehezkiel1086/go-prometheus-grafana/internal/adapter/handler"
	"github.com/yehezkiel1086/go-prometheus-grafana/internal/adapter/monitoring"
	"github.com/yehezkiel1086/go-prometheus-grafana/internal/adapter/storage/postgres"
	"github.com/yehezkiel1086/go-prometheus-grafana/internal/adapter/storage/postgres/repository"
	"github.com/yehezkiel1086/go-prometheus-grafana/internal/core/domain"
	"github.com/yehezkiel1086/go-prometheus-grafana/internal/core/service"
)

func handleError(msg string, err error) {
	if err != nil {
		slog.Error(msg, "error", err)
		os.Exit(1)
	}
}

func main() {
	// load .env configs
	conf, err := config.New()
	handleError("failed to load .env configs", err)
	slog.Info(".env config loaded successfully", "name", conf.App.Name, "env", conf.App.Env)

	ctx := context.Background()

	// init postgres db
	db, err := postgres.New(ctx, conf.DB)
	handleError("failed to connect to postgres db", err)
	slog.Info("postgres db connected successfully", "db", conf.DB.Name)

	// migrate dbs
	err = db.Migrate(&domain.User{}, &domain.Product{}, &domain.Stock{}, &domain.Order{}, &domain.OrderItem{})
	handleError("failed to migrate dbs", err)
	slog.Info("dbs migrated successfully")

	// init prometheus metrics
	mtx := monitoring.NewPrometheusMetrics()

	// dependency injections
	userRepo := repository.NewUserRepository(db)
	userSvc := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userSvc)

	productRepo := repository.NewProductRepository(db)
	productSvc := service.NewProductService(productRepo)
	productHandler := handler.NewProductHandler(productSvc)

	stockRepo := repository.NewStockRepository(db)
	stockSvc := service.NewStockService(stockRepo, mtx)
	stockHandler := handler.NewStockHandler(stockSvc)

	orderRepo := repository.NewOrderRepository(db)
	orderSvc := service.NewOrderService(orderRepo)
	orderHandler := handler.NewOrderHandler(orderSvc)

	// init router
	r := handler.NewRouter(
		userHandler,
		productHandler,
		stockHandler,
		orderHandler,
	)

	// run api server
	err = r.Run(conf.HTTP)
	handleError("failed to run api server", err)
}
