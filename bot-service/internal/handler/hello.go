package handler

import (
	"fmt"
	"log"
	"net/http"
)

func (h *HttpHandler) Hello() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hello!")
		_, err := w.Write([]byte("My hello page!"))
		if err != nil {
			log.Fatal("write failed:", err)
		}
	})
}
