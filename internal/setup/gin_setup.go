package setup

import (
	"database/sql"
	"fmt"

	"github.com/ayrtonsato/video-catalog-golang/internal/routes"
	"github.com/ayrtonsato/video-catalog-golang/pkg/logger"
	"github.com/gin-gonic/gin"
)

type Server struct {
	store  *sql.DB
	router *gin.Engine
	config *Config
	logger logger.Logger
}

func NewServer(store *sql.DB, config *Config, logger logger.Logger) Server {
	server := Server{
		store: store, config: config, logger: logger,
	}
	server.setupRouter()
	server.initRoutes()
	return server
}

func (s *Server) setupRouter() {
	router := gin.Default()
	s.router = router
}

func (s *Server) initRoutes() {
	routes.NewCategoryRoutes(s.router, s.store, s.logger).Routes()
}

func (s *Server) Start() error {
	addr := fmt.Sprintf("%v:%v", s.config.ServerAddress, s.config.Port)
	return s.router.Run(addr)
}
