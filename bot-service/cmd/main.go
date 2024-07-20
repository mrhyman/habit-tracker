package main

import (
	"context"
	"main/internal/repo/eventbus/habitactivated"
	"main/internal/repo/eventbus/usercreated"
	"main/internal/usecase/activatehabit"
	"main/pkg"
	"syscall"

	"github.com/ds248a/closer"
	"main/internal/eventrouter"
	"main/internal/repo/database"

	"main/internal/config"
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

// @externalDocs.description	OpenAPI
// @externalDocs.url			https://swagger.io/resources/open-api/
func main() {
	ctx := context.Background()
	initDefaultLogger()
	cfg := config.MustLoad()
	initLogger(cfg.Logger)

	db, err := database.New(ctx, cfg.Database)
	if err != nil {
		pkg.LogFatal(ctx, "error create db pool", err)
	}
	closer.Add(db.Close)

	userCreatedEventRepo, err := usercreated.NewRepo(ctx, cfg.Kafka.Host, cfg.UserCreatedEventProducerConfig)
	if err != nil {
		pkg.LogFatal(ctx, "error create user_created producer", err)
	}
	habitActivatedEventRepo, err := habitactivated.NewRepo(ctx, cfg.Kafka.Host, cfg.HabitActivatedEventProducerConfig)
	if err != nil {
		pkg.LogFatal(ctx, "error create habit_activated producer", err)
	}

	userRepo := database.NewRepo(ctx, db.Pool)

	// Event router
	eventRouter := eventrouter.NewService(userCreatedEventRepo, habitActivatedEventRepo)

	httpHandler := handler.New(
		createuser.NewCommandHandler(userRepo, eventRouter),
		getuserbyid.NewQueryHandler(userRepo),
		activatehabit.NewCommandHandler(userRepo, eventRouter),
	)

	go metrics.RecordMetrics(userRepo, cfg.BusinessMetricsScrapeInterval)

	s := server.New(cfg.Port, *httpHandler)
	go s.Start()
	closer.Add(s.Shutdown)

	closer.ListenSignal(syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)
}
