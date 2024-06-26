package gapi

import (
	"os"
	"testing"

	db "github.com/Streamfair/streamfair_token_svc/db/sqlc"
	"github.com/Streamfair/streamfair_token_svc/util"
	"github.com/stretchr/testify/require"
)

func newTestServer(t *testing.T, store db.Store) *Server {
	config, err := util.LoadConfig()
	require.NoError(t, err)

	server, err := NewServer(config, store)
	require.NoError(t, err)

	return server
}

func TestMain(m *testing.M) {
	os.Chdir("../")
	os.Exit(m.Run())
}
