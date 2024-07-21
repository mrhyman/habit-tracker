package outbox

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log/slog"
	"main/internal/domain"
	"main/internal/repo/database/event"
)

func SaveEvents(ctx context.Context, tx pgx.Tx, events []domain.Event) error {
	for _, ev := range events {
		record := event.FromDomain(ev)

		query := `INSERT INTO events (id, event_type, created_at, payload, status) VALUES (@id, @eventType, @createdAt, @payload, @status) ON CONFLICT (id) DO NOTHING`
		args := pgx.NamedArgs{
			"id":        record.Id,
			"eventType": record.EventType.String(),
			"createdAt": record.CreatedAt,
			"payload":   record.Payload,
			"status":    record.Status.String(),
		}
		slog.DebugContext(ctx, "create user_created event DB record", slog.String("query_params", fmt.Sprintf("%+v\n", record)))

		_, err := tx.Exec(ctx, query, args)
		if err != nil {
			return err
		}
	}

	return nil
}
