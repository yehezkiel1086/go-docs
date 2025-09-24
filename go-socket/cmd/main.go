package main

import (
	"context"
	"fmt"
	"go-socket/internal/adapter/config"
	"go-socket/internal/adapter/handler"
	storage "go-socket/internal/adapter/storage/postgres"
	"go-socket/internal/adapter/storage/postgres/repository"
	"go-socket/internal/core/domain"
	"go-socket/internal/core/service"
)

func main() {
	conf, err := config.InitConfig()
	if err != nil {
		panic(err)
	}
	fmt.Println("Configs initialized. ✅")

	// init database
	ctx := context.Background()
	db, err := storage.InitDB(ctx, conf.DB)
	if err != nil {
		panic(err)
	}
	fmt.Println("Database connected. ✅")	

	// migrate tables
	if err := db.Migrate(&domain.User{}); err != nil {
		panic(err)
	}
	fmt.Println("Tables migrated. ✅")

	// dependency injection
	userRepo := repository.InitUserRepository(db)
	userSvc := service.InitUserService(userRepo)
	userHandler := handler.InitUserHandler(userSvc)

	// router config
	r, err := handler.InitRouter(
		conf.HTTP,
		*userHandler,
	)
	if err != nil {
		panic(err)
	}

	// start server
	uri := fmt.Sprintf("%v:%v", conf.HTTP.Host, conf.HTTP.Port)
	r.Serve(uri)
}
