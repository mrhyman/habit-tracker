package middleware

import (
	"bytes"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"time"
)

func LoggingMW(next *http.ServeMux) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var bodyBytes []byte
		if r.Body != nil {
			bodyBytes, _ = io.ReadAll(r.Body)
		}

		r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

		start := time.Now()

		next.ServeHTTP(w, r)
		slog.InfoContext(r.Context(),
			"request received",
			"method", r.Method,
			"url", r.URL.String(),
			"body", string(bodyBytes),
			"time_taken_ms", fmt.Sprintf("%d ms", time.Since(start).Milliseconds()),
		)
	})
}
