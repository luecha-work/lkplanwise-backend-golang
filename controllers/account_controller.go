package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/lkplanwise-api/db/sqlc"
	models "github.com/lkplanwise-api/models"
	"github.com/lkplanwise-api/services"
)

func newAccountResponse(account db.Account) models.AccountResponse {
	return models.AccountResponse{
		UserName:  account.UserName.String,
		FullName:  account.FirstName.String + " " + account.LastName.String,
		Email:     account.Email.String,
		CreatedAt: account.CreatedAt.Time,
		CreatedBy: account.CreatedBy.String,
	}
}

func (server *Server) CreateAccount(ctx *gin.Context) {
	var req models.CreateAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	account, err := services.CreateAccount(server.store, ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}
