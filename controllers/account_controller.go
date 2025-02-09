package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lkplanwise-api/constant"
	db "github.com/lkplanwise-api/db/sqlc"
	models "github.com/lkplanwise-api/models"
	"github.com/lkplanwise-api/services"
	"github.com/rs/zerolog/log"
)

func (server *Server) listAccounts(ctx *gin.Context) {
	accounts, err := services.GetListAccounts(ctx, server.store)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, constant.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, accounts)
}

func (server *Server) getAccountById(ctx *gin.Context) {
	var req models.GetAccountByIdRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		log.Error().Err(err).Msg("error binding uri")
		ctx.JSON(http.StatusBadRequest, constant.ErrorResponse(err))
		return
	}

	account, err := services.GetAccountById(ctx, server.store, req)
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, constant.ErrorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, constant.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}

func (server *Server) pagedAccounts(ctx *gin.Context) {
	var req models.PagedAccountRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		log.Error().Err(err).Msg("error binding query")
		ctx.JSON(http.StatusBadRequest, constant.ErrorResponse(err))
		return
	}

	accounts, err := services.PagedAccounts(ctx, server.store, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, constant.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, accounts)
}

func (server *Server) createAccount(ctx *gin.Context) {
	var req models.CreateAccountRequest
	//TODO: Validate request for json body
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, constant.ErrorResponse(err))
		return
	}

	rsp, err := services.CreateAccount(ctx, server.store, server.taskDistributor, req)
	if err != nil {
		if db.ErrorCode(err) == db.UniqueViolation {
			ctx.JSON(http.StatusForbidden, constant.ErrorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, constant.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, rsp)
}

func (server *Server) deleteAccount(ctx *gin.Context) {
	var req models.DeleteAccountRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		log.Error().Err(err).Msg("error binding uri")
		ctx.JSON(http.StatusBadRequest, constant.ErrorResponse(err))
		return
	}

	err := services.DeleteAccount(ctx, server.store, req)
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, constant.ErrorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, constant.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, nil)
}
