package auth

import "fmt"

type CustomError struct {
	Code    int
	Message string
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}

type ErrorItem struct {
	Domain  string `json:"domain"`
	Reason  string `json:"reason"`
	Message string `json:"message"`
}

type ErrorGlobalItem struct {
	Errors  []ErrorItem `json:"errors"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
}

// Reference: https://firebase.google.com/docs/reference/rest/auth?hl=en#section-error-format
type ErrorResponse struct {
	Error ErrorGlobalItem `json:"error"`
}

func (e *ErrorResponse) ToCustomError() *CustomError {
	return &CustomError{
		Code:    e.Error.Code,
		Message: e.Error.Message,
	}
}
