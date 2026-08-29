package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type (
	Container struct {
		App *App
		HTTP *HTTP
		RabbitMQ *RabbitMQ
	}

	App struct {
		Name string
		Env string
	}

	HTTP struct {
		Host string
		Port string
	}

	RabbitMQ struct {
		Host string
		Port string
		User string
		Pass string
	}
)

func InitConfig() (*Container, error) {
	if os.Getenv("") != "production" {
		err := godotenv.Load()
		if err != nil {
			return nil, fmt.Errorf("error loading .env file: %v", err.Error())
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

	Rabbitmq := &RabbitMQ{
		Host: os.Getenv("RABBITMQ_HOST"),
		Port: os.Getenv("RABBITMQ_PORT"),
		User: os.Getenv("RABBITMQ_USER"),
		Pass: os.Getenv("RABBITMQ_PASS"),
	}

	return &Container{
		App: App,
		HTTP: HTTP,
		RabbitMQ: Rabbitmq,
	}, nil
}