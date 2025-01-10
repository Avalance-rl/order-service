package usecase

import (
	"context"
	"github.com/Avalance-rl/order-service/internal/domain/model"
	"go.uber.org/zap"
)

type Repository interface {
	InsertOrder(ctx context.Context, order model.Order) (model.Order, error)
	SelectOrders(ctx context.Context, id string) ([]model.Order, error)
	UpdateOrderStatus(ctx context.Context, id string) (model.OrderStatus, error)
	UpdateOrderStatusToConfirm(ctx context.Context, id string) (model.OrderStatus, error)
	GetTotalPrice(ctx context.Context, productList []string) (uint, error)
	GetTotalPriceByID(ctx context.Context, id string) (uint, error)
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
	price, err := uc.GetTotalPrice(ctx, order.ProductList)
	if err != nil {
		return model.Order{}, err
	}

	order.TotalPrice = price
	writtenOrder, err := uc.repository.InsertOrder(ctx, order)
	if err != nil {
		return model.Order{}, err
	}

	return writtenOrder, nil
}

func (uc *Usecase) GetOrders(ctx context.Context, id string) ([]model.Order, error) {
	orders, err := uc.repository.SelectOrders(ctx, id)
	if err != nil {
		return nil, err
	}

	return orders, nil
}

func (uc *Usecase) UpdateOrderStatus(ctx context.Context, id string) (model.OrderStatus, error) {
	status, err := uc.repository.UpdateOrderStatus(ctx, id)
	if err != nil {
		return "", err
	}

	return status, nil
}

func (uc *Usecase) ConfirmOrder(ctx context.Context, id string) (model.OrderStatus, error) {
	status, err := uc.repository.UpdateOrderStatusToConfirm(ctx, id)
	if err != nil {
		return "", err
	}

	return status, nil
}

func (uc *Usecase) GetTotalPrice(ctx context.Context, productList []string) (uint, error) {
	price, err := uc.repository.GetTotalPrice(ctx, productList)
	if err != nil {
		return 0, err
	}

	return price, nil
}

func (uc *Usecase) GetTotalPriceByID(ctx context.Context, id string) (uint, error) {
	price, err := uc.repository.GetTotalPriceByID(ctx, id)
	if err != nil {
		return 0, err
	}

	return price, nil
}
