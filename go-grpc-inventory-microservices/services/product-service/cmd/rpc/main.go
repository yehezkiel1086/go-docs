package main

import (
	"context"
	"log/slog"
	"os"

	inventory "github.com/yehezkiel1086/go-grpc-inventory-microservices/services/common/protobuf/inventory/protobuf"
	"github.com/yehezkiel1086/go-grpc-inventory-microservices/services/product-service/internal/adapter/config"
	"github.com/yehezkiel1086/go-grpc-inventory-microservices/services/product-service/internal/adapter/handler"
	"github.com/yehezkiel1086/go-grpc-inventory-microservices/services/product-service/internal/adapter/storage/postgres"
	"github.com/yehezkiel1086/go-grpc-inventory-microservices/services/product-service/internal/adapter/storage/postgres/repository"
	"github.com/yehezkiel1086/go-grpc-inventory-microservices/services/product-service/internal/core/domain"
	"github.com/yehezkiel1086/go-grpc-inventory-microservices/services/product-service/internal/core/service"
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
	err = db.Migrate(&domain.Product{})
	handleError("failed to migrate dbs", err)
	slog.Info("dbs migrated successfully")

	// init grpc server
	grpcServer, err := handler.NewGRPCServer(conf.RPC)
	handleError("failed to init grpc server", err)
	slog.Info("product grpc server initialized successfully", "addr", conf.RPC.Host+":"+conf.RPC.Port)

	// connect to inventory gRPC service
	grpcClient, err := handler.NewGRPCClient(conf.Inventory)
	handleError("failed to create grpc client", err)
	defer grpcClient.Close()

	slog.Info("connected to inventory service", "addr", conf.Inventory.Host+":"+conf.Inventory.Port)

	// dependency injections
	productRepo := repository.NewProductRepository(db)
	inventoryClient := inventory.NewInventoryServiceClient(grpcClient.GetConn())
	productSvc := service.NewProductService(productRepo, inventoryClient)
	handler.NewProductGRPCHandler(grpcServer.GetServer(), productSvc)

	// start server
	err = grpcServer.Run(conf.RPC)
	handleError("failed to start server", err)
	slog.Info("product grpc server running successfully")
}
