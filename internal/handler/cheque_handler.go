package handler

import "net/http"

func GetCheque(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("pong"))
	return ni
}