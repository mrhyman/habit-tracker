package repository

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type HabitRepository interface {
	Create(ctx context.Context, order *UserModel) error
	GetByID(ctx context.Context, id string) (*UserModel, error)
	SetActive(ctx context.Context, id string) error
	SetSchedule(ctx context.Context, id string, scheduleId string) error
}

func NewHabitRepository(conn *pgxpool.Pool) *UserRepositoryImpl {
	return &UserRepositoryImpl{conn: conn}
}

type HabitRepositoryImpl struct {
	conn *pgxpool.Pool
}

func (r *UserRepositoryImpl) CreateAndGetId(ctx context.Context, user *UserModel) error {
	query := `INSERT INTO users 
    	(id, nickname, created_at, birthday, active_habit_id) VALUES ($1, $2, $3, $4, $5)
    	ON CONFLICT (id) DO NOTHING`
	args := make([]interface{}, 5)
	args[0] = user.Id
	args[1] = user.Nickname
	args[2] = user.CreatedAt
	args[3] = user.Birthday
	args[4] = user.ActiveHabitId

	_, err := r.conn.Exec(ctx, query, args...)
	return err
}

func (r *UserRepositoryImpl) GetByID(ctx context.Context, id string) (*UserModel, error) {
	var user UserModel
	sql := `SELECT * FROM users WHERE id = $1`
	err := r.conn.QueryRow(ctx, sql, id).Scan(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
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
