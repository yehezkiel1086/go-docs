package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/yehezkiel1086/go-socket-chat/internal/adapter/config"
	"github.com/yehezkiel1086/go-socket-chat/internal/adapter/handler"
	"github.com/yehezkiel1086/go-socket-chat/internal/adapter/storage/postgres"
	"github.com/yehezkiel1086/go-socket-chat/internal/adapter/storage/postgres/repository"
	"github.com/yehezkiel1086/go-socket-chat/internal/core/domain"
	"github.com/yehezkiel1086/go-socket-chat/internal/core/service"
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

	// init postgres db
	db, err := postgres.New(ctx, conf.DB)
	handleError("db connection failed", err)
	slog.Info("db connected successfully", "driver", conf.DB.Connection, "db", conf.DB.Name)

	// migrate dbs
	err = db.Migrate(&domain.User{})
	handleError("failed to migrate dbs", err)
	slog.Info("dbs migrated successfully")

	// dependency injection
	userRepo := repository.NewUserRepository(db)
	userSvc := service.NewUserService(conf.JWT, userRepo)
	userHandler := handler.NewUserHandler(userSvc)

	// init router
	r := handler.NewRouter(userHandler)

	// run backend api
	err = r.Run(conf.HTTP)
	handleError("failed to run go backend api", err)
}
