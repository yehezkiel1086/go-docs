package controllers

import (
	"errors"
	"gin-oauth/configs"
	"gin-oauth/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateNewRole(c *gin.Context) {
	// connect DB
	db, err := configs.ConnDB()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "DB Connection failed",
		})
		return
	}

	// check empty input
	role := &models.Role{}

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
