package models

import "time"

type NewAccessTokenResponse struct {
	AccessToken          string    `json:"access_token"`
	AccessTokenExpiresAt time.Time `json:"access_token_expires_at"`
}

type NewAccessTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}
