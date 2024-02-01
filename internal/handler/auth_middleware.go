package handler

import (
	"context"
	"net/http"

	"rest-api/internal/logger"
	"rest-api/internal/value"

	"go.uber.org/zap"
)

func AuthMid(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log := logger.FromCtx(r.Context())

		if value.GetApiEndpointFlagsFromCtx(r.Context()).Auth {

			//check auth...
			log.Info("check auth...")

			tkn := r.Header.Get("token")
			log.Info("token=" + tkn)

			ctx := context.WithValue(r.Context(), value.UserNameKey, tkn)
			r = r.WithContext(ctx)

			logger := logger.FromCtx(r.Context())
			*logger = *logger.With(zap.String(value.UserNameKey, tkn)) //replace old ctx logger with new child logger with username attr added
		}

		next.ServeHTTP(w, r)
	})
}
