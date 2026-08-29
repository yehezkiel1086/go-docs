package handler

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/yehezkiel1086/go-jwt-auth/internal/adapter/config"
	"github.com/yehezkiel1086/go-jwt-auth/internal/core/domain"
	"github.com/yehezkiel1086/go-jwt-auth/internal/core/util"
)

func AuthMiddleware(conf *config.JWT) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := c.Cookie("access_token")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized: no token provided"})
			c.Abort()
			return
		}

		claims, err := util.ParseToken(tokenString, conf.AccessTokenSecret)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized: " + err.Error()})
			c.Abort()
			return
		}

		c.Set("userID", claims.UserID)
		c.Set("userEmail", claims.Email)
		c.Set("userRole", claims.Role)
		c.Next()
	}
}

func RoleMiddleware(role domain.Role) gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole, exists := c.Get("userRole")
		if !exists || userRole.(domain.Role) != role {
			c.JSON(http.StatusForbidden, gin.H{"error": "forbidden: insufficient privileges"})
			c.Abort()
			return
		}
		c.Next()
	}
}

func SelfOrAdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		paramID, err := strconv.ParseUint(c.Param("id"), 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
			c.Abort()
			return
		}

		userRole, _ := c.Get("userRole")
		if userRole.(domain.Role) == domain.AdminRole {
			c.Next()
			return
		}

		userID, exists := c.Get("userID")
		if !exists || userID.(uint) != uint(paramID) {
			c.JSON(http.StatusForbidden, gin.H{"error": "forbidden: access denied"})
			c.Abort()
			return
		}

		c.Next()
	}
}

func RateLimitMiddleware(rl *util.RateLimiter) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := c.ClientIP()
		if userID, exists := c.Get("userID"); exists {
			key = strconv.FormatUint(uint64(userID.(uint)), 10)
		}

		allowed, remaining, err := rl.Allow(c.Request.Context(), key)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "rate limiter error"})
			c.Abort()
			return
		}

		c.Header("X-RateLimit-Remaining", strconv.Itoa(remaining))

		if !allowed {
			c.JSON(http.StatusTooManyRequests, gin.H{"error": "too many requests, please try again later"})
			c.Abort()
			return
		}

		c.Next()
	}
}

func SecureHeadersMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// prevent clickjacking
		c.Header("X-Frame-Options", "DENY")
		// prevent MIME sniffing
		c.Header("X-Content-Type-Options", "nosniff")
		// enable XSS protection in older browsers
		c.Header("X-XSS-Protection", "1; mode=block")
		// force HTTPS for 1 year, include subdomains
		c.Header("Strict-Transport-Security", "max-age=31536000; includeSubDomains")
		// restrict referrer information
		c.Header("Referrer-Policy", "strict-origin-when-cross-origin")
		// disable browser features not needed by the API
		c.Header("Permissions-Policy", "geolocation=(), microphone=(), camera=()")
		// content security policy for API responses
		c.Header("Content-Security-Policy", "default-src 'none'; frame-ancestors 'none'")
		// remove server fingerprint
		c.Header("Server", "")
		c.Next()
	}
}

func CORSMiddleware(conf *config.CORS) gin.HandlerFunc {
	allowedOrigins := make(map[string]struct{}, len(conf.AllowedOrigins))
	for _, o := range conf.AllowedOrigins {
		allowedOrigins[strings.TrimSpace(o)] = struct{}{}
	}

	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")

		if _, ok := allowedOrigins[origin]; ok {
			c.Header("Access-Control-Allow-Origin", origin)
			c.Header("Access-Control-Allow-Credentials", "true")
			c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
			c.Header("Access-Control-Max-Age", "86400")
			c.Header("Vary", "Origin")
		}

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}