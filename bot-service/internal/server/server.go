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
	"main/pkg"
	"net/http"
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
		pkg.LogFatal(s.Ctx, "server start error", err)
	}
}

func (s *Server) Shutdown() {
	slog.InfoContext(s.Ctx, "Shutting service down")
	if err := s.Instance.Shutdown(s.Ctx); err != nil {
		pkg.LogFatal(s.Ctx, "server shutdown error", err)
	}
}

func SetupMux(h *handler.HttpHandler) http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/ping", h.Ping())
	mux.Handle("/metrics", promhttp.Handler())
	mux.Handle("POST /user/create", h.CreateUser())
	mux.Handle("GET /user/search", h.GetUserById())
	mux.Handle("PUT /user/activateHabit", h.ActivateHabit())
	mux.Handle("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"),
	))

	return middleware.LoggingMW(mux)
}
