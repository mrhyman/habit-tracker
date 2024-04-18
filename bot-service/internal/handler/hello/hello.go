package hello_handler

import (
	"fmt"
	"log"
	"net/http"
)

type Handler struct{}

func NewHelloHandler() *Handler {
	return &Handler{}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello!")
	_, err := w.Write([]byte("My hello page!"))
	if err != nil {
		log.Fatal("write failed:", err)
	}
}
