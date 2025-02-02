package controllers

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/lkplanwise-api/db/sqlc"
	"github.com/lkplanwise-api/token"
	"github.com/lkplanwise-api/utils"
	"github.com/rs/zerolog/log"
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
	//TODO: Select PasetoMaker or JWTMaker to generate token
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	// tokenMaker, err := token.NewJWTMaker(config.TokenSymmetricKey)
	if err != nil {
		log.Error().Err(err).Msg("Cannot create token maker")
		return nil, err
	}

	server := &Server{config: config, store: store, tokenMaker: tokenMaker}

	server.setupRouter()

	return server, nil
}

// func (server *Server) Start(address string) error {
// 	return server.router.Run(address)
// }

// Start creates an HTTP server and uses ListenAndServe to start it.
func (server *Server) Start(address string) error {
	// Create an HTTP server using the router and address
	srv := &http.Server{
		Addr:    address,
		Handler: server.router, // set the Gin router as the HTTP handler
	}

	// Create a channel to listen for OS signals (e.g., termination signals)
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGINT, syscall.SIGTERM)

	// Start the HTTP server in a goroutine
	go func() {
		log.Info().Msgf("Starting Gin server at %s", address)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal().Err(err).Msg("Server failed")
		}
	}()

	// Wait for a termination signal (SIGINT or SIGTERM)
	<-stopChan
	log.Info().Msg("Shutting down Gin server gracefully...")

	// Create a context with a timeout to ensure a graceful shutdown
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Attempt to gracefully shutdown the server
	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Error().Err(err).Msg("Error shutting down server")
		return err
	}

	log.Info().Msg("Server gracefully stopped")
	return nil
}
