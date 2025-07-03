package main

import (
	"gin-jwt-test/controllers"
	"gin-jwt-test/setup"

	"github.com/gin-gonic/gin"
)

func init() {
  setup.InitEnv()
  setup.MigrateDB()
}

func main() {
	r := gin.Default()

  pb := r.Group("/api") // public routes

  pb.GET("/users", controllers.GetAllUsers)
  pb.GET("/users/:user", controllers.GetUserByUsername)
  pb.POST("/roles", controllers.CreateNewRole)
  pb.POST("/register", controllers.Register)
  pb.POST("/login", controllers.Authenticate)

  r.Run() // listen and serve on 0.0.0.0:<PORT>
}
