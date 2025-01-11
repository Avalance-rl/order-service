package order

type APIError struct {
	Status  int
	Message string
}

//func FromError(err error) APIError {
//	var apiError APIError
//	var svcError service.Error
//	if errors.As(err, &svcError) {
//		apiError.Message = svcError.AppErr().Error()
//		svcError := svcError.SvcErr()
//		switch {
//		case errors.Is(svcError, service.ErrInternalFailure):
//			apiError.Status = codes.Internal
//		}
//	}
//}
