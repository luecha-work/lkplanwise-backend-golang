package models

import (
	"time"

	"github.com/google/uuid"
)

type CreateAccountRequest struct {
	FirstName   string `json:"first_name" binding:"required,alphanum"`
	LastName    string `json:"last_name" binding:"required,alphanum"`
	UserName    string `json:"username" binding:"required"`
	Email       string `json:"email" binding:"required,email"`
	Password    string `json:"password" binding:"required,min=8"`
	DateOfBirth string `json:"date_of_birth" binding:"required"`
	RoleId      string `json:"role_id" binding:"required"`
}

type AccountResponse struct {
	Id        uuid.UUID `json:"id"`
	UserName  string    `json:"username"`
	FullName  string    `json:"full_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	CreatedBy string    `json:"created_by"`
}

type UpdateAccountRequest struct {
	FirstName   string `json:"first_name" binding:"required,alphanum"`
	LastName    string `json:"last_name" binding:"required,alphanum"`
	UserName    string `json:"username" binding:"required"`
	Email       string `json:"email" binding:"required,email"`
	Password    string `json:"password" binding:"required,min=8"`
	DateOfBirth string `json:"date_of_birth" binding:"required"`
	RoleId      string `json:"role_id" binding:"required"`
}

type GetAccountByIdRequest struct {
	Id string `uri:"id" binding:"required,uuid"`
}

type DeleteAccountRequest struct {
	Id string `uri:"id" binding:"required,uuid"`
}

type PagedAccountRequest struct {
	PageNumber int32 `form:"page_number" binding:"required,min=1"`
	PageSize   int32 `form:"page_size" binding:"required,min=5,max=10"`
}
