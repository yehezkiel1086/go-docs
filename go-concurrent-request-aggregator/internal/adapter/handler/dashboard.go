package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yehezkiel1086/go-concurrent-request-aggregator/internal/core/port"
)

type DashboardHandler struct {
	svc port.DashboardService
}

func NewDashboardHandler(svc port.DashboardService) *DashboardHandler {
	return &DashboardHandler{
		svc,
	}
}

func (h *DashboardHandler) GetDashboard(c *gin.Context) {
	userId := c.Query("userId")
	if userId == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "userId query is required",
		})
		return
	}

	dashboard, err := h.svc.GetDashboard(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dashboard)
}
