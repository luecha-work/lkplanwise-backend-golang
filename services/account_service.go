package services

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/lkplanwise-api/constant"
	db "github.com/lkplanwise-api/db/sqlc"
	"github.com/lkplanwise-api/models"
	"github.com/lkplanwise-api/utils"
)

// Define a method on the server.Server type using pointer receiver
func CreateAccount(ctx *gin.Context, store db.Store, req models.CreateAccountRequest) (models.AccountResponse, error) {

	// Hash รหัสผ่าน
	hashPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return models.AccountResponse{}, err
	}

	// สร้าง Account
	arg := db.CreateAccountParams{
		Id:           uuid.New(),
		UserName:     req.UserName,
		FirstName:    pgtype.Text{String: req.FirstName, Valid: true},
		LastName:     pgtype.Text{String: req.LastName, Valid: true},
		Email:        pgtype.Text{String: req.Email, Valid: true},
		PasswordHash: pgtype.Text{String: hashPassword, Valid: true},
		DateOfBirth:  pgtype.Text{String: req.DateOfBirth, Valid: true},
		RoleId:       uuid.MustParse(req.RoleId),
		CreatedAt:    pgtype.Timestamptz{Time: time.Now(), Valid: true},
		CreatedBy:    pgtype.Text{String: "system", Valid: true},
	}

	account, err := store.CreateAccount(ctx, arg)
	if err != nil {
		return models.AccountResponse{}, err
	}

	// เตรียมข้อมูลที่จะส่งกลับ
	userResponse := constant.NewAccountResponse(account)

	return userResponse, nil
}
