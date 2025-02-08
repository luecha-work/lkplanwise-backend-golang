package services

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/lkplanwise-api/db/sqlc"
	"github.com/lkplanwise-api/models"
	"github.com/lkplanwise-api/val"
	"github.com/rs/zerolog/log"
)

func VerifyEmailHandler(ctx *gin.Context, store db.Store, req models.VerifyEmailRequest) (models.VerifyEmailResponse, error) {
	//TODO: Check mail request
	violations := validateVerifyEmailRequest(&req)
	if violations != nil {
		fmt.Println("Violations : ", violations)
		return models.VerifyEmailResponse{}, errors.New(violations[1])
	}

	//TODO: Verify email by email_id and secret_code
	vfEmail, err := store.UpdateVerifyEmail(ctx, db.UpdateVerifyEmailParams{
		ID:         int64(req.EmailId),
		Secretcode: req.SecretCode,
	})
	if err != nil {
		log.Error().Err(err).Msg("failed to verify email")
		return models.VerifyEmailResponse{}, err
	}

	//TODO: Get account to update IsMailVerified
	account, err := store.GetAccountByUsername(ctx, vfEmail.UserName)
	if err != nil {
		log.Error().Err(err).Msg("failed to get account")
		return models.VerifyEmailResponse{}, err
	}

	//TODO: Update IsMailVerified
	accountUpdated, err := store.UpdateAccount(ctx, db.UpdateAccountParams{
		ID:             account.Id,
		Username:       pgtype.Text{String: vfEmail.UserName, Valid: true},
		Ismailverified: pgtype.Bool{Bool: true, Valid: true},
		Updatedat:      pgtype.Timestamptz{Time: account.UpdatedAt.Time, Valid: true},
		Updatedby:      pgtype.Text{String: "system", Valid: true},
	})
	if err != nil {
		log.Error().Err(err).Msg("failed to update account")
		return models.VerifyEmailResponse{}, err
	}

	rsp := models.VerifyEmailResponse{
		IsVerified: accountUpdated.IsMailVerified,
	}

	return rsp, nil
}

// validateVerifyEmailRequest ตรวจสอบความถูกต้องของข้อมูล
func validateVerifyEmailRequest(req *models.VerifyEmailRequest) (violations []string) {
	if err := val.ValidateEmailId(int64(req.EmailId)); err != nil {
		violations = append(violations, "invalid email_id")
	}

	if err := val.ValidateSecretCode(req.SecretCode); err != nil {
		violations = append(violations, "invalid secret_code")
	}

	return violations
}
