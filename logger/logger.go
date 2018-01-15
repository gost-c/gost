package logger

import (
	"github.com/gost-c/gost/internal/utils"
	"github.com/sirupsen/logrus"
)

// Logger is logger instance
var Logger *logrus.Logger

func init() {
	log := logrus.New()
	if utils.GetEnvOrDefault("ENV", "prod") == "debug" {
		log.SetLevel(logrus.DebugLevel)
	}
	Logger = log
}
