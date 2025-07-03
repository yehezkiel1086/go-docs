package main

import (
	"go-middleware/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// middlewares implementations
	r.Use(middlewares.Logger(), middlewares.KeyAccess())

	r.GET("/v1/hello", func(c *gin.Context) {
		msg := c.MustGet("message").(string)

		c.JSON(http.StatusOK, gin.H{
			"message": msg,
		})
	})

	r.Run()
}
