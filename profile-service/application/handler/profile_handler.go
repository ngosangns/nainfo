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
	var req dto.UpdateProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
		return
	}

	err := h.profileService.UpdateProfile(req)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.MessageResponse{Message: "profile updated"})
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
