package handler

import (
	"go-socket/internal/adapter/config"

	"github.com/gin-gonic/gin"
)

type Router struct {
	*gin.Engine
}

func InitRouter(
	config *config.HTTP,
	userHandler UserHandler,
) (*Router, error) {
	r := gin.New()

	v1 := r.Group("/api/v1") // public routes
	v1.POST("/register", userHandler.Register)

	return &Router{r}, nil
}

func (r *Router) Serve(uri string) error {
	return r.Run(uri)
}
