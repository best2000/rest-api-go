package handler

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func PrePost(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//pre handle...
		start := time.Now()

		//set request timeout (ctx timeout)
		ctx, cancel := context.WithTimeout(r.Context(), time.Second*3)
		defer cancel()

		//set ctx request id
		reqId := uuid.NewString()
		ctx = context.WithValue(ctx, "requestId", reqId)

		//call next handler...
		next.ServeHTTP(w, r.WithContext(ctx))
		//...

		//post handle...

		slog.Info("request id: " + reqId + ", elapsed time: " + time.Since(start).String())
	})
}
