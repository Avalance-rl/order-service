package order

import (
	"context"
	"github.com/Avalance-rl/order-service/internal/converter"
	desc "github.com/Avalance-rl/order-service/proto/pkg/order_v1"
)

func (i *Implementation) CreateOrder(
	ctx context.Context,
	req *desc.CreateOrderRequest,
) (*desc.CreateOrderResponse, error) {
	order := &desc.Order{
		CustomerId:  req.GetCustomerId(),
		ProductList: req.GetProductList(),
	}
	orderFromDomain, err := i.usecaseOrder.CreateOrder(ctx, *converter.ToOrderFromDesc(order))
	if err != nil {
		return nil, err
	}
	order = converter.ToOrderFromUsecase(&orderFromDomain)
	return &desc.CreateOrderResponse{
		Order: order,
	}, nil
}

func (i *Implementation) GetOrders(
	ctx context.Context,
	req *desc.GetOrdersRequest,
) (*desc.GetOrdersResponse, error) {
	ordersFromUsecase, err := i.usecaseOrder.GetOrders(ctx, req.GetCustomerId())
	if err != nil {
		return nil, err
	}
	orders := make([]*desc.Order, len(ordersFromUsecase))
	for _, order := range ordersFromUsecase {
		orders = append(orders, converter.ToOrderFromUsecase(&order))
	}
	return &desc.GetOrdersResponse{
		Orders: orders,
	}, nil
}

func (i *Implementation) UpdateOrderStatus(
	ctx context.Context,
	req *desc.UpdateOrderStatusRequest,
) (*desc.UpdateOrderStatusResponse, error) {
	orderStatusFromUsecase, err := i.usecaseOrder.UpdateOrderStatus(ctx, req.GetOrderId())
	if err != nil {
		return nil, err
	}
	status := converter.ToOrderStatusFromUsecase(orderStatusFromUsecase)
	return &desc.UpdateOrderStatusResponse{
		OrderStatus: status,
	}, nil
}

func (i *Implementation) ConfirmOrder(
	ctx context.Context,
	req *desc.ConfirmOrderRequest,
) (*desc.ConfirmOrderResponse, error) {
	orderStatusFromUsecase, err := i.usecaseOrder.ConfirmOrder(ctx, req.GetOrderId())
	if err != nil {
		return nil, err
	}

	status := converter.ToOrderStatusFromUsecase(orderStatusFromUsecase)
	return &desc.ConfirmOrderResponse{
		Order: &desc.Order{
			Status: status,
		},
	}, nil
}

func (i *Implementation) GetTotalPrice(
	ctx context.Context,
	req *desc.GetTotalPriceRequest,
) (*desc.GetTotalPriceResponse, error) {
	price, err := i.usecaseOrder.GetTotalPriceByID(ctx, req.GetOrderId())
	if err != nil {
		return nil, err
	}
	return &desc.GetTotalPriceResponse{
		Price: uint64(price),
	}, nil
}
