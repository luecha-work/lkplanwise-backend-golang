package models

import (
	"time"

	"github.com/google/uuid"
)

type CreateAccountRequest struct {
	FirstName   string    `json:"first_name" binding:"required,alphanum"`
	LastName    string    `json:"last_name" binding:"required,alphanum"`
	UserName    string    `json:"username" binding:"required,alphanum"`
	Email       string    `json:"email" binding:"required,email"`
	Password    string    `json:"password" binding:"required"`
	DateOfBirth time.Time `json:"date_of_birth" binding:"required"`
	RoleId      uuid.UUID `json:"role_id" binding:"required"`
}

type AccountResponse struct {
	UserName  string    `json:"username"`
	FullName  string    `json:"full_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	CreatedBy string    `json:"created_by"`
}
