package controllers

import (
	"gin-jwt-test/models"
	"gin-jwt-test/setup"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Authenticate(c *gin.Context) {
	// connect DB
	db, err := setup.ConnDB()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "DB Connection failed",
		})
		return
	}

	// check empty user input
	var input models.LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Username and password are required",
		})
		return
	}

	// check username
	var user, emptyUser models.User
	db.First(&user, "username = ?", input.Username)

	if user == emptyUser {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid credentials",
		})
		return
	}

	// check password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid credentials",
		})		
		return
	}

	// generate token
	expirationTime := time.Now().Add(5 * time.Minute)

	claims := &models.Claims{
		Username: input.Username,
		RoleID:     user.RoleID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	jwtKey := os.Getenv("ACCESS_TOKEN")

	tokenString, err := token.SignedString([]byte(jwtKey))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Set token string ke dalam cookie response
	maxAge := 1000 * 60 * 5
	c.SetCookie("jwt_token", tokenString, maxAge, "/api", "127.0.0.1:3500", false, true)

	// auth success: return generated token
	c.JSON(http.StatusOK, gin.H{
		"jwt_token": tokenString,
	})
}
