package database

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"time"
)

const defaultTimeout = 3 * time.Second

type DB struct {
	*pgxpool.Pool
}

func New(ctx context.Context, connString string) (*DB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	db, err := pgxpool.New(ctx, connString)
	if err != nil {
		log.Fatal("unable to create connection pool:", err)
		return nil, err
	}

	return &DB{db}, nil
}
