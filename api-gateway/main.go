package main

import (
	"log"
	"net/http/httputil"
	"net/url"
	"os"
	"shared/middleware"
	"strconv"

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
	// Modify the request URL to include the username and userID in query params
	profileGroup.Use(func(c *gin.Context) {
		username, _ := c.Get("username")
		userID, _ := c.Get("userID")

		if c.Request.URL.RawQuery != "" {
			c.Request.URL.RawQuery += "&"
		}
		c.Request.URL.RawQuery += "username=" + username.(string) + "&userID=" + strconv.FormatFloat(userID.(float64), 'f', -1, 64)

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
