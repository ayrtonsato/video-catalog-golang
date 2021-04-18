package setup

import (
	"database/sql"
	"github.com/ayrtonsato/video-catalog-golang/internal/config"
	"github.com/ayrtonsato/video-catalog-golang/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

type TestSetup struct {
	Config *config.Config
	DB     *sql.DB
	Log    logger.Logger
	Server *gin.Engine
}

func (ts *TestSetup) BuildConfig(t *testing.T, path string) *TestSetup {
	gin.SetMode(gin.TestMode)
	ts.Config = &config.Config{}
	err := ts.Config.Load(path)
	require.NoError(t, err)

	return ts
}

func (ts *TestSetup) BuildLogger(t *testing.T) *TestSetup {
	loggerSetup := NewLogger(ts.Config)
	loggerSetup.Start()
	ts.Log = loggerSetup.Log

	return ts
}

func (ts *TestSetup) BuildDB(t *testing.T, c *config.Config) *TestSetup {
	if c == nil {
		c = ts.Config
	}
	db := NewDB(ts.Config)
	err := db.StartConn()
	ts.DB = db.DB
	require.NoError(t, err)

	return ts
}

func (ts *TestSetup) BuildServer(t *testing.T) *TestSetup {
	gin.SetMode(gin.TestMode)

	server := NewServer(ts.DB, ts.Config, ts.Log)
	ts.Server = server.router

	return ts
}

func (ts TestSetup) Serve(recorder *httptest.ResponseRecorder, request *http.Request) {
	ts.Server.ServeHTTP(recorder, request)
}
