package model

import "time"

type OrderStatus string

const (
	unpaid    = OrderStatus("UNPAID")
	paid      = OrderStatus("PAID")
	completed = OrderStatus("COMPLETED")
)

type Order struct {
	ID          string    `db:"id"`
	CustomerID  string    `db:"customer_id"`
	OrderStatus string    `db:"order_status"`
	ProductList []string  `db:"product_list"`
	TotalPrice  uint      `db:"total_price"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}
