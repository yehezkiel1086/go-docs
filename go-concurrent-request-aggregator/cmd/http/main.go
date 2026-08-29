package main

import (
	"log/slog"
	"os"

	"github.com/yehezkiel1086/go-concurrent-request-aggregator/internal/adapter/api"
	"github.com/yehezkiel1086/go-concurrent-request-aggregator/internal/adapter/api/repository"
	"github.com/yehezkiel1086/go-concurrent-request-aggregator/internal/adapter/config"
	"github.com/yehezkiel1086/go-concurrent-request-aggregator/internal/adapter/handler"
	"github.com/yehezkiel1086/go-concurrent-request-aggregator/internal/core/service"
)

func handleError(err error, msg string) {
	if err != nil {
		slog.Error(msg, "error", err)
		os.Exit(1)
	}
}

func main() {
	// load .env configs
	conf, err := config.New()
	handleError(err, "failed to load .env configs")
	slog.Info(".env configs loaded successfully", "name", conf.App.Name, "env", conf.App.Env)

	// init api
	api := api.New(conf.API)

	// dependency injection
	userRepo := repository.NewUserRepository(api)
	postRepo := repository.NewPostRepository(api)

	dashboardSvc := service.NewDashboardService(userRepo, postRepo)
	dashboardHandler := handler.NewDashboardHandler(dashboardSvc)

	// init router
	r := handler.NewRouter(dashboardHandler)

	// serve api
	err = r.Run(conf.HTTP)
	handleError(err, "failed to serve api")
}
