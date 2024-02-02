package handler

import (
	"net/http"

	"rest-api/internal/logger"
	"rest-api/internal/value"
)

func LogMid(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)

		log := logger.FromCtx(r.Context())

		info, isType := r.Context().Value(value.ApiEndpointInfoKey).(value.ApiEndpointInfo)
		if isType && info.UserAuditLogFlag {
			log.Info("logging USER AUDIT LOG")			
		}

	})
}
