package redis

import (
	"context"
	"github.com/Avalance-rl/order-service/internal/domain/model"
)

func (r *Redis) SelectOrders(ctx context.Context, id string) ([]model.Order, error) {
	panic("implement me")
}
