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

func GetAccountById(ctx *gin.Context, store db.Store, req models.GetAccountByIdRequest) (models.AccountResponse, error) {
	accountId, err := uuid.Parse(req.Id)
	if err != nil {
		return models.AccountResponse{}, err
	}
	account, err := store.GetAccountById(ctx, accountId)
	if err != nil {
		return models.AccountResponse{}, err
	}

	return constant.NewAccountResponse(account), nil
}

func GetListAccounts(ctx *gin.Context, store db.Store) ([]models.AccountResponse, error) {
	accounts, err := store.GetAllAccounts(ctx)
	if err != nil {
		ctx.JSON(500, constant.ErrorResponse(err))
		return []models.AccountResponse{}, err
	}

	var accountResponses []models.AccountResponse
	for _, account := range accounts {
		accountResponses = append(accountResponses, constant.NewAccountResponse(account))
	}

	return accountResponses, nil
}

func PagedAccounts(ctx *gin.Context, store db.Store, req models.PagedAccountRequest) ([]models.AccountResponse, error) {
	accounts, err := store.PagedAccounts(ctx, db.PagedAccountsParams{
		Limit:  req.PageSize,
		Offset: (req.PageNumber - 1) * req.PageSize,
	})
	if err != nil {
		return []models.AccountResponse{}, err
	}

	var accountResponses []models.AccountResponse
	for _, account := range accounts {
		accountResponses = append(accountResponses, constant.NewAccountResponse(account))
	}

	return accountResponses, nil
}

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

	accountResponse := constant.NewAccountResponse(account)

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

	return accountResponse, nil
}

func DeleteAccount(ctx *gin.Context, store db.Store, req models.DeleteAccountRequest) error {
	accountId, err := uuid.Parse(req.Id)
	if err != nil {
		return err
	}

	err = store.DeleteAccount(ctx, accountId)
	if err != nil {
		return err
	}

	return nil
}

func UpdateAccount(ctx *gin.Context, store db.Store, accountId uuid.UUID, req models.UpdateAccountRequest) (models.AccountResponse, error) {
	if accountId == uuid.Nil {
		return models.AccountResponse{}, errors.New("invalid account ID")
	}

	account, err := store.GetAccountById(ctx, accountId)
	if err != nil {
		return models.AccountResponse{}, err
	}

	updateParams := db.UpdateAccountParams{
		ID:        account.Id,
		Updatedat: pgtype.Timestamptz{Time: time.Now(), Valid: true},
		Updatedby: pgtype.Text{String: "system", Valid: true},
	}

	if req.FirstName != "" {
		updateParams.Firstname = pgtype.Text{String: req.FirstName, Valid: true}
	}
	if req.LastName != "" {
		updateParams.Lastname = pgtype.Text{String: req.LastName, Valid: true}
	}
	if req.Email != "" {
		updateParams.Email = pgtype.Text{String: req.Email, Valid: true}
	}

	updatedAccount, err := store.UpdateAccount(ctx, updateParams)
	if err != nil {
		return models.AccountResponse{}, err
	}

	return constant.NewAccountResponse(updatedAccount), nil
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
