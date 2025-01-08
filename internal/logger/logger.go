package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func InitLogger() *zap.Logger {
	config := zap.NewProductionConfig()
	config.OutputPaths = []string{
		"logs/application.log",
		"stderr",
	}

	config.ErrorOutputPaths = []string{
		"logs/error.log",
		"stderr",
	}

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
