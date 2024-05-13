package getuserbyid

import (
	"context"
	"main/internal/domain"
)

func (qh QueryHandler) Handle(ctx context.Context, q Query) (*domain.User, error) {
	user, err := qh.userRepo.GetUserByID(ctx, q.UserID)

	if err != nil {
		return nil, err
	}

	return user, nil
}
