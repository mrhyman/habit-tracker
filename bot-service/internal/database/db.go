package database

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"log/slog"
	"main/internal/config"
	"os"
)

type DB struct {
	*pgxpool.Pool
}

func New(config config.DatabaseConfig) (*DB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), config.GetTimeout())
	defer cancel()

	db, err := pgxpool.New(ctx, config.GetConnection())
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	return &DB{db}, nil
}

func (db *DB) Close() {
	slog.Info("Closing database connection pool")
	db.Pool.Close()
}
