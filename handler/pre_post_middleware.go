package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/best2000/rest-api-go/logger"
	"github.com/best2000/rest-api-go/value"
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

		//get request ID
		requestId := r.Header.Get(value.RequestIdHeaderKey)
		if requestId == "" {
			//generate a request ID for the request
			requestId = xid.New().String()
		}

		//add request ID header 
		w.Header().Add(value.RequestIdHeaderKey , requestId)

		//create a child logger from main logger then add the request ID to the child
		reqLogger := logger.Get().With(zap.String(string(value.RequestIdCtxKey), requestId))
		//attach logger to context
		ctx = context.WithValue(ctx, value.LoggerCtxKey, reqLogger)

		//add X-Request-ID to context
		ctx = context.WithValue(ctx, value.RequestIdCtxKey, requestId)

		//attach context it to request
		next.ServeHTTP(w, r.WithContext(ctx))	//call next handler
		//...

		//post handle...

		reqLogger.Info("end processing request",zap.String("elapse_time",time.Since(start).String()))
	})
}
