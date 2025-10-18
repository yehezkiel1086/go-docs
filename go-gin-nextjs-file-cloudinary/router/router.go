package router

import (
	"fmt"
	"go-gin-nextjs-file-cloudinary/config"
	"go-gin-nextjs-file-cloudinary/controller"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Router struct {
	r *gin.Engine
	conf *config.HTTP
}

func InitRouter(
	conf *config.HTTP,
	fileCtl *controller.FileController,
) *Router {
	// init router
	r := gin.New()

	// cors config
	r.Use(cors.New(cors.Config{
		AllowOrigins:     strings.Split(conf.AllowedOrigins, ","),
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// routes
	r.POST("/api/upload", fileCtl.UploadFile)
	r.GET("/api/uploads", fileCtl.GetUploads)

	// return router and config
	return &Router{
		r: r,
		conf: conf,
	}
}

func (r *Router) Serve() error {
	uri := fmt.Sprintf("%v:%v", r.conf.Host, r.conf.Port)
	return r.r.Run(uri)
}
