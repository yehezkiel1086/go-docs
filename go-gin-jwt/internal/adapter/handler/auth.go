package handler

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
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

type JWTClaims struct {
	Username string `json:"username"`
	Role domain.Role `json:"role"`
	jwt.RegisteredClaims
}

func (ah *AuthHandler) Login(c *gin.Context) {
	// bind request input
	var input *LoginInput
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	// authenticate
	user, err := ah.svc.Login(c, &domain.User{
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
	expirationTime := time.Now().Add(time.Hour * 24)
	claims := JWTClaims{
		Username: user.Username,
		Role: user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			// Also fixed dates can be used for the NumericDate
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			Issuer:    "admin",
		},
	}

	// get .env secret
	secret := os.Getenv("JWT_SECRET")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(secret))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// set jwt token to cookie
	c.SetCookie(
    "jwt_token",
    ss,
    int(time.Until(expirationTime).Seconds()), // maxAge in seconds
    "/",
    "",
    false, // set to true if you use HTTPS
    true,  // httpOnly for security
	)
	
	// success response
	c.JSON(http.StatusOK, gin.H{
		"message": "Login success.",
	})
}
