package controllers

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/lkplanwise-api/middleware"
)

// setupRouter sets up the routes for the server.
func (server *Server) setupRouter() {
	core := cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{
			http.MethodHead, http.MethodOptions, http.MethodGet,
			http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete,
		},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	router := gin.Default()

	router.Use(core)

	router.POST("/api/register", server.resister)
	router.POST("/api/auth/login", server.login)
	// router.POST("/token/refresh-token", server.renewAccessToken)
	router.GET("/api/verify_email", server.verifyEmail)

	// TODO: Use Middleware for routes
	authRoutes := router.Group("/").Use(middleware.AuthMiddleware(server.tokenMaker, server.store))

	authRoutes.POST("/accounts", server.createAccount)
	// authRoutes.GET("/accounts/:id", server.getAccount)
	// authRoutes.GET("/accounts", server.listAccounts)
	// authRoutes.POST("/transfers", server.createTransfer)

	server.router = router
}
