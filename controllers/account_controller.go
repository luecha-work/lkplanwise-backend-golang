package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lkplanwise-api/constant"
	db "github.com/lkplanwise-api/db/sqlc"
	models "github.com/lkplanwise-api/models"
	"github.com/lkplanwise-api/services"
)

func (server *Server) CreateAccount(ctx *gin.Context) {
	var req models.CreateAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, constant.ErrorResponse(err))
		return
	}

	rsp, err := services.CreateAccount(ctx, server.store, req)
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
