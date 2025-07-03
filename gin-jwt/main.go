package main

import (
	"gin-jwt-test/controllers"
	"gin-jwt-test/middlewares"
	"gin-jwt-test/setup"

	"github.com/gin-gonic/gin"
)

func init() {
  setup.InitEnv()
  setup.MigrateDB()
}

func main() {
	r := gin.Default()

  pb := r.Group("/api") // public routes

  // auth routes
  pb.POST("/register", controllers.Register)
  pb.POST("/login", controllers.Authenticate)

  // authenticated routes
  au := pb.Group("/v1", middlewares.AuthHandler())
  au.GET("/users/:user", controllers.GetUserByUsername)

  // admin only routes
  adm := au.Group("/admin", middlewares.AdminHandler())
  adm.GET("/users", controllers.GetAllUsers)
  adm.POST("/roles", controllers.CreateNewRole)
  
  r.Run() // listen and serve on 0.0.0.0:<PORT>
}
