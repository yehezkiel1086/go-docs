package setup

import (
	"fmt"
	"gin-jwt-test/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func MigrateDB() {
	// get .env DB values
	dsn := getDBEnv()
	
	// connect DB
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	// ✅ Enable uuid-ossp extension (safe even if already exists)
	if err := db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`).Error; err != nil {
		log.Fatal("Failed to enable uuid-ossp extension:", err)
	}

	fmt.Println("Connected to DB!")

	// Migrate DB schemas
	db.AutoMigrate(&models.User{}, &models.Role{})
}

func ConnDB() (*gorm.DB, error) {
	// get .env DB values
	dsn := getDBEnv()
	
	// connect DB
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, nil
}