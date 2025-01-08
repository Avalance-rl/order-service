package application

import (
	"context"
	desc "github.com/Avalance-rl/order-service/proto/pkg/order_v1"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type App struct {
	serviceProvider *provider
	grpcServer      *grpc.Server
	logger          *zap.Logger
}

func NewApp(ctx context.Context, logger *zap.Logger) (*App, error) {
	a := &App{logger: logger}

	err := a.initDeps(ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (a *App) Run() error {
	return a.runGRPCServer()
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initServiceProvider,
		a.initGRPCServer,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initServiceProvider(ctx context.Context) error {
	a.serviceProvider = newServiceProvider(ctx, a.logger)
	return nil
}

func (a *App) initGRPCServer(_ context.Context) error {
	a.grpcServer = grpc.NewServer(grpc.Creds(insecure.NewCredentials()))

	reflection.Register(a.grpcServer)

	desc.RegisterOrderServiceServer(a.grpcServer, a.serviceProvider.OrderImpl())

	return nil
}

func (a *App) runGRPCServer() error {
	log.Printf("GRPC server is running on %s", a.serviceProvider.Config().GRPCServer.Address)

	list, err := net.Listen(
		"tcp",
		a.serviceProvider.Config().GRPCServer.Address+
			":"+a.serviceProvider.Config().GRPCServer.Port)

	if err != nil {
		return err
	}

	err = a.grpcServer.Serve(list)
	if err != nil {
		return err
	}

	return nil
}
