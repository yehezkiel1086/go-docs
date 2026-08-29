package main

import (
	"fmt"
	"go-gin-next-oauth/internal/handler"
	"go-gin-next-oauth/internal/model"
	"go-gin-next-oauth/internal/service"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	godotenv.Load()

	// Connect PostgreSQL
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect database:", err)
	}

	// Auto migrate user
	db.AutoMigrate(&model.User{})

	// Handlers
	authService := service.NewAuthService(db)
	authHandler := handler.NewAuthHandler(authService)

	r := gin.Default()

	r.GET("/api/v1/auth/google/login", authHandler.GoogleLogin)
	r.GET("/api/v1/auth/google/callback", authHandler.GoogleCallback)

	port := os.Getenv("HTTP_PORT")
	r.Run(":" + port)
}
