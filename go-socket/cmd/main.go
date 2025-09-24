package main

import (
	"context"
	"fmt"
	"go-socket/internal/adapter/config"
	storage "go-socket/internal/adapter/storage/postgres"
	"go-socket/internal/core/domain"
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

	// router config

	// start server
}
