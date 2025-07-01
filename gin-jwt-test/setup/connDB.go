package setup

import (
	"fmt"
	"gin-jwt-test/models"

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

	fmt.Println("Connected to DB!")

	// Migrate DB schemas
	db.AutoMigrate(&models.User{})
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