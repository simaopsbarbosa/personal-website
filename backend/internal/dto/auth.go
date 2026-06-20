package dto

type UserLoginRequest struct {
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Message string `json:"message"`
	Token   string `json:"token,omitempty"`
}
