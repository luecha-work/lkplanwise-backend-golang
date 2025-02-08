package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lkplanwise-api/constant"
	models "github.com/lkplanwise-api/models"
	"github.com/lkplanwise-api/services"
)

func (server *Server) verifyEmail(ctx *gin.Context) {
	fmt.Println("Verify Email Handler")
	var req models.VerifyEmailRequest
	//TODO: Validate request for Query
	if err := ctx.ShouldBindQuery(&req); err != nil {

		ctx.JSON(http.StatusInternalServerError, constant.ErrorResponse(err))
		return
	}

	//TODO: Verify email
	vfEmail, err := services.VerifyEmailHandler(ctx, server.store, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, constant.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, vfEmail)
}
