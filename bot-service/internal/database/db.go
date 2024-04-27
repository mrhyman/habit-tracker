package database

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"main/internal/config"
)

type DB struct {
	*pgxpool.Pool
}

func New(ctx context.Context, config config.DatabaseConfig) (*DB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), config.GetTimeout())
	defer cancel()

	db, err := pgxpool.New(ctx, config.GetConnection())
	if err != nil {
		log.Fatal("unable to create connection pool:", err)
		return nil, err
	}

	return &DB{db}, nil
}

func (db *DB) Close() {
	log.Println("Closing database connection pool")
	db.Pool.Close()
}
