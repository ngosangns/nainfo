package main

import (
	"log"
	"os"
	_ "profile-service/docs" // Import Swagger docs
	"profile-service/infrastructure/router"
	"shared/config"
)

// @title Profile Service API
// @version 1.0
// @description This is a sample profile service.
// @host localhost:8001
// @BasePath /

func main() {
	config.Load()           // Load configuration
	r := router.NewRouter() // Create router
	addr := os.Getenv("PROFILE_SERVICE_ADDRESS")
	if addr == "" {
		addr = ":8001"
	}

	log.Fatal(r.Run(addr))
}
