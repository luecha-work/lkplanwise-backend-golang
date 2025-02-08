package services

import (
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/lkplanwise-api/constant"
	db "github.com/lkplanwise-api/db/sqlc"
	"github.com/lkplanwise-api/models"
	"github.com/lkplanwise-api/token"
)

func CheckLKPlanWiseSessionForLogin(ctx *gin.Context, store db.Store, account db.Account) (db.LKPlanWiseSession, error) {
	session, err := store.GetLKPlanWiseSessionForLogin(ctx, db.GetLKPlanWiseSessionForLoginParams{
		AccountId: pgtype.UUID{Bytes: account.Id, Valid: true},
		LoginIp:   ctx.ClientIP(),
	})
	if err != nil {
		return db.LKPlanWiseSession{}, err
	}

	return session, nil
}

func CheckLKPlanWiseSessionUnavailable(ctx *gin.Context, store db.Store, accessPayload *token.Payload) (bool, error) {
	session, err := store.GetLKPlanWiseSessionForLogin(ctx, db.GetLKPlanWiseSessionForLoginParams{
		AccountId: pgtype.UUID{Bytes: accessPayload.AccountId, Valid: true},
		LoginIp:   ctx.ClientIP(),
	})
	if err != nil {
		return true, err
	}

	//TODO: If AccountId and LoginIp do not match then it is an attack. Blocking the session
	if session.AccountId.Bytes != accessPayload.ID && session.LoginIp != ctx.ClientIP() {
		//TODO: Block session
		_, err = store.UpdateLKPlanWiseSession(ctx, db.UpdateLKPlanWiseSessionParams{
			ID:            session.Id,
			Sessionstatus: pgtype.Text{String: constant.SessionBlocked, Valid: true},
			Updatedat:     pgtype.Timestamptz{Time: time.Now(), Valid: true},
			Updatedby:     pgtype.Text{String: "system", Valid: true},
		})
		if err != nil {
			return true, err
		}

		return true, errors.New("session is blocked. Please log in and log in again")
	}

	//TODO: If session is not active, the system will not be able to be used. You will need to log in again.
	if session.SessionStatus != constant.SessionActive {
		return true, errors.New("the session has a problem. Please log in and log in again")
	}

	//TODO: If the session expires, the system will not be able to be used. You will need to log in again.
	if session.ExpirationTime.Valid && session.ExpirationTime.Time.After(time.Now()) {
		_, err = store.UpdateLKPlanWiseSession(ctx, db.UpdateLKPlanWiseSessionParams{
			ID:            session.Id,
			Sessionstatus: pgtype.Text{String: constant.SessionExpired, Valid: true},
			Updatedat:     pgtype.Timestamptz{Time: time.Now(), Valid: true},
			Updatedby:     pgtype.Text{String: "system", Valid: true},
		})
		if err != nil {
			return true, err
		}

		return true, errors.New("the session has expired. Please log in and log in again")
	}

	return false, nil
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
		ExpirationTime: pgtype.Timestamptz{Time: refreshPayload.ExpiredAt, Valid: true},
		SessionStatus:  constant.SessionActive,
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
