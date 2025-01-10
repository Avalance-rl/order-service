package order

import (
	"context"
	"fmt"
	"strings"

	"github.com/Avalance-rl/order-service/internal/infrastructure/db/order/converter"

	"github.com/Avalance-rl/order-service/internal/domain/model"
	repoModel "github.com/Avalance-rl/order-service/internal/infrastructure/db/order/model"
	"github.com/huandu/go-sqlbuilder"
)

func (r *Repository) InsertOrder(ctx context.Context, order model.Order) (model.Order, error) {
	repoOrder := converter.ToOrderFromUsecase(&order)

	productList := "{" + strings.Join(repoOrder.ProductList, ",") + "}"

	sb := sqlbuilder.NewInsertBuilder()
	sb.SetFlavor(sqlbuilder.PostgreSQL)
	sb.InsertInto("orders").
		Cols(
			"customer_id",
			"order_status",
			"product_list",
			"total_price",
			"created_at",
			"updated_at",
		).
		Values(
			repoOrder.CustomerID,
			repoModel.Unpaid,
			productList,
			repoOrder.TotalPrice,
			repoOrder.CreatedAt,
			repoOrder.UpdatedAt,
		).
		SQL("RETURNING id, created_at, updated_at")

	sql, args := sb.Build()

	row := r.pool.QueryRow(ctx, sql, args...)

	err := row.Scan(
		&repoOrder.ID,
		&repoOrder.CreatedAt,
		&repoOrder.UpdatedAt,
	)
	if err != nil {
		return model.Order{}, fmt.Errorf("insert order: %w", err)
	}

	return *converter.ToOrderFromRepo(repoOrder), nil
}

func (r *Repository) SelectOrders(ctx context.Context, id string) ([]model.Order, error) {
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

func (r *Repository) UpdateOrderStatusToConfirm(ctx context.Context, id string) (model.OrderStatus, error) {
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

func (r *Repository) GetTotalPrice(ctx context.Context, productList []string) (uint, error) {
	sb := sqlbuilder.NewSelectBuilder()
	sb.SetFlavor(sqlbuilder.PostgreSQL)
	sb.Select("SUM(price)").
		From("products").
		Where(sb.In("id", productList))

	sql, args := sb.Build()

	row := r.pool.QueryRow(ctx, sql, args...)

	var totalPrice uint
	err := row.Scan(&totalPrice)
	if err != nil {
		return 0, fmt.Errorf("get total price: %w", err)
	}

	return totalPrice, nil
}

func (r *Repository) GetTotalPriceByID(ctx context.Context, id string) (uint, error) {
	sb := sqlbuilder.NewSelectBuilder()
	sb.SetFlavor(sqlbuilder.PostgreSQL)
	sb.Select("total_price").From("orders").Where(sb.Equal("id", id))

	sql, args := sb.Build()

	row := r.pool.QueryRow(ctx, sql, args...)

	var totalPrice uint
	err := row.Scan(&totalPrice)
	if err != nil {
		return 0, fmt.Errorf("get total price: %w", err)
	}

	return totalPrice, nil
}
