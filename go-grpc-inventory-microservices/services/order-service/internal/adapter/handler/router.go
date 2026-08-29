package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/yehezkiel1086/go-grpc-inventory-microservices/services/order-service/internal/adapter/config"
)

type Router struct {
	r *gin.Engine
}

func NewRouter(
	jwtConf *config.JWT,
	orderHandler *OrderHandler,
) (*Router, error) {
	r := gin.New()

	r.Use(gin.Logger(), gin.Recovery())

	// rbac
	pb := r.Group("/api/v1")
	au := pb.Group("/", AuthMiddleware(jwtConf))
	ad := au.Group("/", AdminOnly())

	// public order routes

	// user order routes
	au.POST("/orders", orderHandler.CreateOrder)
	au.GET("/orders", orderHandler.GetMyOrders)
	au.GET("/orders/:id", orderHandler.GetOrderByID)

	// admin order routes
	ad.PATCH("/orders/:id", orderHandler.UpdateOrder)
	ad.DELETE("/orders/:id", orderHandler.DeleteOrder)

	return &Router{
		r: r,
	}, nil
}

func (r *Router) Run(httpConf *config.HTTP) error {
	addr := httpConf.Host + ":" + httpConf.Port
	return r.r.Run(addr)
}
