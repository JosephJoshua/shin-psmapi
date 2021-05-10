package db

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

// TODO: Use environment variables instead
const (
	DBHost     = "localhost"
	DBPort     = "5432"
	DBUser     = "postgres"
	DBPassword = "12345"
	DBName     = "psmapi"
)

func Init() {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		DBHost, DBPort, DBUser, DBPassword, DBName)

	tmp, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	db = tmp
}

func Migrate() {
	if db == nil {
		Init()
	}

	// Migrate models
	db.AutoMigrate()
}

func GetDB() *gorm.DB {
	if db == nil {
		Init()
	}

	return db
}
