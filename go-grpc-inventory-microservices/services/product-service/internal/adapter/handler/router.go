package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/yehezkiel1086/go-grpc-inventory-microservices/services/product-service/internal/adapter/config"
)

type Router struct {
	r *gin.Engine
}

func NewRouter(
	jwtConf *config.JWT,
	productHandler *ProductHandler,
) (*Router, error) {
	r := gin.New()

	r.Use(gin.Logger(), gin.Recovery())

	// rbac
	pb := r.Group("/api/v1")
	au := pb.Group("/", AuthMiddleware(jwtConf))
	ad := au.Group("/", AdminOnly())

	// public product routes
	pb.GET("/products", productHandler.GetAllProducts)
	pb.GET("/products/:id", productHandler.GetProductByID)

	// admin product routes
	ad.POST("/products", productHandler.CreateProduct)
	ad.PUT("/products/:id", productHandler.UpdateProduct)
	ad.DELETE("/products/:id", productHandler.DeleteProduct)

	return &Router{
		r: r,
	}, nil
}

func (r *Router) Run(httpConf *config.HTTP) error {
	addr := httpConf.Host + ":" + httpConf.Port
	return r.r.Run(addr)
}
