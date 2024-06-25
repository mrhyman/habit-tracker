package main

import (
	"context"
	"log/slog"
	"os"
	"syscall"

	"github.com/ds248a/closer"

	"main/internal/config"
	"main/internal/database"
	"main/internal/database/repository"
	"main/internal/handler"
	"main/internal/server"
	"main/internal/usecase/createuser"
	"main/internal/usecase/getuserbyid"
	"main/metrics"
)

//	@title			Habit Tracker Bot
//	@version		1.0
//	@termsOfService	http://swagger.io/terms/

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host	localhost:8080

//	@externalDocs.description	OpenAPI
//	@externalDocs.url			https://swagger.io/resources/open-api/
func main() {
	ctx := context.Background()
	initDefaultLogger()
	cfg := config.MustLoad()
	initLogger(cfg.Logger)

	db, err := database.New(ctx, cfg.Database)
	if err != nil {
		slog.ErrorContext(ctx, "error create db pool", err)
		os.Exit(1)
	}
	closer.Add(db.Close)

	userRepo := repository.NewUserRepository(ctx, db.Pool)

	httpHandler := handler.New(
		createuser.NewCommandHandler(userRepo),
		getuserbyid.NewQueryHandler(userRepo),
	)

	go metrics.RecordMetrics(userRepo, cfg.BusinessMetricsScrapeInterval)

	s := server.New(cfg.Port, *httpHandler)
	go s.Start()
	closer.Add(s.Shutdown)

	closer.ListenSignal(syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)
}
