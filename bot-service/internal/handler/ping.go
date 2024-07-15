package handler

import (
	"log/slog"
	"net/http"
)

// PingHandler godoc
//
//	@Summary	Ping
//	@Tags		handler
//	@Produce	json
//	@Router		/ping [get]
//	@Success	200 {string} Pong!
func (h *HttpHandler) Ping() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("Pong!"))
		if err != nil {
			slog.ErrorContext(r.Context(), "ping-handler write error", slog.String("err", err.Error()))
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
}
