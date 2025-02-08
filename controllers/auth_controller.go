package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/lkplanwise-api/constant"
	db "github.com/lkplanwise-api/db/sqlc"
	models "github.com/lkplanwise-api/models"
	"github.com/lkplanwise-api/services"
	"github.com/lkplanwise-api/utils"
	"github.com/lkplanwise-api/worker"
)

func (server *Server) login(ctx *gin.Context) {
	logger := worker.NewLogger()
	logger.Info("Login request")
	var req models.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, constant.ErrorResponse(err))
		return
	}

	isBlocked, err := services.CheckBlockedBruteForce(ctx, server.store, req.Email)
	if isBlocked {
		ctx.JSON(http.StatusLocked, constant.ErrorResponse(err))
		return
	}

	account, err := services.Login(ctx, server.store, req, server.tokenMaker, server.config)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, constant.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}

func (server *Server) refreshAccessToken(ctx *gin.Context) {
	var req models.RefreshAccessTokenRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, constant.ErrorResponse(err))
		return
	}

	refreshPayload, err := server.tokenMaker.VerifyToken(req.RefreshToken)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, constant.ErrorResponse(err))
		return
	}

	session, err := server.store.GetLKPlanWiseSessionForAuth(ctx, db.GetLKPlanWiseSessionForAuthParams{
		AccountId: pgtype.UUID{Bytes: refreshPayload.AccountId, Valid: true},
		LoginIp:   ctx.ClientIP(),
	})
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, constant.ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, constant.ErrorResponse(err))
		return
	}

	// Check if the session is blocked and return an error if it is
	if session.SessionStatus == constant.SessionBlocked {
		err := fmt.Errorf("blocked session, please login again")
		ctx.JSON(http.StatusUnauthorized, constant.ErrorResponse(err))
		return
	}
	// Check if the session status is expired
	if session.SessionStatus == constant.SessionExpired || (session.ExpirationTime.Valid && time.Now().After(session.ExpirationTime.Time)) {
		err := fmt.Errorf("expired session, please login again")
		ctx.JSON(http.StatusUnauthorized, constant.ErrorResponse(err))
		return
	}

	// Ensure the session account ID matches the account ID in the refresh token payload
	if session.AccountId.Bytes != refreshPayload.AccountId {
		err := fmt.Errorf("incorrect session account")
		ctx.JSON(http.StatusUnauthorized, constant.ErrorResponse(err))
		return
	}

	// Ensure the session refresh token matches the refresh token in the request
	if session.RefreshToken.String != req.RefreshToken {
		err := fmt.Errorf("mismatched session token")
		ctx.JSON(http.StatusUnauthorized, constant.ErrorResponse(err))
		return
	}

	// Create a new access token
	newAccessToken, _, err := server.tokenMaker.CreateToken(
		refreshPayload.AccountId,
		refreshPayload.Email,
		utils.DepositorRole,
		server.config.AccessTokenDuration,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, constant.ErrorResponse(err))
		return
	}

	session, err = server.store.UpdateLKPlanWiseSession(ctx, db.UpdateLKPlanWiseSessionParams{
		ID:             session.Id,
		Refreshtokenat: pgtype.Timestamptz{Time: time.Now(), Valid: true},
		Updatedat:      pgtype.Timestamptz{Time: time.Now(), Valid: true},
		Updatedby:      pgtype.Text{String: "system", Valid: true},
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, constant.ErrorResponse(err))
		return
	}

	rsp := models.RefreshAccessTokenResponse{
		AccessToken: newAccessToken,
	}

	ctx.JSON(http.StatusOK, rsp)
}

func (server *Server) resister(ctx *gin.Context) {
	var req models.CreateAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, constant.ErrorResponse(err))
		return
	}

	account, err := services.CreateAccount(ctx, server.store, server.taskDistributor, req)
	if err != nil {
		errCode := db.ErrorCode(err)
		if errCode == db.ForeignKeyViolation || errCode == db.UniqueViolation {
			ctx.JSON(http.StatusForbidden, constant.ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, constant.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}
