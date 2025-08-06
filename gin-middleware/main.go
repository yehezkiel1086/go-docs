package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		c.Set("example", "12345")

		c.Next()

		latency := time.Since(t)
		log.Print(latency)

		// access the status we are sending
		status := c.Writer.Status()
		log.Println(status)
	}
}

func KeyAccess() gin.HandlerFunc {
	return func(c *gin.Context) {
		key := c.Query("key")

		if key == "" {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
						"message": "unauthorized",
				})
				return
		}

		c.Next()
	}
}


func main() {
	r := gin.Default()

	r.Use(Logger(), KeyAccess())

	r.GET("/", func(c *gin.Context) {
		example := c.MustGet("example")
		c.JSON(http.StatusOK, gin.H{
			"message": example,
		})
	})

	r.Run()
}
