package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/lkplanwise-api/middleware"
)

// setupRouter sets up the routes for the server.
func (s *Server) setupRouter() {
	router := gin.Default()

	// Register routes
	router.GET("/hello", func(c *gin.Context) {
		c.String(200, "hello")
	})

	router.POST("/auth/register", s.resister)
	router.POST("/auth/login", s.login)
	// router.POST("/token/refresh-token", server.renewAccessToken)

	// TODO: Use Middleware for routes
	authRoutes := router.Group("/").Use(middleware.AuthMiddleware(s.tokenMaker))

	authRoutes.POST("/accounts", s.CreateAccount)
	// authRoutes.GET("/accounts/:id", server.getAccount)
	// authRoutes.GET("/accounts", server.listAccounts)

	// authRoutes.POST("/transfers", server.createTransfer)

	s.router = router
}
