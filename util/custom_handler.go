package util

import (
	"errors"
	"net/http"

	"go.uber.org/zap"
)

// custom handlerfunc that implement 'http.handler' interface
type HandlerFunc func(w ResponseWriter, r *http.Request) error

func (f HandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//get logger from context
	log := r.Context().Value("logger").(*zap.Logger)

	//call main handler
	err := f(ResponseWriter{ResponseWriter: w}, r)

	//error handling...
	if err != nil {
		log.Error(err.Error(), zap.Error(err))

		var busErr BusinessError
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
