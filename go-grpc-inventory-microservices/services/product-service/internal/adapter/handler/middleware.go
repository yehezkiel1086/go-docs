package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yehezkiel1086/go-grpc-inventory-microservices/services/product-service/internal/adapter/config"
	"github.com/yehezkiel1086/go-grpc-inventory-microservices/services/product-service/internal/core/domain"
	"github.com/yehezkiel1086/go-grpc-inventory-microservices/services/product-service/internal/core/util"
)

type contextKey string

const ContextKeyUser contextKey = "user"

func AuthMiddleware(jwtConf *config.JWT) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("jwt_token")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized - no token"})
			c.Abort()
			return
		}

		claims, err := util.ParseJWTToken(jwtConf.Secret, token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized - invalid token"})
			c.Abort()
			return
		}

		c.Set(string(ContextKeyUser), claims)
		c.Next()
	}
}

func AdminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, exists := c.Get(string(ContextKeyUser))
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}

		claims, ok := user.(*util.JWTClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}

		if claims.Role != domain.RoleAdmin {
			c.JSON(http.StatusForbidden, gin.H{"error": "forbidden - admin only"})
			c.Abort()
			return
		}

		c.Next()
	}
}

func AuthenticatedOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, exists := c.Get(string(ContextKeyUser))
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}
		c.Next()
	}
}
