package handler

import (
	"auth-service/application/service"
	"auth-service/domain/repository"
	"auth-service/dto"
	"fmt"
	"net/http"
	"shared/utils"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	userRepository repository.UserRepository
	userService    *service.AuthService
}

func NewAuthHandler(userRepository repository.UserRepository) *AuthHandler {
	userService := service.NewAuthService(userRepository)

	return &AuthHandler{userRepository, &userService}
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.LoginRequest

	// Bind JSON to LoginRequest struct
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
		return
	}

	// Find user by username
	user, err := h.userRepository.FindByUsername(req.Username)
	if err != nil {
		c.JSON(http.StatusUnauthorized, dto.ErrorResponse{Error: "invalid username or password"})
		return
	}

	// Compare passwords
	if !utils.CheckPasswordHash(req.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, dto.ErrorResponse{Error: "invalid username or password"})
		return
	}

	// Generate JWT token
	token, err := utils.GenerateJWT(user.Username, user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "failed to generate token"})
		return
	}

	// Respond with token
	c.JSON(http.StatusOK, dto.LoginResponse{Token: token})
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req dto.RegisterRequest

	// Bind JSON to RegisterRequest struct
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
		return
	}

	// Save user
	if err := h.userService.Register(req); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "failed to register user"})
		return
	}

	c.Status(http.StatusNoContent)
}
