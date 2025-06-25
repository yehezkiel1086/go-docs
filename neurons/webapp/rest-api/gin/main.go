package main

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Login struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type User struct {
	Id uint8 `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Age uint8 `json:"age"`
	Job string `json:"job"`
}

func main() {
	r := gin.Default()

	r.POST("/login", loginUser)

	r.GET("/users", getUsers)
	r.GET("/users/:id", getUser)

	r.Run(":8080")
}

func loginUser(c *gin.Context) {
	var login Login
	err := c.ShouldBind(&login)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// manual check
	// if login.Username == "" || login.Password == "" {
	// 	c.JSON(http.StatusUnauthorized, gin.H{
	// 		"error": "Username and password are required.",
	// 	})
	// 	return
	// }

	c.JSON(http.StatusOK, gin.H{
		"message": "User logged in!",
	})
}

func getUsers(c *gin.Context) {
	content, err := os.ReadFile("users.json")
	if err != nil {
		panic(err)
	}

	var users []User

	err = json.Unmarshal(content, &users)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, users)
}

func getUser(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(err)
	}
	
	content, err := os.ReadFile("users.json")
	if err != nil {
		panic(err)
	}

	var users []User

	err = json.Unmarshal(content, &users)
	if err != nil {
		panic(err)
	}
	
	for _, user := range users {
		if user.Id == uint8(userId) {
			c.JSON(http.StatusOK, user)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"message": "User not found.",
	})
}
