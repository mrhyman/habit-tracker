package database

import (
	"context"
	"fmt"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
	"log/slog"
	"main/internal/config"
	"os"
	"path/filepath"
	"time"
)

func StartDbContainer(fixtureName string) (*postgres.PostgresContainer, *DB) {
	ctx := context.Background()
	testDataDir, _ := os.Getwd()
	testDataPath := filepath.Join(testDataDir, "testdata", fixtureName)

	pgc, err := postgres.RunContainer(ctx,
		testcontainers.WithImage("docker.io/postgres:15-alpine"),
		postgres.WithDatabase("dbName"),
		postgres.WithUsername("dbUser"),
		postgres.WithPassword("dbPassword"),
		postgres.WithInitScripts(testDataPath),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(5*time.Second)),
	)

	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	dbConn, err := pgc.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	db, err := New(config.DatabaseConfig{Connection: dbConn})
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	err = WaitForPostgres(db)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	return pgc, db
}

func StopDbContainer(pgc *postgres.PostgresContainer, db *DB) {
	pgc.Terminate(context.Background())
	db.Close()
}

func WaitForPostgres(db *DB) error {
	for i := 0; i < 10; i++ {
		err := db.Ping(context.Background())
		if err == nil {
			return nil
		}
		time.Sleep(1 * time.Second)
	}
	return fmt.Errorf("PostgreSQL not ready")
}
