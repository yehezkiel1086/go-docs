package controllers

import (
	"gin-jwt-test/models"
	"gin-jwt-test/setup"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllUsers(c *gin.Context) {
	// connect DB
	db, err := setup.ConnDB()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "DB Connection failed",
		})
		return
	}

	// get all users
	var users []models.UserOutput

	rows, err := db.Table("users").Select("users.id, users.username, roles.role").Joins("left join roles on roles.id = users.role_id").Rows()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	defer rows.Close()

	for rows.Next() {
    db.ScanRows(rows, &users)
  }

	c.JSON(http.StatusOK, users)
}

func GetUserByUsername(c *gin.Context) {
	// connect DB
	db, err := setup.ConnDB()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "DB Connection failed",
		})
		return
	}

	// get username from param
	userParam := c.Param("user")
	if userParam == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "User param is required",
		})
		return
	}

	// get user
	var user models.UserOutput

	if err := db.Table("users").Select("users.id, users.username, roles.role").Joins("left join roles on roles.id = users.role_id").Where("users.username = ?", userParam).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "User not found",
		})
		return
	}

	c.JSON(http.StatusOK, user)
}
