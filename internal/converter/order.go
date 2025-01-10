package converter

import (
	"github.com/Avalance-rl/order-service/internal/domain/model"
	desc "github.com/Avalance-rl/order-service/proto/pkg/order_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToOrderFromUsecase(order *model.Order) *desc.Order {
	var status desc.OrderStatus
	switch order.OrderStatus {
	case model.Unpaid:
		status = desc.OrderStatus(0)
	case model.Paid:
		status = desc.OrderStatus(1)
	case model.Completed:
		status = desc.OrderStatus(2)
	}
	return &desc.Order{
		Id:          order.ID,
		CustomerId:  order.CustomerID,
		Status:      status,
		ProductList: order.ProductList,
		TotalPrice:  uint64(order.TotalPrice),
		CreatedAt:   timestamppb.New(order.CreatedAt),
		UpdatedAt:   timestamppb.New(order.UpdatedAt),
	}
}

func ToOrderFromDesc(order *desc.Order) *model.Order {
	var status model.OrderStatus
	switch order.Status {
	case desc.OrderStatus(0):
		status = model.Unpaid
	case desc.OrderStatus(1):
		status = model.Paid
	case desc.OrderStatus(2):
		status = model.Completed
	}
	return &model.Order{
		ID:          order.Id,
		CustomerID:  order.CustomerId,
		OrderStatus: status,
		ProductList: order.ProductList,
		TotalPrice:  uint(order.TotalPrice),
		CreatedAt:   order.CreatedAt.AsTime(),
		UpdatedAt:   order.UpdatedAt.AsTime(),
	}
}
