package services

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	db "github.com/lkplanwise-api/db/sqlc"
	"github.com/lkplanwise-api/models"
	"github.com/lkplanwise-api/token"
	"github.com/lkplanwise-api/utils"
)

func Login(ctx *gin.Context, store db.Store, req models.LoginRequest, tokenMaker token.Maker, config utils.Config) (models.LoginResponse, error) {
	account, err := store.GetAccountByEmail(ctx, req.Email)
	if err != nil {
		fmt.Printf("get account error: %s\n", err)
		if errors.Is(err, db.ErrRecordNotFound) {
			//TODO: If account not found, then check to create for brute force
			ManageBlockBruteForce(ctx, store, req.Email)
			return models.LoginResponse{}, errors.New("username or password is incorrect")
		}
	}

	err = utils.CheckPassword(req.Password, account.PasswordHash.String)
	if err != nil {
		fmt.Printf("check password error: %s\n", err)
		//TODO: If password not match, then check to create for brute force
		ManageBlockBruteForce(ctx, store, req.Email)
		return models.LoginResponse{}, errors.New("username or password is incorrect")
	}

	role, err := store.GetRoleById(ctx, account.RoleId)
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			return models.LoginResponse{}, errors.New("this account has no role found")
		}
		return models.LoginResponse{}, err
	}

	accessToken, accessPayload, err := tokenMaker.CreateToken(
		account.Id,
		account.Email,
		role.RoleName.String,
		config.AccessTokenDuration,
	)

	if err != nil {
		return models.LoginResponse{}, err
	}

	refreshToken, refreshPayload, err := tokenMaker.CreateToken(
		account.Id,
		account.Email,
		role.RoleName.String,
		config.RefreshTokenDuration,
	)
	if err != nil {
		return models.LoginResponse{}, err
	}

	//TODO: Check session and create new session
	var sessionId uuid.UUID

	if session, err := CheckLKPlanWiseSessionForLogin(ctx, store, account); err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			//TODO: If session not found, then create new session
			newSession, err := CreateLKPlanWiseSession(ctx, store, tokenMaker, account, req, accessPayload, refreshPayload, accessToken)

			if err != nil {
				return models.LoginResponse{}, err
			}

			sessionId = newSession.Id
		}
	} else {
		//TODO: If session found, then delete old session and create new session
		DeleteLKPlanWiseSession(ctx, store, session.Id)

		newSession, err := CreateLKPlanWiseSession(ctx, store, tokenMaker, account, req, accessPayload, refreshPayload, accessToken)
		if err != nil {
			return models.LoginResponse{}, err
		}

		sessionId = newSession.Id
	}

	//TODO: Check for unlock brute force
	if _, err = checkForUnLockBruteForce(ctx, store, req.Email); err != nil {
		return models.LoginResponse{}, err
	}

	rsp := models.LoginResponse{
		SessionID:    sessionId,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return rsp, nil
}
