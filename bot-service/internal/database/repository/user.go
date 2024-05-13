package repository

import (
	"context"
	"fmt"
	"github.com/google/uuid"
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

func (r *UserRepositoryImpl) CreateUser(ctx context.Context, user *domain.User) error {
	query := `INSERT INTO users 
   	(id, nickname, created_at, birthday, active_habit_id) VALUES ($1, $2, $3, $4, $5)
   	ON CONFLICT (id) DO NOTHING`
	args := make([]interface{}, 5)
	args[0] = user.Id
	args[1] = user.Nickname
	args[2] = user.CreatedAt
	args[3] = user.Birthday
	args[4] = user.ActiveHabitId

	u, err := r.conn.Exec(ctx, query, args...)
	fmt.Println(u)
	return err
}

func (r *UserRepositoryImpl) GetUserByID(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	var user domain.User
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
