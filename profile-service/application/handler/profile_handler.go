package handler

import (
	"fmt"
	"net/http"
	"profile-service/application/service"
	"profile-service/domain/repository"
	"profile-service/dto"

	"github.com/gin-gonic/gin"
)

type ProfileHandler struct {
	profileService service.ProfileService
}

func NewProfileHandler(repo repository.ProfileRepository) *ProfileHandler {
	return &ProfileHandler{profileService: service.NewProfileService(repo)}
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

	c.JSON(http.StatusOK, dto.ProfileResponse{Username: profile.Username, Email: profile.Email})
}
