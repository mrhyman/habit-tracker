package server

import (
	"context"
	"fmt"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/swaggo/http-swagger/v2"
	"log/slog"
	_ "main/docs"
	"main/internal/handler"
	"main/internal/server/middleware"
	"net/http"
	"os"
)

type Server struct {
	Instance *http.Server
	Ctx      context.Context
}

func New(port int, h handler.HttpHandler) *Server {
	return &Server{
		Ctx: context.Background(),
		Instance: &http.Server{
			Addr:    fmt.Sprintf(":%d", port),
			Handler: SetupMux(&h),
		},
	}
}

func (s *Server) Start() {
	slog.InfoContext(s.Ctx, fmt.Sprintf("Listening on port%s", s.Instance.Addr))
	if err := s.Instance.ListenAndServe(); err != nil {
		slog.ErrorContext(s.Ctx, "server start error", slog.String("err", err.Error()))
		os.Exit(1)
	}
}

func (s *Server) Shutdown() {
	slog.InfoContext(s.Ctx, "Shutting service down")
	if err := s.Instance.Shutdown(s.Ctx); err != nil {
		slog.ErrorContext(s.Ctx, "server shutdown error", slog.String("err", err.Error()))
		os.Exit(1)
	}
}

func SetupMux(h *handler.HttpHandler) http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/ping", h.Ping())
	mux.Handle("/metrics", promhttp.Handler())
	mux.Handle("POST /createUser", h.CreateUser())
	mux.Handle("GET /getUser", h.GetUserById())
	mux.Handle("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"),
	))

	return middleware.LoggingMW(mux)
}
