package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/yehezkiel1086/go-grpc-inventory-microservices/services/user-service/internal/adapter/config"
)

type Router struct {
	r *gin.Engine
}

func NewRouter(
	jwtConf *config.JWT,
	userHandler *UserHandler,
	authHandler *AuthHandler,
) (*Router, error) {
	r := gin.New()

	r.Use(gin.Logger(), gin.Recovery())

	// rbac
	pb := r.Group("/api/v1")
	au := pb.Group("/", AuthMiddleware(jwtConf))
	// ad := au.Group("/", AdminOnly())

	// user routes
	pb.POST("/register", userHandler.RegisterUser)
	au.GET("/users/:id", AuthenticatedOnly(), userHandler.GetUserByID)
	au.PUT("/users/:id", AuthenticatedOnly(), userHandler.UpdateUser)
	au.DELETE("/users/:id", AuthenticatedOnly(), userHandler.DeleteUser)

	// auth routes
	pb.POST("/login", authHandler.Login)

	return &Router{
		r: r,
	}, nil
}

func (r *Router) Run(httpConf *config.HTTP) error {
	addr := httpConf.Host + ":" + httpConf.Port
	return r.r.Run(addr)
}
