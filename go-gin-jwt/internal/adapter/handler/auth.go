package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yehezkiel1086/go-docs/go-gin-jwt/internal/core/domain"
	"github.com/yehezkiel1086/go-docs/go-gin-jwt/internal/core/port"
)

type AuthHandler struct {
	svc port.AuthService
}

func InitAuthHandler(svc port.AuthService) *AuthHandler {
	return &AuthHandler{
		svc: svc,
	}
}

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (ah *AuthHandler) Login(c *gin.Context) {
	// bind request input
	var input *LoginInput
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	// authenticate
	_, err := ah.svc.Login(c, &domain.User{
		Username: input.Username,
		Password: input.Password,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid username or password.",
		})
		return
	}

	// generate jwt token

	// set jwt token to cookie

	// success response
	c.JSON(http.StatusOK, gin.H{
		"message": "Login success.",
	})
}
