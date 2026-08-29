package config

import (
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type (
	Container struct {
		App   *App
		HTTP  *HTTP
		DB    *DB
		Redis *Redis
		JWT   *JWT
		CORS  *CORS
	}

	App struct {
		Name string
		Env  string
		Host string
	}

	HTTP struct {
		Host string
		Port string
	}

	DB struct {
		Host     string
		Port     string
		User     string
		Password string
		Name     string
	}

	Redis struct {
		Host     string
		Port     string
		Password string
		DB       string
	}

	JWT struct {
		Host                 string
		AccessTokenSecret    string
		RefreshTokenSecret   string
		AccessTokenDuration  string
		RefreshTokenDuration string
	}

	CORS struct {
		AllowedOrigins []string
	}
)

func New() (*Container, error) {
	if os.Getenv("APP_ENV") != "production" {
		if err := godotenv.Load(); err != nil {
			return nil, err
		}
	}

	origins := strings.Split(os.Getenv("CORS_ALLOWED_ORIGINS"), ",")
	if len(origins) == 0 || origins[0] == "" {
		origins = []string{"http://localhost:3000"}
	}

	return &Container{
		App: &App{
			Name: os.Getenv("APP_NAME"),
			Env:  os.Getenv("APP_ENV"),
			Host: os.Getenv("HTTP_HOST"),
		},
		HTTP: &HTTP{
			Host: os.Getenv("HTTP_HOST"),
			Port: os.Getenv("HTTP_PORT"),
		},
		DB: &DB{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Name:     os.Getenv("DB_NAME"),
		},
		Redis: &Redis{
			Host:     os.Getenv("REDIS_HOST"),
			Port:     os.Getenv("REDIS_PORT"),
			Password: os.Getenv("REDIS_PASSWORD"),
			DB:       os.Getenv("REDIS_DB"),
		},
		JWT: &JWT{
			Host:                 os.Getenv("HTTP_HOST"),
			AccessTokenSecret:    os.Getenv("ACCESS_TOKEN_SECRET"),
			RefreshTokenSecret:   os.Getenv("REFRESH_TOKEN_SECRET"),
			AccessTokenDuration:  os.Getenv("ACCESS_TOKEN_DURATION"),
			RefreshTokenDuration: os.Getenv("REFRESH_TOKEN_DURATION"),
		},
		CORS: &CORS{
			AllowedOrigins: origins,
		},
	}, nil
}