package model

import "time"

type OrderStatus string

const (
	Unpaid    = OrderStatus("UNPAID")
	Paid      = OrderStatus("PAID")
	Completed = OrderStatus("COMPLETED")
)

type Order struct {
	ID          string
	CustomerID  string
	OrderStatus OrderStatus
	ProductList []string
	TotalPrice  uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
