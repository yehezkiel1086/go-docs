package router

import (
	"fmt"
	"go-single-file-upload/config"
	"go-single-file-upload/controller"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Router struct {
	*gin.Engine
}

func InitRouter(
	conf *config.HTTP,
	fileController *controller.FileController,
) *Router {
	r := gin.New()

	allowedOrigins := strings.Split(conf.AllowedOrigins, ",")

	// CORS: allow frontend (Next.js)
	r.Use(cors.New(cors.Config{
		AllowOrigins:     allowedOrigins,
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// routes
	pb := r.Group("/api")

	pb.GET("/uploads", fileController.GetFiles)
	pb.POST("/uploads", fileController.UploadFile)

	return &Router{r}
}

func (r *Router) Serve(conf *config.HTTP) error {
	uri := fmt.Sprintf("%v:%v", conf.Host, conf.Port)
	if err := r.Run(uri); err != nil {
		return err
	}

	return nil
}
