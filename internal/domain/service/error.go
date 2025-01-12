package service

import "errors"

var (
	ErrBadRequest      = errors.New("bad request")
	ErrInternalFailure = errors.New("internal failure error")
	ErrNotFound        = errors.New("not found")
)

type Error struct {
	appErr error
	svcErr error
}

func NewError(appErr, svcErr error) *Error {
	return &Error{
		appErr: appErr,
		svcErr: svcErr,
	}
}

func (e *Error) AppErr() error {
	return e.appErr
}

func (e *Error) SvcErr() error {
	return e.svcErr
}

func (e *Error) Error() string {
	return errors.Join(e.appErr, e.svcErr).Error()
}
