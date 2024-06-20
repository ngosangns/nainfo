package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"shared/config"
	"shared/middleware"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var minioClient *minio.Client

func uploadAvatar(c *gin.Context) {
	// Retrieve the file from the request
	file, err := c.FormFile("avatar")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}

	// Open the file
	openedFile, err := file.Open()
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("open file err: %s", err.Error()))
		return
	}
	defer openedFile.Close()

	// Upload the file to MinIO
	objectName := file.Filename
	bucketName := config.MinIOBucketName()
	_, err = minioClient.PutObject(
		context.Background(),
		bucketName,
		objectName,
		openedFile,
		file.Size,
		minio.PutObjectOptions{ContentType: file.Header.Get("Content-Type")},
	)
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("upload to minio err: %s", err.Error()))
		return
	}

	// Generate the public URL
	url, err := minioClient.PresignedGetObject(context.Background(), config.MinIOBucketName(), objectName, 7*24*time.Hour, nil) // Expire after 7 days
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("generate presigned url err: %s", err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Avatar uploaded successfully",
		"url":     "/storage/get/avatar/" + strings.Split(url.Path, "/")[2],
	})
}

func main() {
	r := gin.Default()

	// Enable CORS
	corsConfig := cors.Config{
		AllowHeaders: []string{
			"Authorization",
			"Content-Type",
		},
	}
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowMethods = append(corsConfig.AllowMethods, "OPTIONS", "PUT", "POST", "GET", "DELETE")
	r.Use(cors.New(corsConfig))

	// MinIO Setup
	minIOEndpoint := config.MinIOHost()             // Replace with your MinIO endpoint
	minIOAccessKeyID := config.MinIOAccessKey()     // Replace with your MinIO access key
	minIOSecretAccessKey := config.MinIOSecretKey() // Replace with your MinIO secret key
	useSSL := false

	// Initialize MinIO client object.
	var err error
	minioClient, err = minio.New(minIOEndpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(minIOAccessKeyID, minIOSecretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatal(err)
	}

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

	r.POST("/storage/post/avatar", uploadAvatar)
	r.GET("/storage/get/avatar/:objectName", func(c *gin.Context) {
		objectName := c.Param("objectName")
		bucketName := config.MinIOBucketName()

		// Get object information (optional, but useful for metadata)
		objectInfo, err := minioClient.StatObject(context.Background(), bucketName, objectName, minio.StatObjectOptions{})
		if err != nil {
			c.String(http.StatusNotFound, "Image not found")
			return
		}

		// Get object data (image)
		object, err := minioClient.GetObject(context.Background(), bucketName, objectName, minio.GetObjectOptions{})
		if err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("Error fetching object: %v", err))
			return
		}
		defer object.Close()

		// Set appropriate headers for image response
		c.Header("Content-Disposition", fmt.Sprintf("inline; filename=%s", objectName))
		c.Header("Content-Type", objectInfo.ContentType)
		c.Header("Content-Length", fmt.Sprintf("%d", objectInfo.Size))
		c.Header("Cache-Control", fmt.Sprintf("max-age=%d", 7*24*60*60)) // Cache for 7 days

		// Stream the image data to the client
		if _, err = io.Copy(c.Writer, object); err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("Error streaming object: %v", err))
			return
		}
	})

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
