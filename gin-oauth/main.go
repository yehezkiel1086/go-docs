package main

import (
	"gin-oauth/configs"
	"gin-oauth/controllers"

	"github.com/gin-gonic/gin"
)

func init() {
	configs.LoadEnv()
	configs.MigrateDB()
}

func main() {
	r := gin.Default()

	pb := r.Group("/api") // public routes

	// authenticated routes
  au := pb.Group("/v1")

	// admin only routes
	adm := au.Group("/admin")
	adm.POST("/roles", controllers.CreateNewRole)

	r.Run()
}
