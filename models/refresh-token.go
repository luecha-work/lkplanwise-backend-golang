package models

type RefreshAccessTokenResponse struct {
	AccessToken string `json:"access_token"`
}

type RefreshAccessTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}
