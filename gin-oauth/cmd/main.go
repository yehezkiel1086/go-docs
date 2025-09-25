package main

import (
	"context"
	"fmt"
	"go-oauth2/internal/adapter/config"
	"go-oauth2/internal/adapter/handler"
	"go-oauth2/internal/adapter/storage/postgres"
	"go-oauth2/internal/adapter/storage/postgres/repository"
	"go-oauth2/internal/core/domain"
	"go-oauth2/internal/core/service"
)

func main() {
	// init configs
	conf, err := config.InitConfig()
	if err != nil {
		panic(err)
	}
	fmt.Println("Configuration imported successfully ✅")

	// init database
	ctx := context.Background()
	db, err := postgres.InitDB(ctx, conf.DB)
	if err != nil {
		panic(err)
	}
	fmt.Println("DB connection success ✅")

	// migrate db
	if err := db.Migrate(&domain.User{}); err != nil {
		panic(err)
	}
	fmt.Println("DB migrations success ✅")

	// dependency injection
	userRepo := repository.InitUserRepository(db)
	userSvc := service.InitUserService(userRepo)
	userHandler := handler.InitUserHandler(userSvc)

	// routing
	r, err := handler.InitRouter(conf.HTTP, *userHandler)
	if err != nil {
		panic(err)
	}

	// run server
	uri := fmt.Sprintf("%v:%v", conf.HTTP.Host, conf.HTTP.Port)
	r.Serve(uri)
}
