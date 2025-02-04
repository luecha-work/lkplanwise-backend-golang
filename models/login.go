package models

import (
	"github.com/google/uuid"
)

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required,min=8"`
	Os       string `json:"os" binding:"required"`
	Platform string `json:"platform" binding:"required"`
	Browser  string `json:"browser" binding:"required"`
}

type LoginResponse struct {
	SessionID    uuid.UUID `json:"session_id"`
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
}
