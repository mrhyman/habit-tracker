//go:generate minimock -g -s .go -o ../../../mocks/usecase/getuserbyid
package getuserbyid

import (
	"context"
	"github.com/google/uuid"
	"main/internal/domain"
)

type iUserRepo interface {
	GetUserByID(ctx context.Context, userID uuid.UUID) (*domain.User, error)
}

type QueryHandler struct {
	userRepo iUserRepo
}

func NewQueryHandler(userRepo iUserRepo) *QueryHandler {
	return &QueryHandler{userRepo: userRepo}
}
