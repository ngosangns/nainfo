package dto

// LoginRequest represents the request payload for login.
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// RegisterRequest represents the request payload for registration.
type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

// LoginResponse represents the response payload for login.
type LoginResponse struct {
	Token string `json:"token"`
}

// RegisterResponse represents the response payload for registration.
type RegisterResponse struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

// ErrorResponse represents a common error response.
type ErrorResponse struct {
	Error string `json:"error"`
}
