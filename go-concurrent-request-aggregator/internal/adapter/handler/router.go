package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/yehezkiel1086/go-concurrent-request-aggregator/internal/adapter/config"
)

type Router struct {
	r *gin.Engine
}

func NewRouter(
	DashboardHandler *DashboardHandler,
) (*Router) {
	r := gin.New()

	r.Use(gin.Logger())

	pb := r.Group("/api/v1")

	pb.GET("/dashboard", DashboardHandler.GetDashboard)

	return &Router{r}
}

func (r *Router) Run(conf *config.HTTP) error {
	uri := fmt.Sprintf("%s:%s", conf.Host, conf.Port)
	return r.r.Run(uri)
}
