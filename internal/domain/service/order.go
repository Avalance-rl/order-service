package service

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

type Cache interface {
	SelectOrders(ctx context.Context, id string) ([]model.Order, error)
}

type Service struct {
	logger     *zap.Logger
	repository Repository
	cache      Cache
}

func NewOrderService(logger *zap.Logger, repository Repository, cache Cache) *Service {
	return &Service{
		logger:     logger,
		repository: repository,
		cache:      cache,
	}
}

func (uc *Service) CreateOrder(ctx context.Context, order model.Order) (model.Order, error) {
	price, err := uc.GetTotalPrice(ctx, order.ProductList)
	if err != nil {
		return model.Order{}, NewError(ErrInternalFailure, err)
	}

	order.TotalPrice = price
	writtenOrder, err := uc.repository.InsertOrder(ctx, order)
	if err != nil {
		return model.Order{}, NewError(ErrInternalFailure, err)
	}

	return writtenOrder, nil
}

func (uc *Service) GetOrders(ctx context.Context, id string) ([]model.Order, error) {
	orders, err := uc.repository.SelectOrders(ctx, id)
	if err != nil {
		return nil, NewError(ErrInternalFailure, err)
	}

	return orders, nil
}

func (uc *Service) UpdateOrderStatus(ctx context.Context, id string) (model.OrderStatus, error) {
	status, err := uc.repository.UpdateOrderStatus(ctx, id)
	if err != nil {
		return "", NewError(ErrInternalFailure, err)
	}

	return status, nil
}

func (uc *Service) ConfirmOrder(ctx context.Context, id string) (model.OrderStatus, error) {
	status, err := uc.repository.UpdateOrderStatusToConfirm(ctx, id)
	if err != nil {
		return "", NewError(ErrInternalFailure, err)
	}

	return status, nil
}

func (uc *Service) GetTotalPrice(ctx context.Context, productList []string) (uint, error) {
	price, err := uc.repository.GetTotalPrice(ctx, productList)
	if err != nil {
		return 0, NewError(ErrInternalFailure, err)
	}

	return price, nil
}

func (uc *Service) GetTotalPriceByID(ctx context.Context, id string) (uint, error) {
	price, err := uc.repository.GetTotalPriceByID(ctx, id)
	if err != nil {
		return 0, NewError(ErrInternalFailure, err)
	}

	return price, nil
}
