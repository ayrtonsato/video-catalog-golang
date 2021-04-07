package setup

import (
	"database/sql"
	"fmt"
	"github.com/ayrtonsato/video-catalog-golang/internal/config"
	"github.com/gin-gonic/gin"
)

type Server struct {
	store *sql.DB
	router *gin.Engine
	config *config.Config
}

func NewServer(store *sql.DB, config *config.Config) Server {
	server := Server{
		store: store, config: config,
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
	// routes.NewCategoryRoutes(s.router, s.store)
	s.router.GET("/", func(ctx *gin.Context){
		ctx.JSON(200, "Hello")
	})
}

func (s *Server) Start() error {
	s.setupRouter()
	s.initRoutes()
	addr := fmt.Sprintf("%v:%v", s.config.ServerAddress, s.config.Port)
	return s.router.Run(addr)
}
