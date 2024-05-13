package createuser

import (
	"errors"
	"github.com/google/uuid"
	"main/internal/domain"
)

var (
	ErrInvalidUserID   = errors.New("user ID should be a valid UUID")
	ErrInvalidUserName = errors.New("user name should not be empty")
)

type Command struct {
	User domain.User
}

func NewCommand(user domain.User) (Command, error) {
	if uuid.Nil == user.Id {
		return Command{}, ErrInvalidUserID
	}

	if user.Nickname == "" {
		return Command{}, ErrInvalidUserName
	}

	return Command{
		User: user,
	}, nil
}
