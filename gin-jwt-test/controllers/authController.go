package controllers

import (
	"gin-jwt-test/models"
	"gin-jwt-test/setup"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	// connect DB
	db, err := setup.ConnDB()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "DB Connection failed",
		})
		return
	}

	// validate user input
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Username and password are required",
		})
		return
	}

	// encrypt password
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Password encryption failed",
		})
		return
	}

	user.Password = string(hashedPwd)

	// create new user
	db.Create(&user)

	c.JSON(http.StatusCreated, gin.H{
		"message": "Registration successful!",
	})
}
