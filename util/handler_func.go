package util

import (
	"errors"
	"log/slog"
	"net/http"
)

// custom handlerfunc that implement 'http.handler' interface
type HandlerFunc func(w ResponseWriter, r *http.Request) error

func (f HandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	//call main handler
	err := f(ResponseWriter{ResponseWriter: w}, r)

	//error handling...
	if err != nil {
		slog.Error(err.Error())

		//prepare error response
		//extract ApiError from error
		var apiErr ApiError
		if errors.As(err, &apiErr) {
			//ApiError type
			w.WriteHeader(apiErr.HttpStatus)
			w.Write([]byte(apiErr.Error()))
		} else {
			//std error
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
		return
	} else {
		//normal case, no error
		w.WriteHeader(http.StatusOK)
	}
}
