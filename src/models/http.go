package models

type AppStatusCode int64

// int mapping
const (
	Success             AppStatusCode = 0
	AuthenticationError               = 1
	AuthorizationError                = 2
	TokenExpiredError                 = 3
	ValidationError                   = 4
	GeneralError                      = 5
)

type HttpResponse struct {
	Code    AppStatusCode `json:"code"`
	Message string        `json:"message"`
	Payload interface{}   `json:"payload"`
}

type UserSignInHttpResponse struct {
	*HttpResponse
	Payload UserSignInResponse `json:"payload"`
}
