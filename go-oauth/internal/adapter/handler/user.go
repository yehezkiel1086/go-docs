package handler

import (
	"go-oauth/internal/core/domain"
	"go-oauth/internal/core/port"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	svc port.UserService
}

func InitUserHandler(svc port.UserService) *UserHandler {
	return &UserHandler{
		svc: svc,
	}
}

type RegisterReq struct {
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (uh *UserHandler) Register(c *gin.Context) {
	// bind input
	var input *RegisterReq
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return		
	}

	// register user
	_, err := uh.svc.Register(c, &domain.User{
		Email: input.Email,
		Password: input.Password,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// return response
	c.JSON(http.StatusCreated, gin.H{
		"message": "user registered successfully",
	})
}