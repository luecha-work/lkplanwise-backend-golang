package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	db "github.com/lkplanwise-api/db/sqlc"
	"github.com/lkplanwise-api/token"
	"github.com/lkplanwise-api/utils"
)

// Server serves HTTP requests for our banking service.
type Server struct {
	config     utils.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

// NewServer creates a new HTTP server.
func NewServer(config utils.Config, store db.Store) (*Server, error) {
	//TODO: Select PasetoMaker or JETMaker to generate token
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	// tokenMaker, err := token.NewJWTMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("Connot create token maker: %w", err)
	}

	server := &Server{config: config, store: store, tokenMaker: tokenMaker}

	server.setupRouter()

	return server, nil
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
