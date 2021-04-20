package main

import (
	"log"

	"github.com/ayrtonsato/video-catalog-golang/internal/setup"
)

func main() {
	c := setup.Config{}
	err := c.Load(".")
	if err != nil {
		log.Fatalf("config: failed to load config: %v", err.Error())
	}

	loggerSetup := setup.NewLogger(&c)
	loggerSetup.Start()
	logger := loggerSetup.Log

	db := setup.NewDB(&c)
	err = db.StartConn()
	if err != nil {
		logger.Fatalf("db: failed to start connection: %v", err.Error())
	}

	server := setup.NewServer(db.DB, &c, logger)

	err = server.Start()
	if err != nil {
		logger.Fatalf("gin-server: failed to start gin: %v", err.Error())
	}
}
