package controllers

import (
	"github.com/gin-gonic/gin"
)

// setupRouter sets up the routes for the server.
func (server *Server) setupRouter() {
	router := gin.Default()

	// Register routes
	router.GET("/hello", func(c *gin.Context) {
		c.String(200, "hello")
	})
	router.POST("/accounts", server.CreateAccount)
	// router.POST("/users", server.createUser)
	// router.POST("/users/login", server.loginUser)
	// router.POST("/token/refresh-token", server.renewAccessToken)

	// TODO: Use Middleware for routes
	// authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))

	// authRoutes.POST("/accounts", server.createAccount)
	// authRoutes.GET("/accounts/:id", server.getAccount)
	// authRoutes.GET("/accounts", server.listAccounts)

	// authRoutes.POST("/transfers", server.createTransfer)

	server.router = router
}
