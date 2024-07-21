package createuser

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

var (
	ErrInvalidUserID   = errors.New("user ID should be a valid UUID")
	ErrInvalidUserName = errors.New("user name should not be empty")
)

type Command struct {
	UserId             uuid.UUID
	UserNickname       string
	UserCreatedAt      time.Time
	UserBirthday       *time.Time
	UserActiveHabitIds []uuid.UUID
}

func NewCommand(
	id string,
	nickname string,
	birthday *time.Time,
	activeHabitIds []uuid.UUID,
) (Command, error) {
	if id == "" {
		return Command{}, ErrInvalidUserID
	}

	uuidId, err := uuid.Parse(id)
	if err != nil {
		return Command{}, ErrInvalidUserID
	}

	if nickname == "" {
		return Command{}, ErrInvalidUserName
	}

	return Command{
		UserId:             uuidId,
		UserNickname:       nickname,
		UserBirthday:       birthday,
		UserActiveHabitIds: activeHabitIds,
	}, nil
}
