package handler

import (
	"go-socket/internal/core/domain"
	"go-socket/internal/core/port"
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

type UserReq struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (uh *UserHandler) Register(c *gin.Context) {
	// bind input
	var input *UserReq
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "username, email and password are required.",
		})
		return
	}

	// register user
	_, err := uh.svc.CreateUser(c, &domain.User{
		Username: input.Username,
		Email: input.Email,
		Password: input.Password,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// response
	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully.",
	})
}
