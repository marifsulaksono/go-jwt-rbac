package config

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(conf *Config) {
	db, err := gorm.Open(mysql.Open(conf.DatabaseURL), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connect to database : %v", err)
	}

	DB = db
}
