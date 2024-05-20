package main

import (
	"log"
	"net/http/httputil"
	"net/url"
	"os"
	"shared/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Enable CORS
	r.Use(cors.Default())

	// Apply authentication middleware for protected routes
	r.Use(middleware.AuthMiddleware())

	authServiceAddress := os.Getenv("AUTH_SERVICE_ADDRESS")
	profileServiceAddress := os.Getenv("PROFILE_SERVICE_ADDRESS")

	if authServiceAddress == "" {
		authServiceAddress = "http://localhost:8000"
	}
	if profileServiceAddress == "" {
		profileServiceAddress = "http://localhost:8001"
	}

	authURL, _ := url.Parse(authServiceAddress)
	profileURL, _ := url.Parse(profileServiceAddress)

	r.Any("/auth/*proxyPath", reverseProxy(authURL))
	r.Any("/profile/*proxyPath", reverseProxy(profileURL))

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Could not start API Gateway: %v", err)
	}
}

func reverseProxy(target *url.URL) gin.HandlerFunc {
	proxy := httputil.NewSingleHostReverseProxy(target)
	return func(c *gin.Context) {
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}
