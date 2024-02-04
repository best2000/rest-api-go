package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func newChequeRouter() chi.Router {
	r := chi.NewRouter()
	
	r.Get("/smth", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("something in the way..."))
	})

	return r 
}