package util

import (
	"net/http"
)

var (
	ErrInternalError error	= ApiError{
		BusStatus: 55,
		HttpStatus: http.StatusInternalServerError,
		Message: "Internal Server Error",
	}
	ErrBadRequest error	= ApiError{
		BusStatus: 55,
		HttpStatus: http.StatusInternalServerError,
		Message: "Bad Request", 
	}
)

type ApiError struct {
	BusStatus int
	HttpStatus int 
	Message string
}

func (e ApiError) Error () string {
	return string(e.BusStatus) + " " + e.Message
}