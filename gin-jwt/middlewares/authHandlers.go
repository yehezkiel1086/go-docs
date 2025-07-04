package middlewares

import (
	"gin-jwt-test/models"
	"gin-jwt-test/setup"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

func AuthHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// get cookie
		cookie, err := c.Cookie("jwt_token")
		if err != nil {
			if err == http.ErrNoCookie {
				c.JSON(http.StatusUnauthorized, gin.H{
					"error": "Unauthorized",
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
					"error": "Unauthorized",
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
				"error": "Unauthorized",
			})

			c.Abort()
			return
		}

		c.Set("role_id", claims.RoleID)

		c.Next()
	}
}

func AdminHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// get role id
		roleId := c.MustGet("role_id")

		// connect db
		db, err := setup.ConnDB()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "DB Connection failed",
			})

			c.Abort()
			return
		}

		// get roles
		var role models.Role

		if err := db.Where("id = ?", roleId).First(&role).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				c.JSON(http.StatusUnauthorized, gin.H{
					"error": "Unauthorized",
				})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": err.Error(),
				})
			}
			
			c.Abort()
			return
		}

		if role.Role != "admin" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized",
			})
	
			c.Abort()
			return
		}

		c.Next()
	}
}