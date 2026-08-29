package config

import (
	"os"

	"github.com/joho/godotenv"
)

type (
	Container struct {
		App          *App
		HTTP         *HTTP
		ProductRPC   *ProductRPC
		InventoryRPC *InventoryRPC
		Rabbitmq     *Rabbitmq
		DB           *DB
		JWT          *JWT
	}

	App struct {
		Name string
		Env  string
	}

	HTTP struct {
		Host string
		Port string
	}

	Rabbitmq struct {
		Host     string
		Port     string
		User     string
		Password string
	}

	ProductRPC struct {
		Host string
		Port string
	}

	InventoryRPC struct {
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
		Port: os.Getenv("HTTP_PORT_ORDER"),
	}

	ProductRPC := &ProductRPC{
		Host: os.Getenv("HTTP_HOST"),
		Port: os.Getenv("RPC_PORT_PRODUCT"),
	}

	InventoryRPC := &InventoryRPC{
		Host: os.Getenv("HTTP_HOST"),
		Port: os.Getenv("RPC_PORT_INVENTORY"),
	}

	DB := &DB{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Name:     os.Getenv("DB_NAME"),
	}

	Rabbitmq := &Rabbitmq{
		Host:     os.Getenv("RABBITMQ_HOST"),
		Port:     os.Getenv("RABBITMQ_PORT"),
		User:     os.Getenv("RABBITMQ_USER"),
		Password: os.Getenv("RABBITMQ_PASSWORD"),
	}

	JWT := &JWT{
		Secret:   os.Getenv("JWT_SECRET"),
		Duration: os.Getenv("JWT_DURATION"),
	}

	return &Container{
		App:          App,
		HTTP:         HTTP,
		ProductRPC:   ProductRPC,
		InventoryRPC: InventoryRPC,
		Rabbitmq:     Rabbitmq,
		DB:           DB,
		JWT:          JWT,
	}, nil
}
