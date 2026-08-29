package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yehezkiel1086/go-grpc-inventory-microservices/services/user-service/internal/adapter/config"
	"github.com/yehezkiel1086/go-grpc-inventory-microservices/services/user-service/internal/core/port"
)

type AuthHandler struct {
	jwtConf *config.JWT
	svc     port.AuthService
}

func NewAuthHandler(jwtConf *config.JWT, svc port.AuthService) *AuthHandler {
	return &AuthHandler{
		jwtConf: jwtConf,
		svc:     svc,
	}
}

type LoginReq struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req LoginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.svc.Login(c.Request.Context(), req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	duration, err := strconv.Atoi(h.jwtConf.Duration)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// set cookie
	c.SetCookie("jwt_token", token, duration*1000, "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{"jwt_token": token})
}
