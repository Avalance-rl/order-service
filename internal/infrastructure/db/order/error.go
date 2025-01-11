package order

import "errors"

var (
	ErrNotFound          = errors.New("resource not found")
	ErrInvalidID         = errors.New("invalid id format")
	ErrDuplicateKey      = errors.New("duplicate key violation")
	ErrDeadlock          = errors.New("deadlock detected")
	ErrConnectionFailed  = errors.New("database connection failed")
	ErrQueryTimeout      = errors.New("query timeout")
	ErrTransactionFailed = errors.New("transaction failed")
	ErrInvalidInput      = errors.New("invalid input")
	ErrForeignKey        = errors.New("foreign key violation")
)
