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
	//uuid.UUID -> string? валидация внутри хендла команды
	UserId            uuid.UUID
	UserNickname      string
	UserCreatedAt     time.Time
	UserBirthday      *time.Time
	UserActiveHabitId *uuid.UUID
}

func NewCommand(
	id uuid.UUID,
	nickname string,
	createdAt time.Time,
	birthday *time.Time,
	activeHabitId *uuid.UUID,
) (Command, error) {
	if uuid.Nil == id {
		return Command{}, ErrInvalidUserID
	}

	if nickname == "" {
		return Command{}, ErrInvalidUserName
	}

	return Command{
		UserId:            id,
		UserNickname:      nickname,
		UserCreatedAt:     createdAt,
		UserBirthday:      birthday,
		UserActiveHabitId: activeHabitId,
	}, nil
}
