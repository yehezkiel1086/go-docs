package main

import (
	"context"
	"fmt"

	"github.com/yehezkiel1086/go-photo-api-redis/config"
	"github.com/yehezkiel1086/go-photo-api-redis/controller"
	"github.com/yehezkiel1086/go-photo-api-redis/router"
	"github.com/yehezkiel1086/go-photo-api-redis/storage/redis"
)

func main() {
	conf, err := config.InitConfig()
	if err != nil {
		panic(err)
	}
	fmt.Println("✅ .env configs initialized successfully")

	ctx := context.Background()

	// init redis
	redis, err := redis.InitRedis(ctx, conf.Redis)
	if err != nil {
		panic(err)
	}
	fmt.Println("✅ redis initialized successfully")

	// dependency injection
	photoController := controller.InitPhotoController(ctx, redis)

	// init router
	r, err := router.InitRouter(
		photoController,
	)
	if err != nil {
		panic(err)
	}
	fmt.Println("✅ router initialized successfully")

	// serve
	if err := r.Serve(conf.HTTP); err != nil {
		panic(err)
	}
}
