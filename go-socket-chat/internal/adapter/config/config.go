package config

import (
	"os"

	"github.com/joho/godotenv"
)

type (
	Container struct {
		App  *App
		HTTP *HTTP
		DB   *DB
		JWT *JWT
	}

	App struct {
		Name string
		Env  string
	}

	HTTP struct {
		Host string
		Port string
	}

	DB struct {
		Connection string
		Host     string
		Port     string
		Name     string
		User     string
		Password string
	}

	JWT struct {
		Secret string
		Duration string
	}
)

func New() (*Container, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
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
		Connection: os.Getenv("DB_CONNECTION"),
		Host: os.Getenv("DB_HOST"),
		Port: os.Getenv("DB_PORT"),
		Name: os.Getenv("DB_NAME"),
		User: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
	}

	JWT := &JWT{
		Secret: os.Getenv("JWT_SECRET"),
		Duration: os.Getenv("JWT_DURATION"),
	}

	return &Container{
		App,
		HTTP,
		DB,
		JWT,
	}, nil
}