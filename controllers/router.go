package controllers

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/lkplanwise-api/middleware"
)

// setupRouter sets up the routes for the server.
func (server *Server) setupRouter() {

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		// AllowOrigins:     server.config.AllowedOrigins,
		AllowOrigins:     []string{"http://example.com", "http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// New route for hello world
	router.GET("/api/hello", func(c *gin.Context) {
		c.String(200, "Hello, World!")
	})

	router.POST("/api/auth/login", server.login)
	router.POST("/api/auth/refresh-token", server.refreshAccessToken)
	router.POST("/api/register", server.resister)
	router.GET("/api/verify_email", server.verifyEmail)

	// TODO: Use Middleware for routes
	authRoutes := router.Group("/").Use(middleware.AuthMiddleware(server.tokenMaker, server.store))

	authRoutes.POST("/accounts", server.createAccount)
	// authRoutes.GET("/accounts/:id", server.getAccount)
	// authRoutes.GET("/accounts", server.listAccounts)
	// authRoutes.POST("/transfers", server.createTransfer)

	server.router = router
}
