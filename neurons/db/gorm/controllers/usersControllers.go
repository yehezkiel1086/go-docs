package controllers

import (
	"gorm-test/configs"
	"gorm-test/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllUsers(c *gin.Context) {
	db, err := configs.ConnDB()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// init users
	var users []models.ReadUsers

	// read users from db
	rows, err := db.Table("users").Select("id, name, age").Rows()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	defer rows.Close()

	for rows.Next() {
		var user models.ReadUsers
		if err := rows.Scan(&user.ID, &user.Name, &user.Age); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		users = append(users, user)
	}

	// response
	c.JSON(http.StatusOK, users)
}
