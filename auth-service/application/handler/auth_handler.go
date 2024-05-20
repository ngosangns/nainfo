package handler

import (
	"auth-service/domain/model"
	"auth-service/domain/repository"
	"auth-service/dto"
	"net/http"
	"shared/utils"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	userRepository repository.UserRepository
}

func NewAuthHandler(userRepository repository.UserRepository) *AuthHandler {
	return &AuthHandler{userRepository}
}

// Login godoc
// @Summary Login user
// @Description Authenticate user and provide a token
// @Tags auth
// @Accept json
// @Produce json
// @Param login body dto.LoginRequest true "User login request"
// @Success 200 {object} dto.LoginResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /login [post]
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

// Register godoc
// @Summary Register a new user
// @Description Create a new user
// @Tags auth
// @Accept json
// @Produce json
// @Param register body dto.RegisterRequest true "User register request"
// @Success 200 {object} dto.RegisterResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /register [post]
func (h *AuthHandler) Register(c *gin.Context) {
	var req dto.RegisterRequest

	// Bind JSON to RegisterRequest struct
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
		return
	}

	// Hash password
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "failed to hash password"})
		return
	}

	// Create user model
	user := &model.User{
		Username: req.Username,
		Password: hashedPassword,
		Email:    req.Email,
	}

	// Save user
	if err := h.userRepository.Save(user); err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "failed to register user"})
		return
	}

	// Respond with the registered user details
	c.JSON(http.StatusOK, dto.RegisterResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	})
}
