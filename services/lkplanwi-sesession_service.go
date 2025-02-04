package services

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/lkplanwise-api/db/sqlc"
	"github.com/lkplanwise-api/models"
	"github.com/lkplanwise-api/token"
)

func CheckSession(ctx *gin.Context, store db.Store, account db.Account) (db.LKPlanWiseSession, error) {
	session, err := store.GetLKPlanWiseSessionForLogin(ctx, db.GetLKPlanWiseSessionForLoginParams{
		AccountId: pgtype.UUID{Bytes: account.Id, Valid: true},
		LoginIp:   ctx.ClientIP(),
	})

	if err != nil {
		return db.LKPlanWiseSession{}, err
	}

	return session, nil
}

func CreateLKPlanWiseSession(
	ctx *gin.Context,
	store db.Store,
	tokenMaker token.Maker,
	account db.Account,
	req models.LoginRequest,
	accessPayload *token.Payload,
	refreshPayload *token.Payload,
	accessToken string) (db.LKPlanWiseSession, error) {
	newSession, err := store.CreateLKPlanWiseSession(ctx, db.CreateLKPlanWiseSessionParams{
		AccountId:      pgtype.UUID{Bytes: account.Id, Valid: true},
		LoginAt:        pgtype.Timestamptz{Time: accessPayload.IssuedAt, Valid: true},
		Platform:       pgtype.Text{String: req.Platform, Valid: true},
		Os:             pgtype.Text{String: req.Os, Valid: true},
		Browser:        pgtype.Text{String: req.Browser, Valid: true},
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
		return db.LKPlanWiseSession{}, err
	}

	return newSession, nil
}

func DeleteLKPlanWiseSession(
	ctx *gin.Context,
	store db.Store,
	sessionId uuid.UUID) error {

	_, err := store.DeleteLKPlanWiseSession(ctx, sessionId)

	if err != nil {
		return err
	}

	return nil
}
