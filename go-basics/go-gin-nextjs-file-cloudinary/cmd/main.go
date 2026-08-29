package main

import (
	"context"
	"fmt"
	"go-gin-nextjs-file-cloudinary/config"
	"go-gin-nextjs-file-cloudinary/config/cloudinary"
	"go-gin-nextjs-file-cloudinary/config/postgres"
	"go-gin-nextjs-file-cloudinary/controller"
	"go-gin-nextjs-file-cloudinary/model"
	"go-gin-nextjs-file-cloudinary/router"
)

func main() {
	// init .env configs
	conf, err := config.InitConfig()
	if err != nil {
		panic(err)
	}
	fmt.Println("✅ .env configs imported successfully")

	ctx := context.Background()

	// init postgres db
	db, err := postgres.InitDB(ctx, conf.DB)
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
	cld, err := cloudinary.InitCloudinary(conf.Cloudinary)
	if err != nil {
		panic(err)
	}
	fmt.Println("✅ cloudinary initialized successfully")

	// init controllers
	fileCtl := controller.InitFileController(conf.Cloudinary, cld, db)

	// init router
	r := router.InitRouter(
		conf.HTTP,
		fileCtl,
	)
	fmt.Println("✅ router initialized successfully")

	// run server
	if err := r.Serve(); err != nil {
		panic(err)
	}
}
