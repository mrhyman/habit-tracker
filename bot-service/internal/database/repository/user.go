package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"main/internal/domain"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository interface {
	CreateUser(ctx context.Context, order UserRecord) error
	GetUserByID(ctx context.Context, id string) (*UserRecord, error)
	SetBirthday(ctx context.Context, id string, birthday time.Time) error
	ActivateHabit(ctx context.Context, id string, habitId string) error
}

func NewUserRepository(conn *pgxpool.Pool) *UserRepositoryImpl {
	return &UserRepositoryImpl{conn: conn}
}

type UserRepositoryImpl struct {
	conn *pgxpool.Pool
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

	_, err := r.conn.Exec(context.Background(), query, args)
	return err
}

func (r *UserRepositoryImpl) GetUserByID(id uuid.UUID) (*domain.User, error) {
	var record UserRecord
	query := `SELECT * FROM users WHERE id = @id`
	args := pgx.NamedArgs{
		"id": id,
	}
	err := r.conn.QueryRow(context.Background(), query, args).Scan(
		&record.Id,
		&record.Nickname,
		&record.CreatedAt,
		&record.Birthday,
		&record.ActiveHabitId,
	)

	if err != nil {
		return nil, err
	}

	return record.toUser(), nil
}

func (r *UserRepositoryImpl) SetBirthday(ctx context.Context, id string, b time.Time) error {
	query := `UPDATE users SET birthday = $1 WHERE id = $2`
	_, err := r.conn.Exec(ctx, query, b, id)
	return err
}

func (r *UserRepositoryImpl) ActivateHabit(ctx context.Context, id string, h string) error {
	query := `UPDATE users SET active_habit_id = $1 WHERE id = $2`
	_, err := r.conn.Exec(ctx, query, h, id)
	return err
}
