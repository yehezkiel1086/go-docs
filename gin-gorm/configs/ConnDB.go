package configs

import (
	"fmt"
	"gin-gorm/models"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnDB() (*gorm.DB, error) {
	DB_HOST := os.Getenv("DB_HOST")
	DB_USER := os.Getenv("DB_USER")
	DB_PASS := os.Getenv("DB_PASS")
	DB_NAME := os.Getenv("DB_NAME")
	DB_PORT := os.Getenv("DB_PORT")
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Asia/Shanghai", DB_HOST, DB_USER, DB_PASS, DB_NAME, DB_PORT)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	return db, err
}

func MigrateDB() {
	db, err := ConnDB()
	if err != nil {
		panic(err)
	}

	// migrate models
	db.AutoMigrate(&models.User{})
}
