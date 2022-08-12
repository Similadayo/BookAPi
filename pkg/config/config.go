package config

import (
	"os"
	"log"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db * gorm.DB
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}
}

func Connect () {
	var (
		DB_DRIVER = os.Getenv("DB_DRIVER")
		DB_PASSWORD = os.Getenv("DB_PASSWORD")
		DB_USER = os.Getenv("DB_USER")
	)
	
	d, err := gorm.Open(DB_DRIVER, DB_USER+":"+DB_PASSWORD+"@/store?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDb () *gorm.DB {
	return db 
} 