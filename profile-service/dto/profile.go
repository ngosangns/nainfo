package dto

type UpdateProfileRequest struct {
	Username    string
	Name        string `json:"name"`
	Description string `json:"description"`
	Email       string `json:"email" binding:"email"`
	Address     string `json:"address"`
	Facebook    string `json:"facebook"`
	LinkedIn    string `json:"linkedin"`
	GitHub      string `json:"github"`
}

type ProfileResponse struct {
	Username    string `json:"username"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Email       string `json:"email"`
	Address     string `json:"address"`
	Facebook    string `json:"facebook"`
	LinkedIn    string `json:"linkedin"`
	GitHub      string `json:"github"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type MessageResponse struct {
	Message string `json:"message"`
}
