package repository

import "errors"

var (
	ErrNotFound     = errors.New("resource not found")
	ErrInvalidID    = errors.New("invalid id format")
	ErrDuplicateKey = errors.New("duplicate key violation")
	ErrQueryTimeout = errors.New("query timeout")
	ErrInvalidInput = errors.New("invalid input")
	ErrForeignKey   = errors.New("foreign key violation")
)

func IsErrNotFound(err error) bool {
	return errors.Is(err, ErrNotFound)
}

func IsErrInvalidID(err error) bool {
	return errors.Is(err, ErrInvalidID)
}

func IsErrDuplicateKey(err error) bool {
	return errors.Is(err, ErrDuplicateKey)
}

func IsErrQueryTimeout(err error) bool {
	return errors.Is(err, ErrQueryTimeout)
}

func IsErrInvalidInput(err error) bool {
	return errors.Is(err, ErrInvalidInput)
}

func IsErrForeignKey(err error) bool {
	return errors.Is(err, ErrForeignKey)
}
