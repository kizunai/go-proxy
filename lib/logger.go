package lib

import (
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	log "github.com/sirupsen/logrus"
)

var logger *log.Logger

func init() {
	// init logger
	logger = log.New()

	path := "./logs/proxy.log"
	writer, _ := rotatelogs.New(
		path+".%Y-%m-%d.log",
		rotatelogs.WithRotationTime(time.Duration(1)*time.Hour),
		rotatelogs.WithRotationTime(24),
	)

	logger.SetOutput(writer)
}

func ProxyLogger() *log.Logger {
	return logger
}
