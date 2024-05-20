package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Load() {
	// Load environment variables from .env file if available
	err := godotenv.Load()
	if err != nil {
		log.Printf(".env file not found, reading configuration from environment variables.")
	}

	if os.Getenv("MYSQL_DSN") == "" {
		log.Fatal("MYSQL_DSN not set in environment variables")
	}
}

func MySQLDSN() string {
	return os.Getenv("MYSQL_DSN")
}
