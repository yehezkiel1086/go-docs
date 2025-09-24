package handler

import (
	"go-socket/internal/core/port"

	"github.com/gin-gonic/gin"
)

type ClientHandler struct {
	svc port.ClientService
}

func InitClientHandler(svc port.ClientService) *ClientHandler {
	return &ClientHandler{
		svc: svc,
	}
}

type MessageReq struct {
	
}

func WriteMessage(c *gin.Context) {
	//
}
