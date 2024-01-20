package util

import (
	"errors"
	"fmt"
	"net/http"
)

var (
	ErrInternalError error	= errors.New("internal server error")
	ErrBadRequest error	= errors.New("bad request")
)
//should try use ENUM instead!!!!!!!!

type ApiError struct {
	AppError error
	BussinessError error
}

//implement 'error' interface
func (e ApiError) Error() string {
	return fmt.Sprintf("Bussiness Error: %s \nApp Error: %s", 
	e.BussinessError.Error(), e.AppError.Error())
}

type ApiErrorResponse struct {
	Status int
	Message string
}

func MatchApiErrorResponse(e error) ApiErrorResponse {
	var apiErr ApiError = ApiError{}
	var	apiErrRes ApiErrorResponse = ApiErrorResponse{}
	if errors.As(e, &apiErr) {
		apiErrRes.Message = apiErr.Error()
		switch apiErr.BussinessError {
		case ErrBadRequest:
			apiErrRes.Status = http.StatusBadRequest
		case ErrInternalError:
			apiErrRes.Status = http.StatusBadRequest
		}
	}
	return apiErrRes
}