package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lkplanwise-api/constant"
	models "github.com/lkplanwise-api/models"
	"github.com/lkplanwise-api/services"
)

func (server *Server) CreateAccount(ctx *gin.Context) {
	var req models.CreateAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, constant.ErrorResponse(err))
		return
	}

	account, err := services.CreateAccount(ctx, server.store, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, constant.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}
