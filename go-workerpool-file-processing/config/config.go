package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost         string
	DBPort         string
	DBUser         string
	DBPassword     string
	DBName         string
	DBSSLMode      string
	DBMaxConns     int
	DBMaxIdleConns int
	TotalWorker    int
	CSVFile        string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("=> error loading .env file")
	}

	return &Config{
		DBHost:         getEnv("DB_HOST", "localhost"),
		DBPort:         getEnv("DB_PORT", "5432"),
		DBUser:         getEnv("DB_USER", "postgres"),
		DBPassword:     getEnv("DB_PASSWORD", "postgres"),
		DBName:         getEnv("DB_NAME", "test"),
		DBSSLMode:      getEnv("DB_SSL_MODE", "disable"),
		DBMaxConns:     getEnvAsInt("DB_MAX_CONNS", 100),
		DBMaxIdleConns: getEnvAsInt("DB_MAX_IDLE_CONNS", 4),
		TotalWorker:    getEnvAsInt("TOTAL_WORKER", 50),
		CSVFile:        getEnv("CSV_FILE", "majestic_million.csv"),
	}
}

func (c *Config) DBConnString() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.DBHost, c.DBPort, c.DBUser, c.DBPassword, c.DBName, c.DBSSLMode,
	)
}

func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}

func getEnvAsInt(key string, fallback int) int {
	if val := os.Getenv(key); val != "" {
		n, err := strconv.Atoi(val)
		if err != nil {
			log.Fatalf("=> invalid value for %s: %v", key, err)
		}
		return n
	}
	return fallback
}