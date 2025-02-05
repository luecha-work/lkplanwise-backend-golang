package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/lkplanwise-api/constant"
	db "github.com/lkplanwise-api/db/sqlc"
	"github.com/lkplanwise-api/services"
	"github.com/lkplanwise-api/token"
)

const (
	authorizationHeaderKey  = "authorization"
	authorizationTypeBearer = "bearer"
	authorizationPayloadKey = "authorization_payload"
)

func AuthMiddleware(tokenMaker token.Maker, store db.Store) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeaderKey := ctx.GetHeader(authorizationHeaderKey)

		if len(authorizationHeaderKey) == 0 {
			err := errors.New("authorization header is not provider")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, constant.ErrorResponse(err))
			return
		}

		fields := strings.Fields(authorizationHeaderKey)
		if len(fields) < 2 {
			err := errors.New("invalid authorization header format")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, constant.ErrorResponse(err))
			return
		}

		authorizationType := strings.ToLower(fields[0])
		if authorizationType != authorizationTypeBearer {
			err := fmt.Errorf("unsupported authorization type %s", authorizationType)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, constant.ErrorResponse(err))
			return
		}

		accessToken := fields[1]
		payload, err := tokenMaker.VerifyToken(accessToken)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, constant.ErrorResponse(err))
			return
		}

		//TODO: Check session for authorization
		if _, err = services.CheckLKPlanWiseSessionUnavailable(ctx, store, payload); err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, constant.ErrorResponse(err))
			return
		}

		ctx.Set(authorizationPayloadKey, payload)
		ctx.Next()
	}
}
