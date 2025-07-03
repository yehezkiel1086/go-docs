package controllers

import (
	"errors"
	"gin-jwt-test/models"
	"gin-jwt-test/setup"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateNewRole(c *gin.Context) {
	// connect DB
	db, err := setup.ConnDB()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "DB Connection failed",
		})
		return
	}

	// check empty input
	var role models.Role

	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// check role existed
	if err := db.Where("role = ?", role.Role).First(&models.Role{}).Error; !errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Role already existed",
		})
		return
	}

	// create new role
	db.Create(&role)

	c.JSON(http.StatusCreated, gin.H{
		"message": "New role created!",
	})
}
