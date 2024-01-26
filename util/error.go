package util

import (
	"fmt"
	"net/http"
)

var (
	ErrInternalError error	= BusError{
		BusStatus: 55,
		HttpStatus: http.StatusInternalServerError,
		Message: "Internal Server Error",
	}
	ErrBadRequest error	= BusError{
		BusStatus: 56,
		HttpStatus: http.StatusBadRequest,
		Message: "Bad Request", 
	}
	ErrServiceUnavailable error	= BusError{
		BusStatus: 55,
		HttpStatus: http.StatusServiceUnavailable,
		Message: "Service Unavailable", 
	}
	ErrUnAuthorize error = BusError{
		BusStatus: 12,
		HttpStatus: 401,
		Message: "this user not allow to use this function",
	}
)

//implement 'error' interface 
type BusError struct {
	BusStatus int
	HttpStatus int 
	Message string
}

func (e BusError) Error() string {
	return fmt.Sprintf("%d : %s", e.BusStatus, e.Message)
}