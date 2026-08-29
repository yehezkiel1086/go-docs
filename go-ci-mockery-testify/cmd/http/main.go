package main

import (
	"log/slog"
	"os"

	"github.com/yehezkiel1086/go-rest-mockery-testify/internal/adapter/config"
	"github.com/yehezkiel1086/go-rest-mockery-testify/internal/adapter/handler"
	"github.com/yehezkiel1086/go-rest-mockery-testify/internal/adapter/storage/postgres"
	"github.com/yehezkiel1086/go-rest-mockery-testify/internal/adapter/storage/postgres/repository"
	"github.com/yehezkiel1086/go-rest-mockery-testify/internal/core/domain"
	"github.com/yehezkiel1086/go-rest-mockery-testify/internal/core/service"
)

func handleError(err error, msg string) {
	if err != nil {
		slog.Error(msg, "error", err)
		os.Exit(1)
	}
}

func main() {
	// init .env configs
	conf, err := config.New()
	handleError(err, "failed to load .env configs")
	slog.Info(".env configs loaded successfully", "app", conf.App.Name)

	// init db
	db, err := postgres.New(conf.DB)
	handleError(err, "failed to load db")
	slog.Info("db loaded successfully", "db", conf.DB.Name)

	// migrate dbs
	err = db.Migrate(&domain.Task{})
	handleError(err, "failed to migrate db")
	slog.Info("dbs migrated successfully")

	// dependency injections
	taskRepo := repository.NewTaskRepository(db)
	taskService := service.NewTaskService(taskRepo)
	taskHandler := handler.NewTaskHandler(taskService)

	// init router
	r := handler.New(taskHandler)

	// run api server
	err = r.Run(conf.HTTP)
	handleError(err, "failed to run api server")
}
