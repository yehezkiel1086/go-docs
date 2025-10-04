package handler

import (
	"fmt"
	"go-oauth/internal/adapter/config"

	"github.com/gin-gonic/gin"
)

type Router struct {
	*gin.Engine
}

func InitRouter(
	conf *config.App,
	userHandler UserHandler,
) (*Router, error) {
	r := gin.New()

	// public routes
	pb := r.Group("/api/v1")
	pb.POST("/register", userHandler.Register)

	// return router object
	return &Router{r}, nil
}

func (r *Router) Serve(conf *config.HTTP) error {
	uri := fmt.Sprintf("%v:%v", conf.Host, conf.Port)
	return r.Run(uri)
}
