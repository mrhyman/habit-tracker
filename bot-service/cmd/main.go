package main

import (
	"github.com/ds248a/closer"
	"log/slog"
	"main/internal/config"
	"main/internal/database"
	"main/internal/database/repository"
	"main/internal/handler"
	"main/internal/server"
	"main/internal/usecase/createuser"
	"main/internal/usecase/getuserbyid"
	"main/metrics"
	"os"
	"syscall"
)

func main() {
	cfg := config.MustLoad()
	config.InitLogger(cfg.Logger)

	db, err := database.New(cfg.Database)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	userRepo := repository.NewUserRepository(db.Pool)

	httpHandler := handler.New(
		createuser.NewCommandHandler(userRepo),
		getuserbyid.NewQueryHandler(userRepo),
	)

	go metrics.RecordMetrics(userRepo, cfg.BusinessMetricsScrapeInterval)

	s := server.New(cfg.Port, *httpHandler)
	go s.Start()

	closer.Add(db.Close)
	closer.Add(s.Shutdown)
	closer.ListenSignal(syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)
}
