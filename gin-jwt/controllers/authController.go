package controllers

import (
	"gin-jwt-test/models"
	"gin-jwt-test/setup"
	"net/http"

	"github.com/gin-gonic/gin"
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
	var input models.UserInput

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

	// auth success: return generated token
	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful!",
	})
}
