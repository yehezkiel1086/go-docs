package main

import (
	"gorm-test/configs"
	"gorm-test/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string
	Age int
}

func main() {
	// load env
	if err := configs.LoadEnv(); err != nil {
		panic(err)
	}

	// init gin
	router := gin.Default()

	router.GET("/api/users", controllers.GetAllUsers)

	router.Run() // listen and serve on 0.0.0.0:<PORT>
}
