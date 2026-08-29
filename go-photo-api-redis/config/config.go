package config

import (
	"os"

	"github.com/joho/godotenv"
)

type (
	Container struct {
		App  *App
		HTTP *HTTP
		// DB *DB
		Redis *Redis
	}

	App struct {
		Name string
		Env  string
	}

	HTTP struct {
		Host string
		Port string
	}

	// DB struct {
	// 	Host     string
	// 	Port     string
	// 	Name string
	// 	User     string
	// 	Password string
	// }

	Redis struct {
		URL string
		Password string
		DB string
	}
)

func InitConfig() (*Container, error) {
	if os.Getenv("APP_ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
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

	// DB := &DB{
	// 	Host:     os.Getenv("DB_HOST"),
	// 	Port:     os.Getenv("DB_PORT"),
	// 	Name: os.Getenv("DB_NAME"),
	// 	User:     os.Getenv("DB_USER"),
	// 	Password: os.Getenv("DB_PASS"),
	// }

	Redis := &Redis{
		URL: os.Getenv("REDIS_URL"),
		Password: os.Getenv("REDIS_PASS"),
		DB: os.Getenv("REDIS_DB"),
	}

	return &Container{
		App: App,
		HTTP: HTTP,
		// DB: DB,
		Redis: Redis,
	}, nil
}