package setup

import (
	"github.com/ayrtonsato/video-catalog-golang/internal/config"
	"go.uber.org/zap"
	"log"
)

type Logger struct {
	conf *config.Config
	Log  *zap.SugaredLogger
}

func NewLogger(conf *config.Config) Logger {
	return Logger{
		conf: conf,
	}
}

func (l *Logger) Start() {
	var logger *zap.Logger
	var err error
	if l.conf.ServerMode == "development" {
		logger, err = zap.NewDevelopment()
	} else {
		logger, err = zap.NewProduction()
	}
	if err != nil {
		log.Fatalln(err)
	}
	defer logger.Sync()

	l.Log = logger.Sugar()
}
