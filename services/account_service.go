package services

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/lkplanwise-api/db/sqlc"
	"github.com/lkplanwise-api/models"
	"github.com/lkplanwise-api/utils"
)

func newAccountResponse(account db.Account) models.AccountResponse {
	return models.AccountResponse{
		UserName:  account.UserName.String,
		FullName:  account.FirstName.String + " " + account.LastName.String,
		Email:     account.Email.String,
		CreatedAt: account.CreatedAt.Time,
		CreatedBy: account.CreatedBy.String,
	}
}

// Define a method on the server.Server type using pointer receiver
func CreateAccount(store db.Store, ctx *gin.Context) (models.AccountResponse, error) {
	var req models.CreateAccountRequest

	// Hash รหัสผ่าน
	hashPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return models.AccountResponse{}, err
	}

	// สร้าง Account
	arg := db.CreateAccountParams{
		UserName: pgtype.Text{
			String: req.UserName,
			Valid:  true,
		},
		FirstName: pgtype.Text{
			String: req.FirstName,
			Valid:  true,
		},
		LastName: pgtype.Text{
			String: req.LastName,
		},
		Email: pgtype.Text{
			String: req.Email,
			Valid:  true,
		},
		PasswordHash: pgtype.Text{
			String: hashPassword,
		},
		DateOfBirth: pgtype.Date{
			Time: req.DateOfBirth,
		},
		RoleId: uuid.UUID(req.RoleId),
		CreatedAt: pgtype.Timestamptz{
			Time:  time.Now(),
			Valid: true,
		},
		CreatedBy: pgtype.Text{
			String: "system",
		},
		UpdatedAt: pgtype.Timestamptz{
			Time:  time.Time{},
			Valid: false,
		},
		UpdatedBy: pgtype.Text{Valid: false},
	}

	account, err := store.CreateAccount(ctx, arg)
	if err != nil {
		return models.AccountResponse{}, err
	}

	// เตรียมข้อมูลที่จะส่งกลับ
	userResponse := newAccountResponse(account)

	return userResponse, nil
}
