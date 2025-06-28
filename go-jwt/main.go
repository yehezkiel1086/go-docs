package main

import (
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
  router.GET("/ping", func(c *gin.Context) {
    c.JSON(200, gin.H{
      "message": "pong",
    })
  })
  router.Run() // by default: listen and serve on 0.0.0.0:8080
}
