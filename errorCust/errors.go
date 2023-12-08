package errorCust

import "net/http"

type AppError struct {
	Code    int    `json:",omitempty"` //remove error message in restapi
	Message string `json:"messsage"`
}

func (e AppError) AsMessage() *AppError {
	return &AppError{
		Message: e.Message,
	}
}

// Error implements error.
func (*AppError) Error() string {
	panic("unimplemented")
}

func NewNotFoundError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusNotFound,
	}
}
func NewUnexpectedError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusInternalServerError,
	}
}
