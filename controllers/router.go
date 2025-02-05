package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/lkplanwise-api/middleware"
)

// setupRouter sets up the routes for the server.
func (server *Server) setupRouter() {
	router := gin.Default()

	// Register routes
	router.GET("/hello", func(c *gin.Context) {
		c.String(200, "hello")
	})

	router.POST("/auth/register", server.resister)
	router.POST("/auth/login", server.login)
	// router.POST("/token/refresh-token", server.renewAccessToken)

	// TODO: Use Middleware for routes
	authRoutes := router.Group("/").Use(middleware.AuthMiddleware(server.tokenMaker, server.store))

	authRoutes.POST("/accounts", server.CreateAccount)
	// authRoutes.GET("/accounts/:id", server.getAccount)
	// authRoutes.GET("/accounts", server.listAccounts)

	// authRoutes.POST("/transfers", server.createTransfer)

	server.router = router
}
