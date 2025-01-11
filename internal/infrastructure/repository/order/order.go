package order

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/Avalance-rl/order-service/internal/infrastructure/repository"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"

	"github.com/Avalance-rl/order-service/internal/infrastructure/repository/order/converter"

	"github.com/Avalance-rl/order-service/internal/domain/model"
	repoModel "github.com/Avalance-rl/order-service/internal/infrastructure/repository/order/model"
	"github.com/huandu/go-sqlbuilder"
)

func (r *Repository) InsertOrder(ctx context.Context, order model.Order) (model.Order, error) {
	if len(order.ProductList) == 0 {
		return model.Order{}, fmt.Errorf("%w: empty product list", repository.ErrInvalidInput)
	}

	repoOrder := converter.ToOrderFromService(&order)
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
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			switch pgErr.Code {
			case "23503":
				return model.Order{}, fmt.Errorf("%w: invalid customer_id or product_id", repository.ErrForeignKey)
			case "23505":
				return model.Order{}, fmt.Errorf("%w: order already exists", repository.ErrDuplicateKey)
			}
		}

		return model.Order{}, fmt.Errorf("insert order unexpected error: %w", err)
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
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("%w: orders not found", repository.ErrNotFound)
		}
		return nil, fmt.Errorf("execute query unexpected error: %w", err)
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

	if len(orders) == 0 {
		return nil, fmt.Errorf("%w: no orders found", repository.ErrNotFound)
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
		if errors.Is(err, pgx.ErrNoRows) {
			return "", fmt.Errorf("%w: order not found", repository.ErrNotFound)
		}
		return "", fmt.Errorf("update order status unexpected error: %w", err)
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
		if errors.Is(err, pgx.ErrNoRows) {
			return "", fmt.Errorf("%w: order not found", repository.ErrNotFound)
		}
		return "", fmt.Errorf("update order status unexpected error: %w", err)
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
		if errors.Is(err, pgx.ErrNoRows) {
			return 0, fmt.Errorf("%w: products not found", repository.ErrNotFound)
		}
		return 0, fmt.Errorf("get total price unexpected error: %w", err)
	}

	return totalPrice, nil
}

func (r *Repository) GetTotalPriceByID(ctx context.Context, id string) (uint, error) {
	if id == "" {
		return 0, fmt.Errorf("%w: empty id", repository.ErrInvalidID)
	}
	sb := sqlbuilder.NewSelectBuilder()
	sb.SetFlavor(sqlbuilder.PostgreSQL)
	sb.Select("total_price").From("orders").Where(sb.Equal("id", id))

	sql, args := sb.Build()

	row := r.pool.QueryRow(ctx, sql, args...)

	var totalPrice uint
	err := row.Scan(&totalPrice)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return 0, fmt.Errorf("%w: order with id %s not found", repository.ErrNotFound, id)
		}

		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "57014" {
				return 0, fmt.Errorf("%w: %v", repository.ErrQueryTimeout, err)
			}
		}

		return 0, fmt.Errorf("get total price unexpected error: %w", err)
	}

	return totalPrice, nil
}
