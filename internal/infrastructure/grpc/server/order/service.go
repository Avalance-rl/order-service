package order

import (
	desc "github.com/Avalance-rl/order-service/proto/pkg/order_v1"
)

type UsecaseOrder interface {
	GetOrders(id string) (*desc.Order, error)
}

type Implementation struct {
	desc.UnimplementedOrderServiceServer
	usecaseOrder UsecaseOrder
}

func NewImplementation(usecaseOrder UsecaseOrder) *Implementation {
	return &Implementation{
		usecaseOrder: usecaseOrder,
	}
}
