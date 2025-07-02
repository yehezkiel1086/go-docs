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

  pub := r.Group("/api") // public routes

  pub.POST("/register", controllers.Register)
  pub.POST("/login", controllers.Authenticate)

  r.Run() // listen and serve on 0.0.0.0:<PORT>
}
