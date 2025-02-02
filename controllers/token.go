package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lkplanwise-api/constant"
	db "github.com/lkplanwise-api/db/sqlc"
	"github.com/lkplanwise-api/models"
	"github.com/lkplanwise-api/utils"
)

func (server *Server) refreshTokenAccessToken(ctx *gin.Context) {
	var req models.NewAccessTokenRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, constant.ErrorResponse(err))
		return
	}

	refreshPayload, err := server.tokenMaker.VerifyToken(req.RefreshToken)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, constant.ErrorResponse(err))
		return
	}

	// session, err := server.store.GetLKPlanWiseSessionById(ctx, refreshPayload.ID)
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, constant.ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, constant.ErrorResponse(err))
		return
	}

	// if session.IsBlocked {
	// 	err := fmt.Errorf("blocked session")
	// 	ctx.JSON(http.StatusUnauthorized, constant.ErrorResponse(err))
	// 	return
	// }

	// if session.Username != refreshPayload.Username {
	// 	err := fmt.Errorf("incorrect session user")
	// 	ctx.JSON(http.StatusUnauthorized, constant.ErrorResponse(err))
	// 	return
	// }

	// if session.RefreshToken != req.RefreshToken {
	// 	err := fmt.Errorf("mismathched session token")
	// 	ctx.JSON(http.StatusUnauthorized, constant.ErrorResponse(err))
	// 	return
	// }

	// if time.Now().After(session.ExpiresAt) {
	// 	err := fmt.Errorf("expired session")
	// 	ctx.JSON(http.StatusUnauthorized, constant.ErrorResponse(err))
	// 	return
	// }

	accessToken, accessPayload, err := server.tokenMaker.CreateToken(
		refreshPayload.Username,
		utils.DepositorRole,
		server.config.AccessTokenDuration,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, constant.ErrorResponse(err))
		return
	}

	rsp := models.NewAccessTokenResponse{
		AccessToken:          accessToken,
		AccessTokenExpiresAt: accessPayload.ExpiredAt,
	}
	ctx.JSON(http.StatusOK, rsp)
}
