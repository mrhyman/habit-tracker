package getuserbyid

import (
	"github.com/google/uuid"
	"main/internal/domain"
)

type iUserRepo interface {
	GetUserByID(userID uuid.UUID) (*domain.User, error)
}

type QueryHandler struct {
	userRepo iUserRepo
}

func NewQueryHandler(userRepo iUserRepo) *QueryHandler {
	return &QueryHandler{userRepo: userRepo}
}
