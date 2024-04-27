package main

import (
	"context"
	"github.com/ds248a/closer"
	"log"
	"main/internal/config"
	"main/internal/database"
	"main/internal/server"
	"syscall"
)

func main() {
	ctx := context.Background()
	cfg := config.MustLoad()

	db, err := database.New(ctx, cfg.Database)
	if err != nil {
		log.Fatal("unable to create connection pool:", err)
	}

	s := server.New(ctx, db)
	go s.Start()

	closer.Add(db.Close)
	closer.Add(s.Shutdown)
	closer.ListenSignal(syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)
}
