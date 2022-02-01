package errors

type APIError struct {
	Message string
}

func (err *APIError) Error() string {
	return err.Message
}

type ValidationError struct {
	*APIError
}

type TokenExpiredError struct {
	*APIError
}

type UnauthenticatedError struct {
	*APIError
}
