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

		var busErr BusError
		//check bussiness error
		if errors.As(err, &busErr) {
			//set bussiness error response
			w.WriteHeader(busErr.HttpStatus)
			w.Write([]byte(busErr.Error()))
		} else {
			//std error
			//set 500 Internal Server Error 
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
		return
	} else {
		//set 200 OK
		w.WriteHeader(http.StatusOK)
	}
}
