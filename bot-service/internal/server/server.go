package server

import (
	"context"
	"fmt"
	"log"
	"main/internal/handler"
	"net/http"
)

type Server struct {
	Instance *http.Server
	Ctx      context.Context
}

func New(port int, httpHandler handler.HttpHandler) *Server {
	return &Server{
		Ctx: context.Background(),
		Instance: &http.Server{
			Addr:    fmt.Sprintf(":%d", port),
			Handler: httpHandler.SetupMux(),
		},
	}
}

func (s *Server) Start() {
	fmt.Println("Starting service at port 8080")
	if err := s.Instance.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func (s *Server) Shutdown() {
	log.Println("Shutting service down")
	if err := s.Instance.Shutdown(s.Ctx); err != nil {
		log.Fatal(err)
	}
}
