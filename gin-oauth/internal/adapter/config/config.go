package config

import (
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
		Host string `json:"host"`
		Name string `json:"name"`
		Username string `json:"username"`
		Password string `json:"password"`
	}
)

func InitConfig() (*Container, error) {
	if os.Getenv("APP_ENV") != "production" {
		if err := godotenv.Load(); err != nil {
			return &Container{}, err
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
		Host: os.Getenv("DB_HOST"),
		Name: os.Getenv("DB_NAME"),
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
	}

	return &Container{
		App: App,
		HTTP: HTTP,
		DB: DB,
	}, nil
}
