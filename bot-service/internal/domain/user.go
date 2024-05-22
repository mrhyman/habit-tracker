package domain

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

var (
	ErrInvalidUserID   = errors.New("user ID should be a valid UUID")
	ErrInvalidUserName = errors.New("user name should not be empty")
)

type User struct {
	Id            uuid.UUID
	Nickname      string
	CreatedAt     time.Time
	Birthday      *time.Time
	ActiveHabitId *uuid.UUID
}

func NewUser(
	userID uuid.UUID, userName string, birthday *time.Time, activeHabitId *uuid.UUID,
) (*User, error) {
	if uuid.Nil == userID {
		return nil, ErrInvalidUserID
	}

	if userName == "" {
		return nil, ErrInvalidUserName
	}

	user := &User{
		Id:            userID,
		Nickname:      userName,
		CreatedAt:     time.Now().UTC(),
		Birthday:      birthday,
		ActiveHabitId: activeHabitId,
	}

	return user, nil
}
