package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/yehezkiel1086/go-grpc-inventory-microservices/services/inventory-service/internal/adapter/config"
	"github.com/yehezkiel1086/go-grpc-inventory-microservices/services/inventory-service/internal/adapter/handler"
	"github.com/yehezkiel1086/go-grpc-inventory-microservices/services/inventory-service/internal/adapter/storage/postgres"
	"github.com/yehezkiel1086/go-grpc-inventory-microservices/services/inventory-service/internal/adapter/storage/postgres/repository"
	"github.com/yehezkiel1086/go-grpc-inventory-microservices/services/inventory-service/internal/core/domain"
	"github.com/yehezkiel1086/go-grpc-inventory-microservices/services/inventory-service/internal/core/service"
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
	err = db.Migrate(&domain.Inventory{})
	handleError("failed to migrate dbs", err)
	slog.Info("dbs migrated successfully")

	// init grpc server
	grpcServer, err := handler.NewGRPCServer(conf.RPC)
	handleError("failed to init grpc server", err)
	slog.Info("inventory grpc server initialized successfully", "addr", conf.RPC.Host+":"+conf.RPC.Port)

	// dependency injections
	inventoryRepo := repository.NewInventoryRepository(db)
	inventorySvc := service.NewInventoryService(inventoryRepo)
	handler.NewInventoryHandler(grpcServer.GetServer(), inventorySvc)

	// run grpc server
	err = grpcServer.Run(conf.RPC)
	handleError("failed to run grpc server", err)
	slog.Info("inventory grpc server running successfully")
}
