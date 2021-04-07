package main

import (
	"github.com/ayrtonsato/video-catalog-golang/internal/config"
	"github.com/ayrtonsato/video-catalog-golang/internal/setup"
	"log"
)

func main() {
	c := config.Config{}
	err := c.Load()
	if err != nil {
		log.Fatalf("config: failed to load config: %v", err.Error())
	}

	db := setup.NewDB(&c)
	err = db.StartConn()
	if err != nil {
		log.Fatalf("db: failed to start connection: %v", err.Error())
	}

	server := setup.NewServer(db.DB, &c)

	err = server.Start()
	if err != nil {
		log.Fatalf("gin-server: failed to start gin: %v", err.Error())
	}
}
