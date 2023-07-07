package test

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/riad/simple_auth/src/db/sqlc"
	HP "github.com/riad/simple_auth/src/http"
	"github.com/riad/simple_auth/src/util"
	"github.com/stretchr/testify/require"
)

func NewTestServer(t *testing.T, store db.Store) *(HP.Server) {
	config, err := util.LoadConfig("../../../../.")
	if err != nil {
		fmt.Printf("cannot load config")
	}
	config = util.Config{
		Token: util.TokenConfig{
			SymetricKey:         util.RandomString(32),
			AccessTokenDuration: time.Minute,
		},
	}

	server, err := HP.NewServer(config, store)
	require.NoError(t, err)

	return server
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)

	os.Exit(m.Run())
}
