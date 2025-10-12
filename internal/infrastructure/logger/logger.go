package logger

import (
	"go.uber.org/zap"
)

type Logger struct {
	*zap.SugaredLogger
}

func NewLogger() *Logger {
	config := zap.NewProductionConfig()
	config.Encoding = "console"
	logger, _ := config.Build()
	return &Logger{SugaredLogger: logger.Sugar()}
}
