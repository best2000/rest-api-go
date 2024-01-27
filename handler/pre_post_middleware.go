package handler

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"time"

	"github.com/best2000/rest-api-go/logger"
	"github.com/best2000/rest-api-go/value"
	"github.com/rs/xid"
	"go.uber.org/zap"
)

func PrePost(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//pre handle.....
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
		w.Header().Add(value.RequestIdHeaderKey, requestId)

		//create a child logger from main logger then add the request ID to the child
		log := logger.Get().With(zap.String(string(value.RequestIdCtxKey), requestId))
		ctx = context.WithValue(ctx, value.LoggerCtxKey, log)          //attach logger to context
		ctx = context.WithValue(ctx, value.RequestIdCtxKey, requestId) //add X-Request-ID to context

		//read request body for logging
		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Error("error reading request body", zap.Error(err))
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		r.Body = io.NopCloser(bytes.NewBuffer(body)) // Replace the body with a new reader after reading from the original

		//log request info
		log.Info("HTTP request info", 
			zap.String("http",r.Proto),
			zap.String("host",r.Host), 
			zap.String("method",r.Method), 
			zap.String("url",r.URL.Path), 
			zap.String("content_type", r.Header.Get("Content-Type")),
			zap.String("body", string(body)))

		//attach context it to request
		next.ServeHTTP(w, r.WithContext(ctx)) //call next handler
		//...

		//post handle...

		log.Info("end processing request", zap.String("elapse_time", time.Since(start).String()))
	})
}
