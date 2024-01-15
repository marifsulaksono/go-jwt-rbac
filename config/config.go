package config

import (
	"fmt"
	"os"
)

const (
	dbHost     = "DB_HOST"
	dbPort     = "DB_PORT"
	dbUser     = "DB_USER"
	dbPass     = "DB_PASS"
	dbName     = "DB_NAME"
	serverPort = "SERVER_PORT"
)

type Config struct {
	ServerPort  string
	DatabaseURL string
}

func GetConfig() *Config {
	return &Config{
		ServerPort:  os.Getenv(serverPort),
		DatabaseURL: fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", os.Getenv(dbUser), os.Getenv(dbPass), os.Getenv(dbHost), os.Getenv(dbPort), os.Getenv(dbName)),
	}
}
