package model

import "time"

type OrderStatus string

const (
	unpaid    = OrderStatus("UNPAID")
	paid      = OrderStatus("PAID")
	completed = OrderStatus("COMPLETED")
)

type Order struct {
	ID          string
	CustomerID  string
	OrderStatus string
	ProductList []string
	TotalPrice  uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
