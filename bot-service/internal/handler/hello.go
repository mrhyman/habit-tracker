package handler

import (
	"fmt"
	"log/slog"
	"net/http"
)

func (h *HttpHandler) Hello() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hello!")
		_, err := w.Write([]byte("My hello page!"))
		if err != nil {
			slog.ErrorContext(r.Context(), "hello-handler write error", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
}
