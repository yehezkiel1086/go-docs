package main

import (
	"gin-oauth2/configs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	configs.GetEnv()
}

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Testing",
		})
	})

	r.Run()
}