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
	hubHandler HubHandler,
) (*Router, error) {
	r := gin.New()

	// auth, register and user routes
	v1 := r.Group("/api/v1") // public routes
	v1.POST("/register", userHandler.Register)

	// web socket routes
	ws := v1.Group("/ws")
	ws.POST("/rooms", hubHandler.CreateRoom)

	return &Router{r}, nil
}

func (r *Router) Serve(uri string) error {
	return r.Run(uri)
}
