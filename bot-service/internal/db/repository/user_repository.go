package repository

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository interface {
	Create(ctx context.Context, order *UserModel) error
	GetByID(ctx context.Context, id string) (*UserModel, error)
	SetBirthday(ctx context.Context, id string, birthday time.Time) error
}

func NewUserRepository(conn *pgxpool.Pool) *UserRepositoryImpl {
	return &UserRepositoryImpl{conn: conn}
}

type UserRepositoryImpl struct {
	conn *pgxpool.Pool
}

func (r *UserRepositoryImpl) Create(ctx context.Context, order *UserModel) error {
	query := `INSERT INTO users 
    	(id, nickname, created_at, birthday, active_habit_id) VALUES ($1, $2, $3, $4, $5)
    	ON CONFLICT (id) DO NOTHING`
	args := make([]interface{}, 5)
	args[0] = order.Id
	args[1] = order.Nickname
	args[2] = order.CreatedAt
	args[3] = order.Birthday
	args[3] = order.ActiveHabitId

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
