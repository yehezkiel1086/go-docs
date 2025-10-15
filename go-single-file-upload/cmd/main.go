package main

import (
	"context"
	"fmt"
	"go-single-file-upload/config"
	"go-single-file-upload/controller"
	"go-single-file-upload/model"
	"go-single-file-upload/router"
	"go-single-file-upload/storage"
)

func main() {
	// load .env
	conf, err := config.InitConfig()
	if err != nil {
		panic(err)
	}
	fmt.Println("✅ .env configs loaded successfully")

	ctx := context.Background()

	// init postgres db
	db, err := storage.InitDB(ctx, conf.DB)
	if err != nil {
		panic(err)
	}
	fmt.Println("✅ postgres db initialized successfully")

	// migrate dbs
	if err := db.Migrate(&model.FileRecord{}); err != nil {
		panic(err)
	}
	fmt.Println("✅ dbs migrated successfully")

	// init cloudinary
	cld, err := storage.InitCloudinary(ctx, conf.Cloudinary)
	if err != nil {
		panic(err)
	}
	fmt.Println("✅ cloudinary initialized successfully")

	// init controllers
	fileCtl := controller.InitFileController(db, cld, conf.Cloudinary)

	// init router
	r := router.InitRouter(
		conf.HTTP,
		fileCtl,
	)

	// serve api
	if err := r.Serve(conf.HTTP); err != nil {
		panic(err)
	}
}
