package order

import (
	"context"

	"github.com/Avalance-rl/order-service/internal/domain/model"
	desc "github.com/Avalance-rl/order-service/proto/pkg/order_v1"
)

type ServiceOrder interface {
	CreateOrder(ctx context.Context, order model.Order) (model.Order, error)
	GetOrders(ctx context.Context, id string) ([]model.Order, error)
	UpdateOrderStatus(ctx context.Context, id string) (model.OrderStatus, error)
	ConfirmOrder(ctx context.Context, id string) (model.OrderStatus, error)
	GetTotalPriceByID(ctx context.Context, id string) (uint, error)
}

type Implementation struct {
	desc.UnimplementedOrderServiceServer
	serviceOrder ServiceOrder
}

func NewImplementation(serviceOrder ServiceOrder) *Implementation {
	return &Implementation{
		serviceOrder: serviceOrder,
	}
}
