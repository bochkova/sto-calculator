package handlers

import (
	"net/http"
	"runtime/debug"
	"time"

	log "sto-calculator/pkg/logging"
)

type responseWriterWrapper struct {
	http.ResponseWriter
	status int
}

func wrapResponseWriter(w http.ResponseWriter) *responseWriterWrapper {
	return &responseWriterWrapper{ResponseWriter: w}
}

func (rw *responseWriterWrapper) WriteHeader(statusCode int) {
	rw.status = statusCode
	rw.ResponseWriter.WriteHeader(statusCode)
}

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := log.WithFields(log.Fields{
			"request_method": r.Method,
			"request_path":   r.URL.Path,
			"request_size":   r.ContentLength,
		})
		ctx := log.CtxWithLogger(r.Context(), logger)

		begin := time.Now()
		wrapper := wrapResponseWriter(w)
		next.ServeHTTP(wrapper, r.WithContext(ctx))

		logger.WithFields(log.Fields{
			"request_duration": time.Since(begin),
			"response_status":  wrapper.status,
		}).Info("Request completed")
	})
}

func Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.GetLoggerFromCtx(r.Context()).Errorf("Panic: %v, Stacktrace: %s", err, debug.Stack())
				http.Error(w, "Internal server error", http.StatusInternalServerError)
			}
		}()

		next.ServeHTTP(w, r)
	})
}
