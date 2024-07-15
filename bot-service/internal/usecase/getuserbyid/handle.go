package getuserbyid

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"main/internal/domain"
)

func (qh QueryHandler) Handle(ctx context.Context, q Query) (*domain.User, error) {
	user, err := qh.userRepo.GetUserByID(q.UserID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domain.ErrUserNotFound
		}
		return nil, err
	}

	return user, nil
}
