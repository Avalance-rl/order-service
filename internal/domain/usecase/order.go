package usecase

import (
	"context"
	"github.com/Avalance-rl/order-service/internal/domain/model"
	"go.uber.org/zap"
)

type Repository interface {
	CreateOrder(ctx context.Context, order model.Order) (model.Order, error)
	GetOrders(ctx context.Context, id string) ([]model.Order, error)
	UpdateOrderStatus(ctx context.Context, id string) (model.OrderStatus, error)
	ConfirmOrder(ctx context.Context, id string) (model.OrderStatus, error)
	GetTotalPrice(ctx context.Context, id string) (uint, error)
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

func (uc *Usecase) CreateOrder(ctx context.Context, order model.Order) (model.Order, error) {
	writtenOrder, err := uc.repository.CreateOrder(ctx, order)
	if err != nil {
		return model.Order{}, err
	}
	return writtenOrder, nil
}

func (uc *Usecase) GetOrders(ctx context.Context, id string) (model.Order, error) {
	order, err := uc.repository.GetOrders(ctx, id)
}

func (uc *Usecase) UpdateOrderStatus(ctx context.Context, id string) (model.OrderStatus, error) {
	panic("implement me")
}

func (uc *Usecase) ConfirmOrder(ctx context.Context, id string) (model.OrderStatus, error) {
	panic("implement me")
}

func (uc *Usecase) GetTotalPrice(ctx context.Context, id string) (uint, error) {
	panic("implement me")
}
