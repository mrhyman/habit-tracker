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
	Id            uuid.UUID  `json:"userId"`
	Nickname      string     `json:"nickname"`
	CreatedAt     time.Time  `json:"createdAt"`
	Birthday      *time.Time `json:"birthday"`
	ActiveHabitId *uuid.UUID `json:"activeHabitId,omitempty"`
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
