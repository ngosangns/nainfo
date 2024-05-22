package main

import (
	"database/sql"
	"log"
	"os"
	"profile-service/infrastructure/grpc"
	"profile-service/infrastructure/persistence"
	"profile-service/infrastructure/router"
	"shared/config"
)

// @title Profile Service API
// @version 1.0
// @description This is a sample profile service.
// @host localhost:8001
// @BasePath /

func main() {
	config.Load() // Load configuration

	db, err := sql.Open("mysql", config.MySQLDSN())
	if err != nil {
		panic(err)
	}
	profileRepository := persistence.NewMySQLProfileRepository(db)

	r := router.NewRouter(db, profileRepository) // Create router
	addr := os.Getenv("PROFILE_SERVICE_ADDRESS")
	if addr == "" {
		addr = ":8001"
	}

	// Run the gRPC server
	if err := grpc.RunGRPCServer(profileRepository); err != nil {
		panic(err)
	}

	// Run the HTTP server
	log.Fatal(r.Run(addr))
}
