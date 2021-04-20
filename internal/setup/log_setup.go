package setup

import (
	"flag"
	"go.uber.org/zap"
	"log"
)

type Logger struct {
	conf *Config
	Log  *zap.SugaredLogger
}

func NewLogger(conf *Config) Logger {
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

	// if the system is under testing, disable logs with nop constructor
	if flag.Lookup("test.v") != nil {
		logger = zap.NewNop()
	}

	if err != nil {
		log.Fatalln(err)
	}
	defer logger.Sync()

	l.Log = logger.Sugar()
}
