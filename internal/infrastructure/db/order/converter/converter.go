package converter

import (
	"github.com/Avalance-rl/order-service/internal/domain/model"
	repoModel "github.com/Avalance-rl/order-service/internal/infrastructure/db/order/model"
)

func ToOrderFromUsecase(order *model.Order) *repoModel.Order {
	return &repoModel.Order{
		ID:          order.ID,
		CustomerID:  order.CustomerID,
		OrderStatus: repoModel.OrderStatus(order.OrderStatus),
		ProductList: order.ProductList,
		TotalPrice:  order.TotalPrice,
		CreatedAt:   order.CreatedAt,
		UpdatedAt:   order.UpdatedAt,
	}
}

func ToOrderFromRepo(order *repoModel.Order) *model.Order {
	return &model.Order{
		ID:          order.ID,
		CustomerID:  order.CustomerID,
		OrderStatus: model.OrderStatus(order.OrderStatus),
		ProductList: order.ProductList,
		TotalPrice:  order.TotalPrice,
		CreatedAt:   order.CreatedAt,
		UpdatedAt:   order.UpdatedAt,
	}
}
