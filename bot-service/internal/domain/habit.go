package domain

import (
	"errors"
	"github.com/google/uuid"
)

var (
	ErrInvalidHabitID = errors.New("habit ID should be a valid UUID")
)

type Habit struct {
	AggregateRoot
	Id     uuid.UUID
	UserId uuid.UUID
}

func NewHabit(
	id uuid.UUID,
	userId uuid.UUID,
) (*Habit, error) {
	if uuid.Nil == id {
		return nil, ErrInvalidHabitID
	}
	if uuid.Nil == userId {
		return nil, ErrInvalidUserID
	}

	habit := &Habit{
		Id:     id,
		UserId: userId,
	}
	return habit, nil
}
