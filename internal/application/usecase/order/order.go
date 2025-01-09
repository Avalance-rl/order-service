package order

import (
	"context"

	"github.com/Avalance-rl/order-service/internal/domain/model"
	repoModel "github.com/Avalance-rl/order-service/internal/infrastructure/db/order/model"
	desc "github.com/Avalance-rl/order-service/proto/pkg/order_v1"
	"go.uber.org/zap"
)

type Repository interface {
	CreateOrder(ctx context.Context, order repoModel.Order) (model.Order, error)
	GetOrders(ctx context.Context, id string) ([]model.Order, error)
}

type Usecase struct {
	logger     *zap.Logger
	repository Repository
}

func NewOrderService(logger *zap.Logger, repository Repository) *Usecase {
	return &Usecase{
		logger:     logger,
		repository: repository,
	}
}

func (uc *Usecase) GetOrders(id string) (*desc.Order, error) {
	panic("implement me")
}
