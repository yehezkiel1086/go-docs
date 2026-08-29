package config

import (
	"os"

	"github.com/joho/godotenv"
)

type (
	Container struct {
		App       *App
		HTTP      *HTTP
		RPC       *RPC
		DB        *DB
		JWT       *JWT
		Inventory *Inventory
	}

	App struct {
		Name string
		Env  string
	}

	HTTP struct {
		Host string
		Port string
	}

	RPC struct {
		Host string
		Port string
	}

	DB struct {
		Name     string
		User     string
		Password string
		Host     string
		Port     string
	}

	JWT struct {
		Secret   string
		Duration string
	}

	Inventory struct {
		Host string
		Port string
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
		Port: os.Getenv("HTTP_PORT_PRODUCT"),
	}

	RPC := &RPC{
		Host: os.Getenv("HTTP_HOST"),
		Port: os.Getenv("RPC_PORT_PRODUCT"),
	}

	DB := &DB{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Name:     os.Getenv("DB_NAME"),
	}

	JWT := &JWT{
		Secret:   os.Getenv("JWT_SECRET"),
		Duration: os.Getenv("JWT_DURATION"),
	}

	Inventory := &Inventory{
		Host: os.Getenv("HTTP_HOST"),
		Port: os.Getenv("RPC_PORT_INVENTORY"),
	}

	return &Container{
		App:       App,
		HTTP:      HTTP,
		RPC:       RPC,
		DB:        DB,
		JWT:       JWT,
		Inventory: Inventory,
	}, nil
}
