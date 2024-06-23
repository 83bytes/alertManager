package logging

import (
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
)

type Logger struct {
	*logrus.Logger
}

var DEFAULT_LOG_LEVEL = "INFO"
var log = logrus.New()

func NewLogger(logLevel string) (*Logger, error) {

	switch strings.ToUpper(logLevel) {
	case "INFO":
		log.Info("info logs enabled")
		log.SetLevel(logrus.InfoLevel)
	case "DEBUG":
		log.Info("debug logs enabled")
		log.SetLevel(logrus.DebugLevel)
	case "ERROR":
		log.Info("error logs enabled")
		log.SetLevel(logrus.ErrorLevel)
	default:
		return nil, fmt.Errorf("unsupported log-level %s; refer docs", logLevel)
	}

	return &Logger{log}, nil
}

func GetLogger() *Logger {
	return &Logger{log}
}
