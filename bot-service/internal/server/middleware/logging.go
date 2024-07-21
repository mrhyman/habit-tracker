package middleware

import (
	"bytes"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"time"
)

type loggingErrorWriter struct {
	http.ResponseWriter
	statusCode int
	error      *bytes.Buffer
}

func (lew *loggingErrorWriter) WriteHeader(code int) {
	lew.statusCode = code
	lew.ResponseWriter.WriteHeader(code)
}

func (lew *loggingErrorWriter) Write(b []byte) (int, error) {
	lew.error.Write(b)
	return lew.ResponseWriter.Write(b)
}

func LoggingMW(next *http.ServeMux) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var bodyBytes []byte
		if r.Body != nil {
			bodyBytes, _ = io.ReadAll(r.Body)
		}

		lew := &loggingErrorWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
			error:          new(bytes.Buffer),
		}

		r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

		start := time.Now()

		next.ServeHTTP(lew, r)
		slog.InfoContext(r.Context(),
			"request received",
			"method", r.Method,
			"url", r.URL.String(),
			"body", string(bodyBytes),
			"time_taken_ms", fmt.Sprintf("%d ms", time.Since(start).Milliseconds()),
		)

		if lew.error.Len() > 0 {
			slog.ErrorContext(r.Context(),
				"request error",
				"status", lew.statusCode,
				"err", lew.error.String(),
				"time_taken_ms", fmt.Sprintf("%d ms", time.Since(start).Milliseconds()),
			)
		}
	})
}
