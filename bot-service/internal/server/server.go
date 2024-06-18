package server

import (
	"context"
	"fmt"
	"log/slog"
	"main/internal/handler"
	"net/http"
	"os"
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
	slog.Info(fmt.Sprintf("Listening on port%s", s.Instance.Addr))
	if err := s.Instance.ListenAndServe(); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}

func (s *Server) Shutdown() {
	slog.Info("Shutting service down")
	if err := s.Instance.Shutdown(s.Ctx); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}
