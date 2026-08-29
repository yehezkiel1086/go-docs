package config

import (
	"os"

	"github.com/joho/godotenv"
)

type (
	Container struct {
		App  *App
		HTTP *HTTP
		API  *API
	}

	App struct {
		Name string
		Env  string
	}

	HTTP struct {
		Host string
		Port string
	}

	API struct {
		Base string
	}
)

func New() (*Container, error) {
	if os.Getenv("APP_ENV") != "production" {
		if err := godotenv.Load(); err != nil {
			return nil, err
		}
	}

	App := &App{
		Name: os.Getenv("APP_NAME"),
		Env:  os.Getenv("APP_ENV"),
	}

	HTTP := &HTTP{
		Host: os.Getenv("HTTP_HOST"),
		Port: os.Getenv("HTTP_PORT"),
	}

	API := &API{
		Base: os.Getenv("API_BASE"),
	}

	return &Container{
		App,
		HTTP,
		API,
	}, nil
}
