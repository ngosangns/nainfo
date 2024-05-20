package main

import (
	"auth-service/infrastructure/router"
	"log"
	"os"
	"shared/config"
)

// @title Auth Service API
// @version 1.0
// @description This is a sample auth service.
// @host localhost:8000
// @BasePath /
func main() {
	config.Load()           // Load configuration
	r := router.NewRouter() // Create router
	addr := os.Getenv("AUTH_SERVICE_ADDRESS")
	if addr == "" {
		addr = ":8000"
	}

	log.Fatal(r.Run(addr))
}
