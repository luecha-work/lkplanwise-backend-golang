package controllers

import (
	"os"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/lkplanwise-api/db/sqlc"
	"github.com/lkplanwise-api/utils"
	"github.com/lkplanwise-api/worker"
	"github.com/stretchr/testify/require"
)

func NewTestServer(t *testing.T, store db.Store, taskDistributor worker.TaskDistributor) *Server {
	config := utils.Config{
		TokenSymmetricKey:   utils.RandomString(32),
		AccessTokenDuration: time.Minute,
	}

	server, err := NewServer(config, store, taskDistributor)
	require.NoError(t, err)

	return server
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}
