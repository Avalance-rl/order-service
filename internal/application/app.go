package application

import (
	"context"
	desc "github.com/Avalance-rl/order-service/proto/pkg/order_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"log"
	"log/slog"
	"net"
)

type App struct {
	serviceProvider *provider
	grpcServer      *grpc.Server
}

func (a *App) initServiceProvider(ctx context.Context) error {
	logger := &slog.Logger{}
	a.serviceProvider = newServiceProvider(ctx, logger)
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
