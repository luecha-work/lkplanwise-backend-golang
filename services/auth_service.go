package services

import (
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/lkplanwise-api/db/sqlc"
	"github.com/lkplanwise-api/models"
	"github.com/lkplanwise-api/token"
	"github.com/lkplanwise-api/utils"
)

func Login(ctx *gin.Context, store db.Store, req models.LoginRequest, tokenMaker token.Maker, config utils.Config) (models.LoginResponse, error) {

	account, err := store.GetAccountByUsername(ctx, req.Username)
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			// ctx.JSON(http.StatusNotFound, constant.ErrorResponse(err))
			return models.LoginResponse{}, err
		}
		// ctx.JSON(http.StatusInternalServerError, constant.ErrorResponse(err))
		return models.LoginResponse{}, err
	}

	err = utils.CheckPassword(req.Password, account.PasswordHash.String)
	if err != nil {
		// ctx.JSON(http.StatusUnauthorized, constant.ErrorResponse(err))
		return models.LoginResponse{}, err
	}

	accessToken, accessPayload, err := tokenMaker.CreateToken(
		account.UserName,
		utils.DepositorRole,
		config.AccessTokenDuration,
	)

	if err != nil {
		// ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return models.LoginResponse{}, err
	}

	refreshToken, refreshPayload, err := tokenMaker.CreateToken(
		account.UserName,
		utils.DepositorRole,
		config.RefreshTokenDuration,
	)

	if err != nil {
		// ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return models.LoginResponse{}, err
	}

	session, err := store.CreateLKPlanWiseSession(ctx, db.CreateLKPlanWiseSessionParams{
		AccountId:      pgtype.UUID{Bytes: account.Id, Valid: true},
		LoginAt:        pgtype.Timestamptz{Time: accessPayload.IssuedAt, Valid: true},
		Platform:       pgtype.Text{String: "web", Valid: true},
		Os:             pgtype.Text{String: "windows", Valid: true},
		Browser:        pgtype.Text{String: "chrome", Valid: true},
		LoginIp:        ctx.ClientIP(),
		IssuedTime:     pgtype.Timestamptz{Time: accessPayload.IssuedAt, Valid: true},
		ExpirationTime: pgtype.Timestamptz{Time: accessPayload.ExpiredAt, Valid: true},
		SessionStatus:  "A",
		Token:          pgtype.Text{String: accessToken, Valid: true},
		RefreshTokenAt: pgtype.Timestamptz{Time: refreshPayload.IssuedAt, Valid: true},
		CreatedAt:      pgtype.Timestamptz{Time: time.Now(), Valid: true},
		CreatedBy:      pgtype.Text{String: "system", Valid: true},
	})
	if err != nil {
		// ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return models.LoginResponse{}, err
	}

	rsp := models.LoginResponse{
		SessionID:    session.Id,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return rsp, nil
}
