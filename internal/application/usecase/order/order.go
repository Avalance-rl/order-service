package order

import (
	desc "github.com/Avalance-rl/order-service/proto/pkg/order_v1"
	"log/slog"
)

type Repository interface {
	GetOrders(id string)
}

type Usecase struct {
	logger     *slog.Logger
	repository Repository
}

func NewOrderService(logger *slog.Logger, repository Repository) *Usecase {
	return &Usecase{
		logger:     logger,
		repository: repository,
	}
}

func (uc *Usecase) GetOrders(id string) (*desc.Order, error) {
	panic("implement me")
}
