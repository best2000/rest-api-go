package util

import (
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
		w.Write([]byte(err.Error()))
	}
}
