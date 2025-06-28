package main

import (
	"go-jwt/controllers"
	"go-jwt/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnDB()
  initializers.SyncDB()
}

func main() {
	router := gin.Default()

  router.POST("/signup", controllers.Signup)

  router.Run() // by default: listen and serve on 0.0.0.0:8080
}
