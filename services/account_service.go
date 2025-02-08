package services

import (
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/hibiken/asynq"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/lkplanwise-api/constant"
	db "github.com/lkplanwise-api/db/sqlc"
	"github.com/lkplanwise-api/models"
	"github.com/lkplanwise-api/utils"
	"github.com/lkplanwise-api/worker"
)

// Define a method on the server.Server type using pointer receiver
func CreateAccount(ctx *gin.Context, store db.Store, taskDistributor worker.TaskDistributor, req models.CreateAccountRequest) (models.AccountResponse, error) {
	hashPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return models.AccountResponse{}, err
	}

	if _, err = store.GetRoleById(ctx, uuid.MustParse(req.RoleId)); err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			return models.AccountResponse{}, errors.New("role not found in the system")
		}
		return models.AccountResponse{}, err
	}

	// สร้าง Account
	arg := db.CreateAccountParams{
		Id:           uuid.New(),
		UserName:     req.UserName,
		FirstName:    pgtype.Text{String: req.FirstName, Valid: true},
		LastName:     pgtype.Text{String: req.LastName, Valid: true},
		Email:        req.Email,
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

	userResponse := constant.NewAccountResponse(account)

	//TODO: send email then create account
	taskPayload := &worker.PayloadSendVerifyEmail{
		Email: account.Email,
	}
	opts := []asynq.Option{
		asynq.MaxRetry(10),
		asynq.ProcessIn(10 * time.Second),
		asynq.Queue(worker.QueueCritical),
	}
	taskDistributor.DistributeTaskSendVerifyEmail(ctx, taskPayload, opts...)

	return userResponse, nil
}

func lockedAccount(ctx *gin.Context, store db.Store, email string) (db.Account, error) {
	account, err := store.GetAccountByEmail(ctx, email)
	if err != nil {
		return db.Account{}, err
	}

	updatedAccount, err := store.UpdateAccount(ctx, db.UpdateAccountParams{
		ID:        account.Id,
		Islocked:  pgtype.Bool{Bool: true, Valid: true},
		Updatedat: pgtype.Timestamptz{Time: time.Now(), Valid: true},
		Updatedby: pgtype.Text{String: "system", Valid: true},
	})
	if err != nil {
		return db.Account{}, err
	}

	return updatedAccount, nil
}

func unLockAccount(ctx *gin.Context, store db.Store, email string) (db.Account, error) {
	account, err := store.GetAccountByEmail(ctx, email)
	if err != nil {
		return db.Account{}, err
	}

	updatedAccount, err := store.UpdateAccount(ctx, db.UpdateAccountParams{
		ID:        account.Id,
		Islocked:  pgtype.Bool{Bool: false, Valid: true},
		Updatedat: pgtype.Timestamptz{Time: time.Now(), Valid: true},
		Updatedby: pgtype.Text{String: "system", Valid: true},
	})
	if err != nil {
		return db.Account{}, err
	}

	return updatedAccount, nil
}
