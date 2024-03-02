package handler

import (
	"bytes"
	"context"
	"errors"
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
		log := logger.Get()
		//pre handle.....
		start := time.Now()

		//set request timeout (ctx timeout)
		ctx, cancel := context.WithTimeout(r.Context(), time.Second*60)
		defer cancel()

		//get request ID
		requestId := r.Header.Get(value.RequestIdHeaderKey)
		if requestId == "" {	//generate a request ID for the request
			requestId = xid.New().String()
		}
		w.Header().Add(value.RequestIdHeaderKey, requestId)	//add request ID header

		//get api endpoint flags (auth, perm, logs)
		endpointFlags, err := value.GetApiEndpointFlags(r)
		if err != nil {
			err := errors.New("api endpoint not implemented")
			log.Error("no matching api endpoint", zap.Error(err))
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("404 Not Found"))
			return
		}

		//create a child logger from main logger then add the addtional info of request
		log = log.With(
			zap.String(string(value.RequestIdKey), requestId),
			zap.Any("request_uri", r.RequestURI),
			zap.Any(value.ApiEndpointFlagsKey, endpointFlags),
		)

		ctx = context.WithValue(ctx, value.LoggerKey, log)          //attach logger to context
		ctx = context.WithValue(ctx, value.RequestIdKey, requestId) //add X-Request-ID to context
		ctx = context.WithValue(ctx, value.ApiEndpointFlagsKey, endpointFlags) //add Api endpoint flags to context

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
			zap.String("http", r.Proto),
			zap.String("host", r.Host),
			zap.String("method", r.Method),
			// zap.String("request_uri", r.RequestURI),
			zap.String("query_string", r.URL.RawQuery),
			zap.String("content_type", r.Header.Get("Content-Type")),
			zap.String("body", string(body)))

		//attach context to request, call next handler
		next.ServeHTTP(w, r.WithContext(ctx)) 
		//...

		//post handle...

		log.Info("end processing request", zap.String("elapse_time", time.Since(start).String()))
	})
}
