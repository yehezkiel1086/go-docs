package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/yehezkiel1086/go-photo-api-redis/config"
	"github.com/yehezkiel1086/go-photo-api-redis/controller"
)

type Router struct {
	r *gin.Engine
}

func InitRouter(
	photoController *controller.PhotoController,
) (*Router, error) {
	r := gin.Default()

	pb := r.Group("/api/v1")

	pb.GET("/photos", photoController.GetPhotos)
	pb.GET("/photos/:albumId", photoController.GetPhotosByAlbumID)
	
	return &Router{r: r}, nil
}

func (r *Router) Serve(conf *config.HTTP) error {
	uri := fmt.Sprintf("%s:%s", conf.Host, conf.Port)
	return r.r.Run(uri)
}
