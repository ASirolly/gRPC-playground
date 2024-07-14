package middleware

import (
	"fmt"
	"log/slog"
	"net/http"
)

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func newLoggingResponseWriter(w http.ResponseWriter) *loggingResponseWriter {
	return &loggingResponseWriter{w, http.StatusOK}
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

func Logging(next http.Handler, l *slog.Logger) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			lrw := newLoggingResponseWriter(w)
			next.ServeHTTP(lrw, r)

			status := lrw.statusCode
			method := r.Method
			path := r.URL.Path
			l.Info(
				fmt.Sprintf("%d %s", status, http.StatusText(status)),
				`status`, status,
				`method`, method,
				`path`, path,
			)
		})
}
