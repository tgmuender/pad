package logger

import "go.uber.org/zap"

// logger is a global logger instance
var logger *zap.Logger

// Get returns a logger instance
func Get() *zap.Logger {
	if logger == nil {
		logger = zap.Must(zap.NewDevelopment())
		defer logger.Sync()
		logger.Debug("Logger initialized")
	}
	return logger
}
