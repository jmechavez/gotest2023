package errorCust

import "net/http"

// AppError represents a custom application error.
type AppError struct {
	Code    int    `json:",omitempty"` // HTTP status code
	Message string `json:"message"`    // Error message
}

// AsMessage returns a new AppError with only the error message.
func (e AppError) AsMessage() *AppError {
	return &AppError{
		Message: e.Message,
	}
}

// Error implements the error interface.
func (e *AppError) Error() string {
	return e.Message
}

// NewNotFoundError creates a new AppError for not found scenarios.
func NewNotFoundError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusNotFound,
	}
}

// NewUnexpectedError creates a new AppError for unexpected/internal server errors.
func NewUnexpectedError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusInternalServerError,
	}
}

func NewValidationError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusUnprocessableEntity,
	}
}
