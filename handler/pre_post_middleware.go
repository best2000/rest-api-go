package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/best2000/rest-api-go/logger"
	"github.com/rs/xid"
	"go.uber.org/zap"
)

func PrePost(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//pre handle...
		start := time.Now()

		//set request timeout (ctx timeout)
		ctx, cancel := context.WithTimeout(r.Context(), time.Second*3)
		defer cancel()

		//generate a correlation ID for the request
		correlationID := xid.New().String()

		//add correlation ID header 
		w.Header().Add("X-Correlation-ID", correlationID)

		//create a child logger from main logger then add the correlation ID to the child
		reqLogger := logger.Get().With(zap.String(string("correlation_id"), correlationID))
		//attach logger to context
		ctx = context.WithValue(ctx, "logger", reqLogger)

		//add X-Correlation-ID to context
		ctx = context.WithValue(ctx, "correlation_id", correlationID)

		//attach context it to request
		next.ServeHTTP(w, r.WithContext(ctx))	//call next handler
		//...

		//post handle...

		reqLogger.Info("end processing request",zap.String("elapse_time",time.Since(start).String()))
	})
}
