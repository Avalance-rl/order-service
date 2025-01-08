package order

import (
	desc "github.com/Avalance-rl/order-service/proto/pkg/order_v1"
	"go.uber.org/zap"
)

type Repository interface {
	GetOrders(id string)
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
