package main

import (
	"context"
	"log/slog"
	"os"

	go_grpc_inventory_microservices "github.com/yehezkiel1086/go-grpc-inventory-microservices/services/common/protobuf/inventory/protobuf"
	product "github.com/yehezkiel1086/go-grpc-inventory-microservices/services/common/protobuf/product/protobuf"
	"github.com/yehezkiel1086/go-grpc-inventory-microservices/services/order-service/internal/adapter/config"
	"github.com/yehezkiel1086/go-grpc-inventory-microservices/services/order-service/internal/adapter/handler"
	"github.com/yehezkiel1086/go-grpc-inventory-microservices/services/order-service/internal/adapter/storage/postgres"
	"github.com/yehezkiel1086/go-grpc-inventory-microservices/services/order-service/internal/adapter/storage/postgres/repository"
	"github.com/yehezkiel1086/go-grpc-inventory-microservices/services/order-service/internal/adapter/storage/rabbitmq"
	"github.com/yehezkiel1086/go-grpc-inventory-microservices/services/order-service/internal/core/domain"
	"github.com/yehezkiel1086/go-grpc-inventory-microservices/services/order-service/internal/core/service"
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
	slog.Info(".env configs loaded successfully", "app", conf.App.Name, "env", conf.App.Env)

	ctx := context.Background()

	// init db
	db, err := postgres.New(ctx, conf.DB)
	handleError("failed to init db", err)
	slog.Info("db initialized successfully", "db", conf.DB.Name)

	// migrate dbs
	err = db.Migrate(&domain.Order{})
	handleError("failed to migrate dbs", err)
	slog.Info("dbs migrated successfully")

	// init rabbitmq
	rabbitmq, err := rabbitmq.New(conf.Rabbitmq)
	handleError("failed to init rabbitmq", err)

	defer rabbitmq.Close()

	slog.Info("rabbitmq initialized successfully")

	// connect to product gRPC service
	grpcProductClient, err := handler.NewGRPCClient(conf.ProductRPC.Host + ":" + conf.ProductRPC.Port)
	handleError("failed to create grpc client", err)
	defer grpcProductClient.Close()

	slog.Info("connected to product service", "addr", conf.ProductRPC.Host+":"+conf.ProductRPC.Port)

	// connect to inventory gRPC service
	grpcInventoryClient, err := handler.NewGRPCClient(conf.InventoryRPC.Host + ":" + conf.InventoryRPC.Port)
	handleError("failed to create grpc client", err)
	defer grpcInventoryClient.Close()

	slog.Info("connected to inventory service", "addr", conf.InventoryRPC.Host+":"+conf.InventoryRPC.Port)

	// init rpc clients
	inventoryClient := go_grpc_inventory_microservices.NewInventoryServiceClient(grpcInventoryClient.GetConn())
	productClient := product.NewProductServiceClient(grpcProductClient.GetConn())

	// dependency injections
	orderRepo := repository.NewOrderRepository(db)
	orderSvc := service.NewOrderService(orderRepo, productClient, inventoryClient, rabbitmq)
	orderHandler := handler.NewOrderHandler(orderSvc)

	// init router
	router, err := handler.NewRouter(
		conf.JWT,
		orderHandler,
	)
	handleError("failed to init router", err)

	// start server
	err = router.Run(conf.HTTP)
	handleError("failed to start server", err)
}
