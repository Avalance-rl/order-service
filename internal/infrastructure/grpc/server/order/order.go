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
	orderFromDomain, err := i.serviceOrder.CreateOrder(ctx, *converter.ToOrderFromDesc(order))
	if err != nil {
		return nil, err
	}
	order = converter.ToOrderFromService(&orderFromDomain)
	return &desc.CreateOrderResponse{
		Order: order,
	}, nil
}

func (i *Implementation) GetOrders(
	ctx context.Context,
	req *desc.GetOrdersRequest,
) (*desc.GetOrdersResponse, error) {
	ordersFromService, err := i.serviceOrder.GetOrders(ctx, req.GetCustomerId())
	if err != nil {
		return nil, err
	}
	orders := make([]*desc.Order, len(ordersFromService))
	for _, order := range ordersFromService {
		orders = append(orders, converter.ToOrderFromService(&order))
	}
	return &desc.GetOrdersResponse{
		Orders: orders,
	}, nil
}

func (i *Implementation) UpdateOrderStatus(
	ctx context.Context,
	req *desc.UpdateOrderStatusRequest,
) (*desc.UpdateOrderStatusResponse, error) {
	orderStatusFromService, err := i.serviceOrder.UpdateOrderStatus(ctx, req.GetOrderId())
	if err != nil {
		return nil, err
	}
	status := converter.ToOrderStatusFromService(orderStatusFromService)
	return &desc.UpdateOrderStatusResponse{
		OrderStatus: status,
	}, nil
}

func (i *Implementation) ConfirmOrder(
	ctx context.Context,
	req *desc.ConfirmOrderRequest,
) (*desc.ConfirmOrderResponse, error) {
	orderStatusFromService, err := i.serviceOrder.ConfirmOrder(ctx, req.GetOrderId())
	if err != nil {
		return nil, err
	}

	status := converter.ToOrderStatusFromService(orderStatusFromService)
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
	price, err := i.serviceOrder.GetTotalPriceByID(ctx, req.GetOrderId())
	if err != nil {
		return nil, err
	}
	return &desc.GetTotalPriceResponse{
		Price: uint64(price),
	}, nil
}
