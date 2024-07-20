package user

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"main/internal/repo/database/outbox"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"

	"main/internal/domain"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewRepo(conn *pgxpool.Pool) *Repo {
	return &Repo{conn: conn}
}

type Repo struct {
	conn *pgxpool.Pool
}

func (r *Repo) CreateUser(ctx context.Context, user *domain.User) error {
	record := userFromDomain(user)
	query := `INSERT INTO users (id, nickname, created_at, birthday, active_habit_ids) VALUES (@id, @nickname, @createdAt, @birthday, @activeHabitIds) ON CONFLICT (id) DO NOTHING`
	args := pgx.NamedArgs{
		"id":             record.Id,
		"nickname":       record.Nickname,
		"createdAt":      record.CreatedAt,
		"birthday":       record.Birthday,
		"activeHabitIds": record.ActiveHabitIds,
	}

	tx, err := r.conn.BeginTx(ctx, pgx.TxOptions{
		IsoLevel: pgx.ReadCommitted,
	})
	defer func() {
		if err != nil {
			if errRollback := tx.Rollback(ctx); errRollback != nil {
			}
			return
		}

		if errCommit := tx.Commit(ctx); errCommit != nil {
		}
	}()

	slog.DebugContext(ctx, "create user DB record", slog.String("query_params", fmt.Sprintf("%+v\n", record)))
	_, err = tx.Exec(ctx, query, args)

	err = outbox.SaveEvents(ctx, tx, user.PopAllEvents())

	return err
}

func (r *Repo) GetUserByID(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	var ur Record
	query := `SELECT * FROM users WHERE id = @id`
	args := pgx.NamedArgs{
		"id": id,
	}
	err := pgxscan.Get(ctx, r.conn, &ur, query, args)
	slog.DebugContext(ctx, "get user DB record", slog.String("query_params", id.String()))

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domain.ErrUserNotFound
		}
		return nil, err
	}

	return &domain.User{
		Id:             ur.Id,
		Nickname:       ur.Nickname,
		CreatedAt:      ur.CreatedAt,
		Birthday:       ur.Birthday,
		ActiveHabitIds: ur.ActiveHabitIds,
	}, nil
}

func (r *Repo) SetBirthday(ctx context.Context, id string, b time.Time) error {
	query := `UPDATE users SET birthday = $1 WHERE id = $2`
	slog.DebugContext(ctx, "update user DB record birthday", slog.String("query_params", fmt.Sprintf("id: %s, birthday: %s", id, b)))
	_, err := r.conn.Exec(ctx, query, b, id)

	return err
}

func (r *Repo) ActivateHabit(ctx context.Context, id uuid.UUID, h uuid.UUID) error {
	query := `UPDATE users SET active_habit_ids = array_append(active_habit_ids, $1) WHERE id = $2`
	slog.DebugContext(ctx, "update user DB record habit", "query_params", fmt.Sprintf("id: %s, habit: %s", id, h))
	_, err := r.conn.Exec(ctx, query, h, id)

	return err
}
