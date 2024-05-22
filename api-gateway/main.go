package main

import (
	"log"
	"net/http"
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

	authGroup := r.Group("/auth")
	authGroup.Any("/*proxyPath", reverseProxy(authURL))

	profileGroup := r.Group("/profile")
	// Profile routes need authentication
	profileGroup.Use(middleware.AuthMiddleware())
	// Modify the request URL to include the username in query params
	profileGroup.Use(func(c *gin.Context) {
		username, _ := c.Get("username")

		// Get all query parameters into an object
		values, err := url.ParseQuery(c.Request.URL.RawQuery)
		if err != nil {
			log.Printf("Error parsing query parameters: %v", err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		// Set the username field
		values.Set("username", username.(string))

		// Format the query string back
		c.Request.URL.RawQuery = values.Encode()

		c.Next()
	})
	profileGroup.Any("/*proxyPath", reverseProxy(profileURL))

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
