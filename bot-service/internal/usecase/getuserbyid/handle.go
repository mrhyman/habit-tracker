package getuserbyid

import (
	"main/internal/domain"
)

func (qh QueryHandler) Handle(q Query) (*domain.User, error) {
	user, err := qh.userRepo.GetUserByID(q.UserID)

	if err != nil {
		return nil, err
	}

	return user, nil
}
