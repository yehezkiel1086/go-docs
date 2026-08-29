package handler

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yehezkiel1086/go-jwt-auth/internal/adapter/config"
	"github.com/yehezkiel1086/go-jwt-auth/internal/adapter/storage/redis"
	"github.com/yehezkiel1086/go-jwt-auth/internal/core/domain"
	"github.com/yehezkiel1086/go-jwt-auth/internal/core/util"
)

type Router struct {
	r *gin.Engine
}

func New(
	jwtConfig *config.JWT,
	corsConfig *config.CORS,
	cache *redis.Redis,
	userHandler *UserHandler,
	authHandler *AuthHandler,
) (*Router, error) {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(SecureHeadersMiddleware())
	r.Use(CORSMiddleware(corsConfig))

	rateLimiter := util.NewRateLimiter(cache.GetClient(), 60, time.Minute)
	strictLimiter := util.NewRateLimiter(cache.GetClient(), 10, time.Minute)

	pb := r.Group("/api/v1", RateLimitMiddleware(rateLimiter))
	au := pb.Group("/auth")
	us := pb.Group("/users", AuthMiddleware(jwtConfig))
	ad := pb.Group("/admin", AuthMiddleware(jwtConfig), RoleMiddleware(domain.AdminRole))

	pb.POST("/register", userHandler.CreateUser)

	au.POST("/login", RateLimitMiddleware(strictLimiter), authHandler.Login)
	au.POST("/refresh", RateLimitMiddleware(strictLimiter), authHandler.RefreshToken)
	au.POST("/logout", AuthMiddleware(jwtConfig), authHandler.Logout)

	us.GET("/:id", SelfOrAdminMiddleware(), userHandler.GetUserByID)
	us.PUT("/:id", SelfOrAdminMiddleware(), userHandler.UpdateUserByID)
	us.DELETE("/:id", SelfOrAdminMiddleware(), userHandler.DeleteUserByID)

	ad.GET("/users", userHandler.GetUsers)
	ad.GET("/users/:id", userHandler.GetUserByID)
	ad.DELETE("/users/:id", userHandler.DeleteUserByID)

	return &Router{r: r}, nil
}

func (r *Router) Run(conf *config.HTTP) error {
	return r.r.Run(conf.Host + ":" + conf.Port)
}