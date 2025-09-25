package main

import (
	"context"
	"fmt"
	"go-oauth2/internal/adapter/config"
	"go-oauth2/internal/adapter/storage/postgres"
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
	_, err = postgres.InitDB(ctx, conf.DB)
	if err != nil {
		panic(err)
	}
	fmt.Println("DB connection success ✅")
}
