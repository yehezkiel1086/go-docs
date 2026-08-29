package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type (
	Container struct {
		App  *App
		HTTP *HTTP
		DB   *DB
		Cloudinary *Cloudinary
	}

	App struct {
		Name string
		Env  string
	}

	HTTP struct {
		Host           string
		Port           string
		AllowedOrigins string
	}

	DB struct {
		Name string
		User string
		Pass string
		Host string
		Port string
	}

	Cloudinary struct {
		Name string
		Key string
		Secret string
		Folder string
	}
)

func InitConfig() (*Container, error) {
	if os.Getenv("HTTP_ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			return &Container{}, errors.New("failed to load .env file")
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
	}

	DB := &DB{
		Name: os.Getenv("DB_NAME"),
		User: os.Getenv("DB_USER"),
		Pass: os.Getenv("DB_PASS"),
		Host: os.Getenv("DB_HOST"),
		Port: os.Getenv("DB_PORT"),
	}

	Cloudinary := &Cloudinary{
		Name: os.Getenv("CLOUDINARY_NAME"),
		Key: os.Getenv("CLOUDINARY_API_KEY"),
		Secret: os.Getenv("CLOUDINARY_API_SECRET"),
		Folder: os.Getenv("CLOUDINARY_FOLDER"),
	}

	return &Container{
		App: App,
		HTTP: HTTP,
		DB: DB,
		Cloudinary: Cloudinary,
	}, nil
}
