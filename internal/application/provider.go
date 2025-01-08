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

func (p *provider) Config() config.Config {
	cfg, err := config.Load(os.Getenv("CONFIG_PATH"))
	if err != nil {
		panic(err)
	}
	p.config = cfg

	return p.config
}

func (p *provider) OrderRepository() usecaseOrder.Repository {
	if p.orderRepository == nil {

		rep, err := orderRepo.NewRepository(
			p.Config().Database.Host,
			p.Config().Database.Port,
			p.Config().Database.SSLMode,
			p.Config().Database.User,
			p.Config().Database.Password,
			p.Config().Database.Name,
			p.Config().Database.MaxConns,
			p.logger,
		)
		if err != nil {
			p.logger.Fatal("orderRepository error", zap.Error(err))
		}
		p.orderRepository = rep
	}

	return p.orderRepository
}

func (p *provider) OrderUsecase() grpcServer.UsecaseOrder {
	if p.orderUsecase == nil {
		p.orderUsecase = usecaseOrder.NewOrderService(
			p.logger,
			p.OrderRepository(),
		)
	}

	return p.orderUsecase
}

func (p *provider) OrderImpl() *grpcServer.Implementation {
	if p.orderImpl == nil {
		p.orderImpl = grpcServer.NewImplementation(p.OrderUsecase())
	}

	return p.orderImpl
}
