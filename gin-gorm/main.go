package main

import (
	"gin-gorm/configs"
	"gin-gorm/controllers"

	"github.com/gin-gonic/gin"
)

func init() {
	configs.GetEnv()
	configs.MigrateDB()
}

func main() {
	route := gin.Default()

	// routes
	r := route.Group("/api/v1")
	r.GET("/users", controllers.GetAllUsers)
	r.POST("/users", controllers.CreateUser)

	route.Run()
}
