package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yehezkiel1086/go-jwt-auth/internal/adapter/config"
	"github.com/yehezkiel1086/go-jwt-auth/internal/core/port"
)

type AuthHandler struct {
	svc  port.AuthService
	conf *config.JWT
	app  *config.App
}

func NewAuthHandler(app *config.App, conf *config.JWT, svc port.AuthService) *AuthHandler {
	return &AuthHandler{svc: svc, conf: conf, app: app}
}

type loginReq struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func (h *AuthHandler) isProduction() bool {
	return h.app.Env == "production"
}

func (h *AuthHandler) cookieHost() string {
	if h.isProduction() {
		return h.app.Host
	}
	return ""
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req loginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	accessToken, refreshToken, err := h.svc.Login(c.Request.Context(), req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	secure := h.isProduction()
	host := h.cookieHost()

	c.SetCookie("access_token", accessToken, 15*60, "/", host, secure, true)
	c.SetCookie("refresh_token", refreshToken, 7*24*60*60, "/api/v1/auth/refresh", host, secure, true)

	c.JSON(http.StatusOK, gin.H{"message": "login successful"})
}

func (h *AuthHandler) RefreshToken(c *gin.Context) {
	refreshToken, err := c.Cookie("refresh_token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "refresh token not found"})
		return
	}

	accessToken, newRefreshToken, err := h.svc.RefreshToken(c.Request.Context(), refreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	secure := h.isProduction()
	host := h.cookieHost()

	c.SetCookie("access_token", accessToken, 15*60, "/", host, secure, true)
	c.SetCookie("refresh_token", newRefreshToken, 7*24*60*60, "/api/v1/auth/refresh", host, secure, true)

	c.JSON(http.StatusOK, gin.H{"message": "token refreshed"})
}

func (h *AuthHandler) Logout(c *gin.Context) {
	accessToken, err := c.Cookie("access_token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "access token not found"})
		return
	}

	if err := h.svc.Logout(c.Request.Context(), accessToken); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	secure := h.isProduction()
	host := h.cookieHost()

	c.SetCookie("access_token", "", -1, "/", host, secure, true)
	c.SetCookie("refresh_token", "", -1, "/api/v1/auth/refresh", host, secure, true)

	c.JSON(http.StatusOK, gin.H{"message": "logout successful"})
}