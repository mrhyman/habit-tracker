package database

import (
	"context"
	"fmt"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
	"main/internal/config"
	"main/utils"
	"os"
	"path/filepath"
	"time"
)

func StartDbContainer(fixtureName string) (*postgres.PostgresContainer, *DB) {
	ctx := context.Background()
	timeout := 5 * time.Second
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
				WithStartupTimeout(timeout)),
	)

	if err != nil {
		utils.LogFatal(ctx, "error run container", err)
	}

	dbConn, err := pgc.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		utils.LogFatal(ctx, "error get connection string", err)
	}

	db, err := New(ctx, config.DatabaseConfig{Connection: dbConn, Timeout: timeout})
	if err != nil {
		utils.LogFatal(ctx, "error create db pool", err)
	}

	err = WaitForPostgres(db)
	if err != nil {
		utils.LogFatal(ctx, "error wait for db connection", err)
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
