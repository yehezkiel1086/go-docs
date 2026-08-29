package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/yehezkiel1086/go-gin-llm-chatbot/internal/adapter/config"
)

type Route struct {
	r *gin.Engine
}

func InitRoute(
	promptHandler *PromptHandler,
) *Route {
	r := gin.New()

	// route groupings
	pb := r.Group("/api/v1")

	// public routes
	pb.POST("/prompt/text-to-text", promptHandler.TextToTextPrompt)
	pb.POST("/prompt/stream-text-to-text", promptHandler.StreamTextToTextPrompt)
	pb.POST("/prompt/image-to-text", promptHandler.ImageToTextPrompt)

	// protected routes

	return &Route{
		r: r,
	}
}

func (r *Route) Serve(conf *config.HTTP) error {
	uri := conf.Host + ":" + conf.Port
	return r.r.Run(uri)
}
