package handler

import (
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

// UpdateProfile godoc
// @Summary Update user profile
// @Description Update the profile details of a user
// @Tags profile
// @Accept json
// @Produce json
// @Param profile body dto.UpdateProfileRequest true "Profile update request"
// @Success 200 {object} dto.MessageResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /profile [put]
func (h *ProfileHandler) UpdateProfile(c *gin.Context) {
	var req dto.UpdateProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
		return
	}

	err := h.profileService.UpdateProfile(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.MessageResponse{Message: "profile updated"})
}

// GetProfile godoc
// @Summary Get user profile
// @Description Get the profile details of a user by username
// @Tags profile
// @Produce json
// @Param username path string true "Username"
// @Success 200 {object} dto.ProfileResponse
// @Failure 404 {object} dto.ErrorResponse
// @Router /profile/{username} [get]
func (h *ProfileHandler) GetProfile(c *gin.Context) {
	username := c.Param("username")

	profile, err := h.profileService.GetProfile(username)
	if err != nil {
		c.JSON(http.StatusNotFound, dto.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.ProfileResponse{Username: profile.Username, Email: profile.Email})
}
