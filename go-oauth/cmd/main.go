package main

import (
	"context"
	"fmt"
	"go-oauth/internal/adapter/config"
	"go-oauth/internal/adapter/handler"
	"go-oauth/internal/adapter/storage/db/postgres"
	"go-oauth/internal/adapter/storage/db/postgres/repository"
	"go-oauth/internal/core/domain"
	"go-oauth/internal/core/service"
)

func main() {
	// get .env configs
	conf, err := config.InitConfig()
	if err != nil {
		panic(err)
	}
	fmt.Println("✅ successfully imported .env configs")

	ctx := context.Background()

	// init db
	db, err := postgres.InitDB(ctx, conf.DB)
	if err != nil {
		panic(err)
	}
	fmt.Println("✅ successfully connected to DB")

	// migrate database
	if err := db.MigrateDB(&domain.User{}); err != nil {
		panic(err)
	}
	fmt.Println("✅ successfully migrated dbs")

	// add dependency injection
	userRepo := repository.InitUserRepository(db)
	userSvc := service.InitUserService(userRepo)
	userHandler := handler.InitUserHandler(userSvc)

	// init router
	r, err := handler.InitRouter(conf.App, *userHandler)
	if err != nil {
		panic(err)
	}
	fmt.Println("✅ successfully initialized routes")

	// run application
	if err := r.Serve(conf.HTTP); err != nil {
		panic(err)
	}
}
