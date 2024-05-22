package config

import (
	"log"
	"os"
)

func Load() {
	if os.Getenv("MYSQL_DSN") == "" {
		log.Fatal("MYSQL_DSN not set in environment variables")
	}
}

func MySQLDSN() string {
	return os.Getenv("MYSQL_DSN")
}
