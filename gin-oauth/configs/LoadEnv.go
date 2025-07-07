package configs

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// s3Bucket := os.Getenv("S3_BUCKET")
	// secretKey := os.Getenv("SECRET_KEY")

	// now do something with s3 or whatever
}