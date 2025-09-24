package handler

import (
	"go-socket/internal/core/domain"
	"go-socket/internal/core/port"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HubHandler struct {
	svc port.HubService
}

func InitHubHandler(svc port.HubService) *HubHandler {
	return &HubHandler{
		svc: svc,
	}
}

func (h *HubHandler) CreateRoom(c *gin.Context) {
	// bind input
	var input *domain.CreateRoomReq
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "id and name are required.",
		})
		return
	}

	// create room
	res, err := h.svc.CreateRoom(c, &domain.Room{
		ID: input.ID,
		Name: input.Name,
		Clients: make(map[string]*domain.Client),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return		
	}

	// return response
	c.JSON(http.StatusCreated, &domain.CreateRoomRes{
		ID: res.ID,
		Name: res.Name,
	})
}
