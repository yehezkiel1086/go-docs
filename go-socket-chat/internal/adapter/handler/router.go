package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/yehezkiel1086/go-socket-chat/internal/adapter/config"
)

type Router struct {
	r *gin.Engine
}

func NewRouter(
	userHandler *UserHandler,
) *Router {
	r := gin.New()

	r.Use(gin.Logger())

	// rbac
	pb := r.Group("/api/v1")

	// user routes
	pb.POST("/register", userHandler.Register)
	pb.POST("/login", userHandler.Login)

	return &Router{r}
}

func (r *Router) Run(conf *config.HTTP) error {
	uri := fmt.Sprintf("%s:%s", conf.Host, conf.Port)
	return r.r.Run(uri)
}
