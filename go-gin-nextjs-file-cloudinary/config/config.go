package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type (
	Container struct {
		App *App
		HTTP *HTTP
		DB *DB
		Cloudinary *Cloudinary
	}

	App struct {
		Name string
		Env string
	}

	HTTP struct {
		Host string
		Port string
		AllowedOrigins string
		ClientURL string
	}

	DB struct {
		Host string
		Port string
		Name string
		User string
		Pass string
	}

	Cloudinary struct {
		Name string
		APIKey string
		APISecret string
		Folder string
	}
)

func InitConfig() (*Container, error) {
	if os.Getenv("APP_ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			return &Container{}, errors.New("failed to load .env config file")
		}
	}

	App := &App{
		Name: os.Getenv("APP_NAME"),
		Env: os.Getenv("APP_ENV"),
	}

	HTTP := &HTTP{
		Host: os.Getenv("HTTP_HOST"),
		Port: os.Getenv("HTTP_PORT"),
		AllowedOrigins: os.Getenv("HTTP_ALLOWED_ORIGINS"),
		ClientURL: os.Getenv("HTTP_CLIENT_URL"),
	}

	DB := &DB{
		Host: os.Getenv("DB_HOST"),
		Port: os.Getenv("DB_PORT"),
		Name: os.Getenv("DB_NAME"),
		User: os.Getenv("DB_USER"),
		Pass: os.Getenv("DB_PASS"),
	}

	Cloudinary := &Cloudinary{
		Name: os.Getenv("CLOUDINARY_NAME"),
		APIKey: os.Getenv("CLOUDINARY_API_KEY"),
		APISecret: os.Getenv("CLOUDINARY_API_SECRET"),
		Folder: os.Getenv("CLOUDINARY_FOLDER"),
	}

	return &Container{
		App: App,
		HTTP: HTTP,
		DB: DB,
		Cloudinary: Cloudinary,
	}, nil
}
