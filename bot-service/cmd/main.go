package main

import (
	"context"
	"github.com/ds248a/closer"
	"log"
	"main/internal/bot"
	"main/internal/config"
	"main/internal/database"
	"main/internal/server"
	"syscall"
)

func main() {
	var (
		ctx = context.Background()
		cfg = config.MustLoad()
	)

	db, err := database.New(ctx, cfg.Database)
	if err != nil {
		log.Fatal("unable to create connection pool:", err)
	}

	b, err := bot.New(ctx, cfg.Bot)
	if err != nil {
		log.Fatal(err)
	}
	go b.Start()

	s := server.New(ctx, db)
	go s.Start()

	closer.Add(db.Close)
	closer.Add(s.Shutdown)
	closer.Add(b.Shutdown)
	closer.ListenSignal(syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)
}
