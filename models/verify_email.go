package models

type VerifyEmailRequest struct {
	EmailId    int    `form:"email_id" binding:"required"`
	SecretCode string `form:"secret_code" binding:"required"`
}

type VerifyEmailResponse struct {
	IsVerified bool `json:"is_verified" binding:"required"`
}
