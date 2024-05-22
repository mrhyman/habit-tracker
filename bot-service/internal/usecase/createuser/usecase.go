package createuser

import (
	"github.com/google/uuid"
	"main/internal/domain"
)

type iUserRepo interface {
	CreateUser(user *domain.User) error
	GetUserByID(userID uuid.UUID) (*domain.User, error)
}

type CommandHandler struct {
	userRepo iUserRepo
}

func NewCommandHandler(userRepo iUserRepo) *CommandHandler {
	return &CommandHandler{userRepo: userRepo}
}
