package logger

import (
	"github.com/sirupsen/logrus"
	"os"
)

// Logger is logger instance
var Logger *logrus.Logger

func init() {
	log := logrus.New()
	if os.Getenv("ENV") == "debug" {
		log.SetLevel(logrus.DebugLevel)
	}
	Logger = log
}
