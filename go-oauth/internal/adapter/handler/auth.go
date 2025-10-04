package handler

import (
	"go-oauth/internal/core/domain"
	"go-oauth/internal/core/port"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	svc port.AuthService
}

func InitAuthHandler(svc port.AuthService) *AuthHandler {
	return &AuthHandler{
		svc: svc,
	}
}

func (ah *AuthHandler) Login(c *gin.Context) {
	// bind input
	var input *domain.LoginReq
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "email and password are required",
		})
		return
	}

	// login
	if _, err := ah.svc.Login(c, input); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid credentials",
		})
		return
	}

	// return response
	c.JSON(http.StatusOK, gin.H{
		"message": "user logged in successfully",
	})
}
