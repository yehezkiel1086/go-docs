package main

import (
	"gin-oauth/configs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	configs.LoadEnv()
}

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello Gin OAuth!",
		})
	})

	r.Run()
}
