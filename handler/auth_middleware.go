package handler

import (
	"net/http"

	"github.com/best2000/rest-api-go/logger"
)

func AuthMid(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log := logger.FromCtx(r.Context())

		log.Info("in auth middleware")

		tkn := r.Header.Get("token")
		log.Info("token="+tkn)

		next.ServeHTTP(w, r)
	})
}
