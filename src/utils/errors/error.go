package errors

import (
	"fmt"
	"net/http"
)

type RestError interface {
	Message() string
	Status() int
	Error() string
}

type restError struct {
	ErrMessage string `json:"message"`
	ErrStatus  int    `json:"status"`
	ErrError   string `json:"error"`
}

func (e restError) Error() string {
	return fmt.Sprintf("message: %s - status: %d - error: %s",
		e.ErrMessage, e.ErrStatus, e.ErrError)
}

func (e restError) Message() string {
	return e.ErrMessage
}

func (e restError) Status() int {
	return e.ErrStatus
}

func NewRestError(message string, status int, err string, causes []interface{}) RestError {
	return restError{
		ErrMessage: message,
		ErrStatus:  status,
		ErrError:   err,
	}
}

func NewBadRequestError(message string) RestError {
	return restError{
		ErrMessage: message,
		ErrStatus:  http.StatusBadRequest,
		ErrError:   "bad_request",
	}
}

func NewNotFoundError(message string) RestError {
	return restError{
		ErrMessage: message,
		ErrStatus:  http.StatusNotFound,
		ErrError:   "not_found",
	}
}

func NewUnauthorizedError(message string) RestError {
	return restError{
		ErrMessage: message,
		ErrStatus:  http.StatusUnauthorized,
		ErrError:   "unauthorized",
	}
}

func NewInternalServerError(message string) RestError {
	result := restError{
		ErrMessage: message,
		ErrStatus:  http.StatusInternalServerError,
		ErrError:   "internal_server_error",
	}
	return result
}
