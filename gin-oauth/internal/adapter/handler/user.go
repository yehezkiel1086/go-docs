package handler

import (
	"go-oauth2/internal/core/domain"
	"go-oauth2/internal/core/port"
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
	Username string `json:"username"`
	Password string `json:"password"`
}

func (uh *UserHandler) Register(c *gin.Context) {
	// bind input
	var input *RegisterReq
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Username and password are required.",
		})
		return
	}

	// create new user
	_, err := uh.svc.CreateUser(c, &domain.User{
		Username: input.Username,
		Password: input.Password,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Username and password are required.",
		})
		return
	}

	// return response
	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully.",
	})
}
