package main

import (
	"context"
	"github.com/ds248a/closer"
	"log"
	"main/internal/config"
	"main/internal/database"
	"main/internal/database/repository"
	"main/internal/handler"
	"main/internal/server"
	"main/internal/usecase/createuser"
	"main/internal/usecase/getuserbyid"
	"main/metrics"
	"syscall"
)

func main() {
	ctx := context.Background()
	cfg := config.MustLoad()

	db, err := database.New(ctx, cfg.Database)
	if err != nil {
		log.Fatal("unable to create connection pool:", err)
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
