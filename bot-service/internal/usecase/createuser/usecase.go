package createuser

import (
	"context"
	"github.com/google/uuid"
	"main/internal/domain"
)

type iUserRepo interface {
	CreateUser(ctx context.Context, user *domain.User) error
	GetUserByID(ctx context.Context, userID uuid.UUID) (*domain.User, error)
}

type CommandHandler struct {
	userRepo iUserRepo
}

func NewCommandHandler(userRepo iUserRepo) *CommandHandler {
	return &CommandHandler{userRepo: userRepo}
}
