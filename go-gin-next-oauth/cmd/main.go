package main

import (
	"go-gin-next-oauth/internal/handler"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	r := gin.Default()
	authHandler := handler.NewAuthHandler()

	r.GET("/api/v1/auth/google/login", authHandler.GoogleLogin)
	r.GET("/api/v1/auth/google/callback", authHandler.GoogleCallback)

	port := os.Getenv("HTTP_PORT")
	if port == "" {
		port = "3500"
	}

	log.Printf("Server running at :%s", port)
	r.Run(":" + port)
}
