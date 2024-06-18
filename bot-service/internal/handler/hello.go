package handler

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
)

func (h *HttpHandler) Hello() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hello!")
		_, err := w.Write([]byte("My hello page!"))
		if err != nil {
			slog.Error(err.Error())
			os.Exit(1)
		}
	})
}
