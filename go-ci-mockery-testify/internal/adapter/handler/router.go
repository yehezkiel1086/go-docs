package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/yehezkiel1086/go-rest-mockery-testify/internal/adapter/config"
)

type Router struct {
	r *gin.Engine
}

func New(
	taskHandler *TaskHandler,
) *Router {
	r := gin.New()

	// router groups
	pb := r.Group("/api/v1")

	pb.POST("/tasks", taskHandler.CreateTask)
	pb.GET("/tasks/:id", taskHandler.GetTaskByID)
	pb.GET("/tasks", taskHandler.GetTasks)
	pb.PUT("/tasks/:id", taskHandler.UpdateTask)
	pb.DELETE("/tasks/:id", taskHandler.DeleteTask)

	return &Router{
		r,
	}
}

func (r *Router) Run(conf *config.HTTP) error {
	uri := conf.Host + ":" + conf.Port

	return r.r.Run(uri)
}
