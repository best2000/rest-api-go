package handler

import (
	"context"
	"net/http"

	"github.com/best2000/rest-api-go/logger"
	"github.com/rs/xid"
	"go.uber.org/zap"
)

func RequestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// create a correlation ID for the request
		correlationID := xid.New().String()

		//add X-Correlation-ID to context
		ctx := context.WithValue(r.Context(), "correlation_id", correlationID)

		//add header 
		w.Header().Add("X-Correlation-ID", correlationID)

		//create a child logger from main logger, add the correlation ID to the child
		childLogger := logger.Get().With(zap.String(string("correlation_id"), correlationID))

		//attach logger to context
		ctx = context.WithValue(ctx, "logger", childLogger)

		//attach context it to request
		next.ServeHTTP(w, r.WithContext(ctx))	//call next handler
	})
}
