package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lkplanwise-api/constant"
	db "github.com/lkplanwise-api/db/sqlc"
	models "github.com/lkplanwise-api/models"
	"github.com/lkplanwise-api/services"
)

func (server *Server) login(ctx *gin.Context) {
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

func (server *Server) resister(ctx *gin.Context) {
	var req models.CreateAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, constant.ErrorResponse(err))
		return
	}

	account, err := services.CreateAccount(ctx, server.store, req)
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
