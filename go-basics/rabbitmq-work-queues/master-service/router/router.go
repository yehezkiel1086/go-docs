package router

import (
	"go-rabbitmq-work-queues/config"
	"go-rabbitmq-work-queues/controller"

	"github.com/gin-gonic/gin"
)

type Router struct {
	r *gin.Engine
}

func InitRouter(
	ctl *controller.Controller,
) *Router {
	r := gin.Default()

	pb := r.Group("/api/v1")

	pb.POST("/publish", ctl.PublishMessage)

	return &Router{
		r: r,
	}
}

func (r *Router) Serve(conf *config.HTTP) error {
	return r.r.Run(conf.Host + ":" + conf.Port)
}