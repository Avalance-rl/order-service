package main

import (
	"context"
	"github.com/Avalance-rl/order-service/internal/application"
	"github.com/Avalance-rl/order-service/internal/logger"
	"go.uber.org/zap"
)

func main() {
	logger := logger.InitLogger()
	defer logger.Sync()

	ctx := context.Background()

	a, err := application.NewApp(ctx, logger)
	if err != nil {
		logger.Fatal("failed to init app: %s", zap.Error(err))
	}

	err = a.Run()
	if err != nil {
		logger.Fatal("failed to run app: %s", zap.Error(err))
	}
}
