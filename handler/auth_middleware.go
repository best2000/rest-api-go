package handler

import (
	"log/slog"
	"net/http"
)

func AuthMid(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Info("in auth middleware")
		//set ctx timeout
		// ctx, cancle := context.WithTimeout(r.Context(), time.Millisecond * 1)

		// defer cancle()
		next.ServeHTTP(w, r)
	})
}
