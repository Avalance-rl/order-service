package order

import desc "github.com/Avalance-rl/order-service/proto/pkg/order_v1"

type Repository interface {
	GetOrders(id string)
}

type Usecase struct {
	repository Repository
}

func NewOrderService(repository Repository) *Usecase {
	return &Usecase{
		repository: repository,
	}
}

func (uc *Usecase) GetOrders(id string) (*desc.Order, error) {
	panic("implement me")
}
