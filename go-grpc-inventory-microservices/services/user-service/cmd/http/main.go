package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/yehezkiel1086/go-grpc-inventory-microservices/services/user-service/internal/adapter/config"
	"github.com/yehezkiel1086/go-grpc-inventory-microservices/services/user-service/internal/adapter/handler"
	"github.com/yehezkiel1086/go-grpc-inventory-microservices/services/user-service/internal/adapter/storage/postgres"
	"github.com/yehezkiel1086/go-grpc-inventory-microservices/services/user-service/internal/adapter/storage/postgres/repository"
	"github.com/yehezkiel1086/go-grpc-inventory-microservices/services/user-service/internal/core/domain"
	"github.com/yehezkiel1086/go-grpc-inventory-microservices/services/user-service/internal/core/service"
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
	err = db.Migrate(&domain.User{})
	handleError("failed to migrate dbs", err)
	slog.Info("dbs migrated successfully")

	// dependency injections
	userRepo := repository.NewUserRepository(db)
	userSvc := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userSvc)

	authSvc := service.NewAuthService(conf.JWT, userRepo)
	authHandler := handler.NewAuthHandler(conf.JWT, authSvc)

	// init router
	router, err := handler.NewRouter(
		conf.JWT,
		userHandler,
		authHandler,
	)
	handleError("failed to init router", err)

	// start server
	err = router.Run(conf.HTTP)
	handleError("failed to start server", err)
}
