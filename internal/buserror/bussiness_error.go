package buserror

import (
	"fmt"
	"net/http"
)

var (
	ErrInternalError error	= BusinessError{
		BusStatus: 55,
		HttpStatus: http.StatusInternalServerError,
		Message: "Internal Server Error",
	}
	ErrBadRequest error	= BusinessError{
		BusStatus: 56,
		HttpStatus: http.StatusBadRequest,
		Message: "Bad Request", 
	}
	ErrServiceUnavailable error	= BusinessError{
		BusStatus: 55,
		HttpStatus: http.StatusServiceUnavailable,
		Message: "Service Unavailable", 
	}
	ErrUnAuthorize error = BusinessError{
		BusStatus: 12,
		HttpStatus: 401,
		Message: "this user not allow to use this function",
	}
)

//implement 'error' interface 
type BusinessError struct {
	BusStatus int
	HttpStatus int 
	Message string
}

func (e BusinessError) Error() string {
	return fmt.Sprintf("%d : %s", e.BusStatus, e.Message)
}