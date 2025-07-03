package middlewares

import (
	"gin-jwt-test/models"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// get cookie
		cookie, err := c.Cookie("jwt_token")
		if err != nil {
			if err == http.ErrNoCookie {
				c.JSON(http.StatusUnauthorized, gin.H{
					// "error": "Unauthorized",
					"error": err.Error(),
				})
			} else {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
			}

			c.Abort()
			return
		}

		jwtKey := os.Getenv("ACCESS_TOKEN")

		// check cookie's token with ACCESS_TOKEN
		claims := &models.Claims{}
		token, err := jwt.ParseWithClaims(cookie, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				c.JSON(http.StatusUnauthorized, gin.H{
					// "error": "Unauthorized",
					"error": err.Error(),
				})
			} else {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
			}

			c.Abort()
			return
		}

		// check if token is still valid
		if !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				// "error": "Unauthorized",
					"error": err.Error(),
			})

			c.Abort()
			return
		}

		c.Next()
	}
}

func AdminHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		//
	}
}