package application

import (
	"context"
	usecaseOrder "github.com/Avalance-rl/order-service/internal/application/usecase/order"
	"github.com/Avalance-rl/order-service/internal/config"
	orderRepo "github.com/Avalance-rl/order-service/internal/infrastructure/db/order"
	grpcServer "github.com/Avalance-rl/order-service/internal/infrastructure/grpc/server/order"
	"go.uber.org/zap"
	"os"
)

type provider struct {
	config          config.Config
	orderRepository usecaseOrder.Repository
	orderUsecase    grpcServer.UsecaseOrder
	orderImpl       *grpcServer.Implementation
	ctx             context.Context
	logger          *zap.Logger
}

func newServiceProvider(ctx context.Context, logger *zap.Logger) *provider {
	return &provider{ctx: ctx, logger: logger}
}

func (s *provider) Config() config.Config {
	if s.Config == nil {
		cfg, err := config.Load(os.Getenv("CONFIG_PATH"))
		if err != nil {
			panic(err)
		}
		s.config = cfg
	}

	return s.config
}

func (s *provider) OrderRepository() usecaseOrder.Repository {
	if s.orderRepository == nil {
		s.orderRepository, _ = orderRepo.NewRepository(s.Config().Database.Name, s.Config().Database.MaxConns, nil)
	}

	return s.orderRepository
}

func (s *provider) OrderUsecase() grpcServer.UsecaseOrder {
	if s.orderUsecase == nil {
		s.orderUsecase = usecaseOrder.NewOrderService(
			s.logger,
			s.OrderRepository(),
		)
	}

	return s.orderUsecase
}

func (s *provider) OrderImpl() *grpcServer.Implementation {
	if s.orderImpl == nil {
		s.orderImpl = grpcServer.NewImplementation(s.OrderUsecase())
	}

	return s.orderImpl
}
