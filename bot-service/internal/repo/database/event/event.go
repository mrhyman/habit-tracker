package event

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log/slog"

	"main/internal/domain"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewRepo(conn *pgxpool.Pool) *Repo {
	return &Repo{conn: conn}
}

type Repo struct {
	conn *pgxpool.Pool
}

func (r *Repo) CreateEvent(ctx context.Context, ev *domain.Event) error {
	record := FromDomain(*ev)
	query := `INSERT INTO events (id, event_type, created_at, payload, status) VALUES (@id, @eventType, @createdAt, @payload, @status) ON CONFLICT (id) DO NOTHING`
	args := pgx.NamedArgs{
		"id":        record.Id,
		"eventType": record.EventType.String(),
		"createdAt": record.CreatedAt,
		"payload":   record.Payload,
		"status":    record.Status.String(),
	}

	slog.DebugContext(ctx, "create user_created event DB record", slog.String("query_params", fmt.Sprintf("%+v\n", record)))

	_, err := r.conn.Exec(ctx, query, args)
	return err
}
