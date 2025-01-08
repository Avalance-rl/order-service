package main

import (
	"context"

	"github.com/Avalance-rl/order-service/internal/application"
	"github.com/Avalance-rl/order-service/internal/logger"
	"go.uber.org/zap"
)

func main() {
	l := logger.InitLogger()
	defer func() {
		err := l.Sync()
		if err != nil {
			panic(err)
		}
	}()

	ctx := context.Background()

	a, err := application.NewApp(ctx, l)
	if err != nil {
		l.Fatal("failed to init app: %s", zap.Error(err))
	}

	err = a.Run()
	if err != nil {
		l.Fatal("failed to run app: %s", zap.Error(err))
	}
}
