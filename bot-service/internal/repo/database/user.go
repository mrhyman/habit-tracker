package database

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"

	"main/internal/domain"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewRepo(ctx context.Context, conn *pgxpool.Pool) *UserRepo {
	return &UserRepo{ctx: ctx, conn: conn}
}

type UserRepo struct {
	conn *pgxpool.Pool
	ctx  context.Context
}

func (r *UserRepo) CreateUser(user *domain.User) error {
	record := userFromDomain(user)
	query := `INSERT INTO users (id, nickname, created_at, birthday, active_habit_ids) VALUES (@id, @nickname, @createdAt, @birthday, @activeHabitIds) ON CONFLICT (id) DO NOTHING`
	args := pgx.NamedArgs{
		"id":             record.Id,
		"nickname":       record.Nickname,
		"createdAt":      record.CreatedAt,
		"birthday":       record.Birthday,
		"activeHabitIds": record.ActiveHabitIds,
	}
	slog.DebugContext(r.ctx, "create user DB record", slog.String("query_params", fmt.Sprintf("%+v\n", record)))

	_, err := r.conn.Exec(context.Background(), query, args)
	return err
}

func (r *UserRepo) GetUserByID(id uuid.UUID) (*domain.User, error) {
	var ur UserRecord
	query := `SELECT * FROM users WHERE id = @id`
	args := pgx.NamedArgs{
		"id": id,
	}
	err := pgxscan.Get(context.Background(), r.conn, &ur, query, args)
	slog.DebugContext(r.ctx, "get user DB record", slog.String("query_params", id.String()))

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

func (r *UserRepo) SetBirthday(id string, b time.Time) error {
	query := `UPDATE users SET birthday = $1 WHERE id = $2`
	slog.DebugContext(r.ctx, "update user DB record birthday", slog.String("query_params", fmt.Sprintf("id: %s, birthday: %s", id, b)))
	_, err := r.conn.Exec(r.ctx, query, b, id)
	return err
}

func (r *UserRepo) ActivateHabit(id uuid.UUID, h uuid.UUID) error {
	query := `UPDATE users SET active_habit_ids = array_append(active_habit_ids, $1) WHERE id = $2`
	slog.DebugContext(r.ctx, "update user DB record habit", "query_params", fmt.Sprintf("id: %s, habit: %s", id, h))
	_, err := r.conn.Exec(r.ctx, query, h, id)
	return err
}
