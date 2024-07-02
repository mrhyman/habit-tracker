//go:generate minimock -g -s .go -o ../../../mocks/usecase/createUser
package createuser

import (
	"github.com/google/uuid"
	"main/internal/domain"
)

type iUserRepo interface {
	CreateUser(user *domain.User) error
	GetUserByID(userID uuid.UUID) (*domain.User, error)
}

type iUserEventBus interface {
	UserCreated(user *domain.User) error
}

type CommandHandler struct {
	userRepo  iUserRepo
	userEvent iUserEventBus
}

func NewCommandHandler(userRepo iUserRepo, userEvent iUserEventBus) *CommandHandler {
	return &CommandHandler{userRepo: userRepo, userEvent: userEvent}
}
