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

	// check empty user input
	var input models.RegisInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Username, roleId and password are required",
		})
		return
	}

	// encrypt password
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Password encryption failed",
		})
		return
	}

	input.Password = string(hashedPwd)

	// create new user
	if err := db.Create(&models.User{
		Username: input.Username,
		Password: input.Password,
		RoleID: input.RoleID,
	}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "User creation failed",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Registration successful!",
	})
}
