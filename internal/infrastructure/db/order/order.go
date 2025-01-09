package order

import (
	"context"
	"fmt"
	"strings"

	"github.com/Avalance-rl/order-service/internal/domain/model"
	repoModel "github.com/Avalance-rl/order-service/internal/infrastructure/db/order/model"
	"github.com/huandu/go-sqlbuilder"
)

func (r *Repository) CreateOrder(ctx context.Context, order repoModel.Order) (model.Order, error) {
	productList := "{" + strings.Join(order.ProductList, ",") + "}"

	sb := sqlbuilder.NewInsertBuilder()
	sb.SetFlavor(sqlbuilder.PostgreSQL)
	sb.InsertInto("orders").
		Cols(
			"id",
			"customer_id",
			"order_status",
			"product_list",
			"total_price",
			"created_at",
			"updated_at",
		).
		Values(
			order.ID,
			order.CustomerID,
			order.OrderStatus,
			productList,
			order.TotalPrice,
			order.CreatedAt,
			order.UpdatedAt,
		).
		SQL("RETURNING id, created_at, updated_at")
	sql, args := sb.Build()

	row := r.pool.QueryRow(ctx, sql, args...)

	var returnedOrder model.Order
	err := row.Scan(
		&returnedOrder.ID,
		&returnedOrder.CreatedAt,
		&returnedOrder.UpdatedAt,
	)
	returnedOrder.CustomerID = order.CustomerID
	returnedOrder.OrderStatus = order.OrderStatus
	returnedOrder.ProductList = order.ProductList
	returnedOrder.TotalPrice = order.TotalPrice

	if err != nil {
		return model.Order{}, fmt.Errorf("insert order: %w", err)
	}

	return returnedOrder, nil
}

func (r *Repository) GetOrders(ctx context.Context, id string) ([]model.Order, error) {
	sb := sqlbuilder.NewSelectBuilder()
	sb.SetFlavor(sqlbuilder.PostgreSQL)
	sb.Select("*").From("orders").Where(sb.Equal("id", id))
	sql, args := sb.Build()

	rows, err := r.pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("execute query: %w", err)
	}
	defer rows.Close()

	var orders []model.Order
	for rows.Next() {
		var order model.Order
		err = rows.Scan(
			&order.ID,
			&order.CustomerID,
			&order.OrderStatus,
			&order.ProductList,
			&order.TotalPrice,
			&order.CreatedAt,
			&order.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("scan order: %w", err)
		}
		orders = append(orders, order)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate rows: %w", err)
	}

	return orders, nil
}

func (r *Repository) UpdateOrderStatus(ctx context.Context, id string) (model.OrderStatus, error) {
	sb := sqlbuilder.NewUpdateBuilder()
	sb.SetFlavor(sqlbuilder.PostgreSQL)
	sb.Update("orders").
		Set("order_status", string(repoModel.Paid)).
		Where(sb.Equal("id", id)).
		SQL("RETURNING order_status")
	sql, args := sb.Build()

	row := r.pool.QueryRow(ctx, sql, args...)

	var orderStatus model.OrderStatus
	err := row.Scan(
		&orderStatus,
	)
	if err != nil {
		return "", fmt.Errorf("update order status: %w", err)
	}

	return orderStatus, nil
}

func (r *Repository) ConfirmOrder(ctx context.Context, id string) (model.OrderStatus, error) {
	sb := sqlbuilder.NewUpdateBuilder()
	sb.SetFlavor(sqlbuilder.PostgreSQL)
	sb.Update("orders").
		Set("order_status", string(repoModel.Completed)).
		Where(sb.Equal("id", id)).
		SQL("RETURNING order_status")
	sql, args := sb.Build()

	row := r.pool.QueryRow(ctx, sql, args...)

	var orderStatus model.OrderStatus
	err := row.Scan(
		&orderStatus,
	)
	if err != nil {
		return "", fmt.Errorf("confirm order: %w", err)
	}

	return orderStatus, nil
}

func (r *Repository) GetTotalPrice(ctx context.Context, id string) (uint, error) {
	sb := sqlbuilder.NewSelectBuilder()
	sb.SetFlavor(sqlbuilder.PostgreSQL)

	sb.Select("SUM(p.price)").
		From("orders o").
		Join("products p", "p.id = ANY(o.product_list)").
		Where(sb.Equal("o.id", id))

	sql, args := sb.Build()

	row := r.pool.QueryRow(ctx, sql, args...)

	var totalPrice uint
	err := row.Scan(&totalPrice)
	if err != nil {
		return 0, fmt.Errorf("get total price: %w", err)
	}

	return totalPrice, nil
}
