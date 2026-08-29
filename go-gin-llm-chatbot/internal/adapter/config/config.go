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
		LLM *LLM
	}

	App struct {
		Name string
		Env string
	}

	HTTP struct {
		Host string
		Port string
	}

	LLM struct {
		APIKey string
		AIModel string
	}
)

func InitConfig() (*Container, error) {
	if os.Getenv("APP_ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			return nil, fmt.Errorf("error loading .env file: %v", err.Error())
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

	LLM := &LLM{
		APIKey: os.Getenv("GEMINI_API_KEY"),
		AIModel: os.Getenv("GEMINI_MODEL"),
	}

	return &Container{
		App: App,
		HTTP: HTTP,
		LLM: LLM,
	}, nil
}