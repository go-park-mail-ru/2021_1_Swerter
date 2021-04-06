package middleware

import (
	"fmt"
	"my-motivation/internal/pkg/utils/logger"
	"net/http"
	"time"
)

func LoggingMiddleware(next http.Handler, logger *logger.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		logger.Log.Infof(fmt.Sprintf("%s %s %s %s", r.Method, r.RequestURI, r.RemoteAddr, time.Since(start)))
	})
}