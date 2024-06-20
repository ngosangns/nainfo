package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"profile-service/application/service"
	"profile-service/domain/repository"
	"profile-service/dto"
	"shared/config"

	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type ProfileHandler struct {
	profileService service.ProfileService
	minioClient    *minio.Client
}

func NewProfileHandler(repo repository.ProfileRepository) *ProfileHandler {
	// MinIO Setup
	endpoint := config.MinIOHost()             // Replace with your MinIO endpoint
	accessKeyID := config.MinIOAccessKey()     // Replace with your MinIO access key
	secretAccessKey := config.MinIOSecretKey() // Replace with your MinIO secret key
	useSSL := false

	// Initialize MinIO client object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalln(err)
	}

	return &ProfileHandler{profileService: service.NewProfileService(repo), minioClient: minioClient}
}

func (h *ProfileHandler) UpdateProfile(c *gin.Context) {
	username := c.Query("username")
	if username == "" {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "username is required"})
		return
	}

	var req dto.UpdateProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
		return
	}
	req.Username = username

	err := h.profileService.UpdateOrCreateProfile(req)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *ProfileHandler) GetProfile(c *gin.Context) {
	username := c.Query("username")

	profile, err := h.profileService.GetProfile(username)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusNotFound, dto.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.ProfileResponse{
		Username:    profile.Username,
		Email:       profile.Email,
		Name:        profile.Name,
		Description: profile.Description,
		Address:     profile.Address,
		Facebook:    profile.Facebook,
		LinkedIn:    profile.LinkedIn,
		GitHub:      profile.GitHub,
	})
}

func (h *ProfileHandler) UploadAvatar(c *gin.Context) {
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
	_, err = h.minioClient.PutObject(
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
	url, err := h.minioClient.PresignedGetObject(context.Background(), config.MinIOBucketName(), objectName, 7*24*60*60, nil) // Expire after 7 days
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("generate presigned url err: %s", err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Avatar uploaded successfully",
		"url":     url.String(),
	})
}
