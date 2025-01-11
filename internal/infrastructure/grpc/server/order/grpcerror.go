package order

import (
	"errors"
	"net/http"

	"github.com/Avalance-rl/order-service/internal/domain/service"
)

type APIError struct {
	Status  int
	Message string
}

func FromError(err error) APIError {
	var apiError APIError
	var svcError service.Error
	if errors.As(err, &svcError) {
		apiError.Message = svcError.AppErr().Error()
		svcErr := svcError.SvcErr()
		switch svcErr {
		case service.ErrInternalFailure:
			apiError.Status = http.StatusInternalServerError
		case service.ErrNotFound:
			apiError.Status = http.StatusNotFound
		case service.ErrBadRequest:
			apiError.Status = http.StatusBadRequest
		}
	}

	return apiError
}
