package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func InitLogger() *zap.Logger {
	config := zap.NewProductionConfig()
	logFile, err := os.OpenFile("application.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	config.OutputPaths = []string{logFile.Name()}
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	logLevel := os.Getenv("LOG_LEVEL")
	if level, err := zapcore.ParseLevel(logLevel); err == nil {
		config.Level = zap.NewAtomicLevelAt(level)
	}

	l, err := config.Build()
	if err != nil {
		panic(err)
	}

	return l
}
