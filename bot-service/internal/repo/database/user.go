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

type UserRepository interface {
	CreateUser(ctx context.Context, order UserRecord) error
	GetUserByID(ctx context.Context, id string) (*UserRecord, error)
	SetBirthday(ctx context.Context, id string, birthday time.Time) error
	ActivateHabit(ctx context.Context, id string, habitId string) error
}

func NewUserRepository(ctx context.Context, conn *pgxpool.Pool) *UserRepositoryImpl {
	return &UserRepositoryImpl{ctx: ctx, conn: conn}
}

type UserRepositoryImpl struct {
	conn *pgxpool.Pool
	ctx  context.Context
}

func (r *UserRepositoryImpl) CreateUser(user *domain.User) error {
	record := userFromDomain(user)
	query := `INSERT INTO users (id, nickname, created_at, birthday, active_habit_id) VALUES (@id, @nickname, @createdAt, @birthday, @activeHabitId) ON CONFLICT (id) DO NOTHING`
	args := pgx.NamedArgs{
		"id":            record.Id,
		"nickname":      record.Nickname,
		"createdAt":     record.CreatedAt,
		"birthday":      record.Birthday,
		"activeHabitId": record.ActiveHabitId,
	}
	slog.DebugContext(r.ctx, "create user DB record", slog.String("query_params", fmt.Sprintf("%+v\n", record)))

	_, err := r.conn.Exec(context.Background(), query, args)
	return err
}

func (r *UserRepositoryImpl) GetUserByID(id uuid.UUID) (*domain.User, error) {
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
		Id:            ur.Id,
		Nickname:      ur.Nickname,
		CreatedAt:     ur.CreatedAt,
		Birthday:      ur.Birthday,
		ActiveHabitId: ur.ActiveHabitId,
	}, nil
}

func (r *UserRepositoryImpl) SetBirthday(ctx context.Context, id string, b time.Time) error {
	query := `UPDATE users SET birthday = $1 WHERE id = $2`
	slog.DebugContext(r.ctx, "update user DB record birthday", slog.String("query_pÏarams", fmt.Sprintf("id: %s, birthday: %s", id, b)))
	_, err := r.conn.Exec(ctx, query, b, id)
	return err
}

func (r *UserRepositoryImpl) ActivateHabit(ctx context.Context, id string, h string) error {
	query := `UPDATE users SET active_habit_id = $1 WHERE id = $2`
	slog.DebugContext(r.ctx, "update user DB record habit", "query_params", fmt.Sprintf("id: %s, habit: %s", id, h))
	_, err := r.conn.Exec(ctx, query, h, id)
	return err
}
