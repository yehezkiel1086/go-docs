package router

import (
	"rabbitmq-hello/producer-api/config"
	"rabbitmq-hello/producer-api/controller"

	"github.com/gin-gonic/gin"
)

type Router struct {
	r *gin.Engine
}

func InitRouter(
	helloController *controller.HelloController,
) (*Router) {
	r := gin.Default()

	pb := r.Group("/api/v1")

	pb.POST("/hello", helloController.Hello)

	return &Router{
		r: r,
	}
}

func (r *Router) Run(conf *config.HTTP) error {
	uri := conf.Host + ":" + conf.Port
	return r.r.Run(uri)
}
