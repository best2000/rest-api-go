package handler

import (
	"context"
	"errors"
	"net/http"

	"rest-api/internal/logger"
	"rest-api/internal/value"

	"go.uber.org/zap"
)

func EndpointInfoMid(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log := logger.FromCtx(r.Context())

		//get api endpoint flags (auth, perm, logs)
		endpointFlags, err := value.GetApiEndpointFlags(r)
		if err != nil {
			err := errors.New("api endpoint not implemented")
			log.Error("no matching api endpoint", zap.Error(err))
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("404 Not Found"))
			return
		}

		//replace old ctx logger with new child logger with username attr added
		logger := logger.FromCtx(r.Context())
		*logger = *logger.With(zap.Any(value.ApiEndpointInfoKey, endpointFlags)) 

		//add Api endpoint flags to context
		ctx := context.WithValue(r.Context(), value.ApiEndpointInfoKey, endpointFlags) 

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
