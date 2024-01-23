package util

import (
	"errors"
	"net/http"
)

var (
	ErrInternalError error	= ApiErrorResponse{
		BusStatus: 15,
		HttpStatus: 500 ,
		Message: "internal server error",
	}
	ErrBadRequest error	= errors.New("bad request")
)
//should try use ENUM instead!!!!!!!!

// type ApiError struct {
// 	AppError error
// 	BussinessError error
// }

// func NewApiError(appErr error, busErr error) ApiError{
// 	return ApiError{
// 		AppError: appErr,
// 		BussinessError: busErr,
// 	}
// }

//implement 'error' interface
// func (e ApiError) Error() string {
// 	errString := "Error: "
// 	if e.BussinessError != nil {
// 		errString += e.BussinessError.Error()
// 	} 
// 	errString += ", " 
// 	if e.AppError != nil {
// 		errString += e.AppError.Error()
// 	} 
// 	return errString
// }
//+unwrap method
// func (e ApiError) Unwrap() error {
// 	return e.AppError
// }

type ApiErrorResponse struct {
	BusStatus int
	HttpStatus int 
	Message string
}

func (e ApiErrorResponse) Error () string {
	return e.Message
}

// func (e ApiError) PrepareErrorResponse(w http.ResponseWriter) {
// 	var	(
// 		status int
// 		message string
// 	)
// 	message = e.Error()
// 	if e.BussinessError != nil {
// 		switch e.BussinessError {
// 		case ErrBadRequest:
// 			status = http.StatusBadRequest
// 		case ErrInternalError:
// 			status = http.StatusInternalServerError
// 		}
// 	}
// 	res := ApiErrorResponse{
// 		Status: status,
// 		Message: message,
// 	}

// 	w.WriteHeader(res.Status)
// 	w.Write([]byte(res.Message))
// }