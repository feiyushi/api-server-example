package api

import "fmt"

// Error is the API error type
type Error struct {
	Code    ErrorCode `yaml:"code"`
	Message string    `yaml:"message"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("error code: %s; message: %s", e.Code, e.Message)
}

// ErrorResponse is the payload sent back to user in case of error
type ErrorResponse struct {
	Error Error `yaml:"error"`
}

// NewErrorResponse initialized an error response with error code and message
func NewErrorResponse(code ErrorCode, message string) *ErrorResponse {
	return &ErrorResponse{
		Error: Error{
			Code:    code,
			Message: message,
		},
	}
}
