package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/yehezkiel1086/go-jwt-auth/internal/adapter/config"
	"github.com/yehezkiel1086/go-jwt-auth/internal/adapter/handler"
	"github.com/yehezkiel1086/go-jwt-auth/internal/adapter/storage/postgres"
	"github.com/yehezkiel1086/go-jwt-auth/internal/adapter/storage/postgres/repository"
	"github.com/yehezkiel1086/go-jwt-auth/internal/adapter/storage/redis"
	"github.com/yehezkiel1086/go-jwt-auth/internal/core/domain"
	"github.com/yehezkiel1086/go-jwt-auth/internal/core/service"
)

func handleError(msg string, err error) {
	if err != nil {
		slog.Error(msg, "error", err)
		os.Exit(1)
	}
}

func main() {
	conf, err := config.New()
	handleError("failed to load .env configs", err)
	slog.Info(".env configs loaded successfully", "app", conf.App.Name, "env", conf.App.Env)

	ctx := context.Background()

	db, err := postgres.New(ctx, conf.DB)
	handleError("failed to init db", err)
	slog.Info("db initialized successfully")

	err = db.Migrate(&domain.User{})
	handleError("failed to migrate db", err)
	slog.Info("db migrated successfully")

	cache, err := redis.New(ctx, conf.Redis)
	handleError("failed to init redis", err)
	slog.Info("redis initialized successfully")

	userRepo := repository.NewUserRepository(db)
	userSvc := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userSvc)

	authSvc := service.NewAuthService(conf.JWT, cache, userRepo)
	authHandler := handler.NewAuthHandler(conf.App, conf.JWT, authSvc)

	r, err := handler.New(conf.JWT, conf.CORS, cache, userHandler, authHandler)
	handleError("failed to init router", err)
	slog.Info("router initialized successfully")

	err = r.Run(conf.HTTP)
	handleError("failed to run server", err)
}