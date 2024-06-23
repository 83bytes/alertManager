package cmd

import (
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
)

var DEFAULT_LOG_LEVEL = "INFO"

func setLogLevelE(log *logrus.Logger, ll string) error {

	switch strings.ToUpper(ll) {
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
		return fmt.Errorf("unsupported log-level %s; refer docs", ll)
	}

	return nil
}
