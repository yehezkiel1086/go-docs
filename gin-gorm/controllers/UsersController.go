package controllers

import (
	"gin-gorm/configs"
	"gin-gorm/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user models.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	db, err := configs.ConnDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "DB Connection failed",
		})
		return		
	}

	db.Create(&user)

	c.JSON(http.StatusCreated, user)
}

func GetAllUsers(c *gin.Context) {
	db, err := configs.ConnDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "DB Connection Error",
		})
		return
	}

	rows, err := db.Table("users").Rows()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	
	defer rows.Close()

	users := []models.User{}

	for rows.Next() {
		db.ScanRows(rows, &users)
	}

	c.JSON(http.StatusOK, users)
}
