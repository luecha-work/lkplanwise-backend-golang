package api

// import (
// 	"fmt"

// 	"github.com/gin-gonic/gin"
// 	"github.com/lkplanwise-api/utils"
// )

// // Server serves HTTP requests for our banking service.
// type Server struct {
// 	config     utils.Config
// 	store      db.Store
// 	tokenMaker token.Maker
// 	router     *gin.Engine
// }

// // NewServer creates a new HTTO server.
// func NewServer(config utils.Config, store db.Store) (*Server, error) {
// 	//TODO: Select PasetoMaker or JETMaker to generate token
// 	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
// 	// tokenMaker, err := token.NewJWTMaker(config.TokenSymmetricKey)
// 	if err != nil {
// 		return nil, fmt.Errorf("Connot create token maker: %w", err)
// 	}

// 	server := &Server{config: config, store: store, tokenMaker: tokenMaker}

// 	//TODO: Add Register Validation name is currency
// 	// if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
// 	// 	v.RegisterValidation("currency", validCurrency)
// 	// }

// 	server.setupRouter()

// 	return server, nil
// }

// func (server *Server) setupRouter() {
// 	router := gin.Default()

// 	// router.POST("/users", server.createUser)
// 	// router.POST("/users/login", server.loginUser)
// 	// router.POST("/token/refresh-token", server.renewAccessToken)

// 	// //TODO: Use Middleware for routes
// 	// authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))

// 	// authRoutes.POST("/accounts", server.createAccount)
// 	// authRoutes.GET("/accounts/:id", server.getAccount)
// 	// authRoutes.GET("/accounts", server.listAccounts)

// 	// authRoutes.POST("/transfers", server.createTransfer)

// 	server.router = router
// }

// func (server *Server) Start(address string) error {
// 	return server.router.Run(address)
// }

// func errorResponse(err error) gin.H {
// 	return gin.H{
// 		"error": err.Error(),
// 	}
// }
