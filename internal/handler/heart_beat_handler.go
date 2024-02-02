package handler

import "net/http"

func Ping(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("pong"))
	return nil
}
