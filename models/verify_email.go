package models

// type VerifyEmailRequest struct {
// 	EmailId    int64  `json:"email_id,omitempty"`
// 	SecretCode string `json:"secret_code,omitempty"`
// }

// type VerifyEmailResponse struct {
// 	IsVerified bool `json:"is_verified,omitempty"`
// }

type VerifyEmailRequest struct {
	EmailId    int    `form:"email_id" binding:"required"`
	SecretCode string `form:"secret_code" binding:"required"`
}

type VerifyEmailResponse struct {
	IsVerified bool `json:"is_verified" binding:"required"`
}
