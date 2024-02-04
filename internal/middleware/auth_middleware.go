package middleware

import (
	"context"
	"net/http"

	"rest-api/internal/logger"
	tkn "rest-api/internal/token"
	"rest-api/internal/value"

	"go.uber.org/zap"
)

func AuthMid(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log := logger.FromCtx(r.Context())

		if value.GetEndpointInfoFromCtx(r.Context()).FunctionCode != "" {
			//check auth...
			log.Info("check auth...")

			tknStr := r.Header.Get("Authorization")
			log.Info("token=" + tknStr)

			pl, _ := tkn.Decode(tknStr)

			log.Info(pl.Userlogin)

			ctx := context.WithValue(r.Context(), value.UserNameKey, pl.Userlogin)
			r = r.WithContext(ctx)

			//replace old ctx logger with new child logger with username attr added
			logger := logger.FromCtx(r.Context())
			*logger = *logger.With(zap.String(value.UserNameKey, pl.Userlogin)) 
		}

		next.ServeHTTP(w, r)
	})
}
