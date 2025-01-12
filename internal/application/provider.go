package application

import (
	"context"
	"github.com/Avalance-rl/order-service/internal/infrastructure/cache/redis"
	orderRepo "github.com/Avalance-rl/order-service/internal/infrastructure/repository/pgx/order"
	"os"

	serviceOrder "github.com/Avalance-rl/order-service/internal/domain/service"

	"github.com/Avalance-rl/order-service/internal/config"
	grpcServer "github.com/Avalance-rl/order-service/internal/infrastructure/grpc/server/order"
	"go.uber.org/zap"
)

type provider struct {
	config          config.Config
	orderRepository serviceOrder.Repository
	orderCache      serviceOrder.Cache
	orderService    grpcServer.ServiceOrder
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

func (p *provider) OrderRepository() serviceOrder.Repository {
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

func (p *provider) OrderCache() serviceOrder.Cache {
	if p.orderCache == nil {

		cache, err := redis.NewCache(
			p.Config().Redis.Address,
			p.Config().Redis.Password,
		)
		if err != nil {
			p.logger.Fatal("orderCache error", zap.Error(err))
		}
		p.orderCache = cache
	}

	return p.orderCache
}

func (p *provider) OrderService() grpcServer.ServiceOrder {
	if p.orderService == nil {
		p.orderService = serviceOrder.NewOrderService(
			p.logger,
			p.OrderRepository(),
			p.OrderCache(),
		)
	}

	return p.orderService
}

func (p *provider) OrderImpl() *grpcServer.Implementation {
	if p.orderImpl == nil {
		p.orderImpl = grpcServer.NewImplementation(p.OrderService())
	}

	return p.orderImpl
}
