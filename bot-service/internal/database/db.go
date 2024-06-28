package database

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"log/slog"
	"main/internal/config"
)

type DB struct {
	ctx context.Context
	*pgxpool.Pool
}

func New(ctx context.Context, config config.DatabaseConfig) (*DB, error) {
	ctx, cancel := context.WithTimeout(ctx, config.GetTimeout())
	defer cancel()

	db, err := pgxpool.New(ctx, config.GetConnection())
	if err != nil {
		return nil, err
	}

	if err = db.Ping(ctx); err != nil {
		return nil, err
	}

	return &DB{ctx, db}, nil
}

func (db *DB) Close() {
	slog.InfoContext(db.ctx, "Closing database connection pool")
	db.Pool.Close()
}
