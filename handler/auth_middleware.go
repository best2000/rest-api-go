package handler

import (
	"log/slog"
	"net/http"
)

func AuthMid(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Info("in auth middleware")

		tkn := r.Header.Get("token")
		slog.Info("token="+tkn)

		next.ServeHTTP(w, r)
	})
}
