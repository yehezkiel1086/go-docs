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
	}

	App struct {
		Name string `json:"name"`
		Env string `json:"env"`
	}

	HTTP struct {
		Host string `json:"host"`
		Port string `json:"port"`
	}

	DB struct {
		Name string `json:"name"`
		User string `json:"user"`
		Pass string `json:"password"`
		Host string `json:"host"`
		Port string `json:"port"`
	}
)

func InitConfig() (*Container, error) {
	if os.Getenv("APP_ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			return &Container{}, errors.New("error loading .env file")
		}
	}

	App := &App{
		Name: os.Getenv("APP_NAME"),
		Env: os.Getenv("APP_ENV"),
	}

	HTTP := &HTTP{
		Host: os.Getenv("HTTP_HOST"),
		Port: os.Getenv("HTTP_PORT"),
	}

	DB := &DB{
		Name: os.Getenv("DB_NAME"),
		User: os.Getenv("DB_USER"),
		Pass: os.Getenv("DB_PASS"),
		Host: os.Getenv("DB_HOST"),
		Port: os.Getenv("DB_PORT"),
	}

	return &Container{
		App: App,
		HTTP: HTTP,
		DB: DB,
	}, nil
}
