package handler

import (
	"context"
	"go-gin-next-oauth/internal/service"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	service *service.AuthService
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{
		service: service.NewAuthService(),
	}
}

func (h *AuthHandler) GoogleLogin(c *gin.Context) {
	url := h.service.GetGoogleLoginURL()
	c.Redirect(http.StatusTemporaryRedirect, url)
}

func (h *AuthHandler) GoogleCallback(c *gin.Context) {
	code := c.Query("code")
	if code == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing code"})
		return
	}

	userInfo, token, err := h.service.HandleGoogleCallback(context.Background(), code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	frontendRedirect := os.Getenv("HTTP_CLIENT_REDIRECT_URL") + "?token=" + token
	c.Redirect(http.StatusTemporaryRedirect, frontendRedirect)

	_ = userInfo // optional: log or save user
}
