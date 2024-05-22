package dto

type UpdateProfileRequest struct {
	Username string `uri:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

type ProfileResponse struct {
	Username string `uri:"username"`
	Email    string `json:"email"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type MessageResponse struct {
	Message string `json:"message"`
}
