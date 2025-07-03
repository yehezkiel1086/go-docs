package middlewares

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		c.Set("message", "Hello middlewares!")

		// before request
		c.Next()

		// after request
		latency := time.Since(t)
		status := c.Writer.Status()

		log.Println(latency, status)
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