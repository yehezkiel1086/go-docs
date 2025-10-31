package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type (
	Container struct {
		HTTP *HTTP
		Rabbitmq *Rabbitmq
	}

	HTTP struct {
		Host string
		Port string
	}

	Rabbitmq struct {
		Host string
		Port string
		User string		
		Pass string
	}
)

func InitConfig() (*Container, error) {
	if os.Getenv("APP_ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			return &Container{}, errors.New("error loading .env file")
		}
	}

	HTTP := &HTTP{
		Host: os.Getenv("HTTP_HOST"),
		Port: os.Getenv("HTTP_PORT"),
	}

	Rabbitmq := &Rabbitmq{
		Host: os.Getenv("RABBITMQ_HOST"),
		Port: os.Getenv("RABBITMQ_PORT"),
		User: os.Getenv("RABBITMQ_USER"),
		Pass: os.Getenv("RABBITMQ_PASS"),
	}
	
	return &Container{
		HTTP: HTTP,
		Rabbitmq: Rabbitmq,
	}, nil
}