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
	userService *service.AuthService
}

func NewAuthHandler(userRepository repository.UserRepository) *AuthHandler {
	userService := service.NewAuthService(userRepository)

	return &AuthHandler{&userService}
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.LoginRequest

	// Bind JSON to LoginRequest struct
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
		return
	}

	user, err := h.userService.Login(req)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "failed to login to user"})
		return
	}

	token, err := utils.GenerateJWT(user.Username, user.ID)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "failed to login to user"})
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
