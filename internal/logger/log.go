package logger

import "go.uber.org/zap"

var logger *zap.Logger

func Get() *zap.Logger {
	if logger == nil {
		logger = zap.Must(zap.NewDevelopment())
		defer logger.Sync()
		logger.Debug("Logger initialized")
	}
	return logger
}
