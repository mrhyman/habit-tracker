package activatehabit

import (
	"errors"
	"github.com/google/uuid"
)

var (
	ErrInvalidUserID  = errors.New("user ID should be a valid UUID")
	ErrInvalidHabitID = errors.New("habit ID should be a valid UUID")
)

type Command struct {
	UserId  uuid.UUID
	HabitId uuid.UUID
}

func NewCommand(
	userId string,
	habitId string,
) (Command, error) {
	if userId == "" {
		return Command{}, ErrInvalidUserID
	}

	if habitId == "" {
		return Command{}, ErrInvalidHabitID
	}

	uuidUserId, err := uuid.Parse(userId)
	if err != nil {
		return Command{}, ErrInvalidUserID
	}

	uuidHabitId, err := uuid.Parse(habitId)
	if err != nil {
		return Command{}, ErrInvalidHabitID
	}

	return Command{
		UserId:  uuidUserId,
		HabitId: uuidHabitId,
	}, nil
}
