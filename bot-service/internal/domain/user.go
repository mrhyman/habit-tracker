package domain

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

var (
	ErrInvalidUserID     = errors.New("user ID should be a valid UUID")
	ErrInvalidUserName   = errors.New("user name should not be empty")
	ErrUserAlreadyExists = errors.New("user with such id already exists")
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

	now := timeNowFn().Truncate(time.Microsecond)

	user := &User{
		Id:            userID,
		Nickname:      userName,
		CreatedAt:     now,
		Birthday:      birthday,
		ActiveHabitId: activeHabitId,
	}

	return user, nil
}

func (u *User) IsAdult() bool {
	if u.Birthday != nil {
		return time.Since(*u.Birthday).Hours()/24/365 >= 18
	}
	return false
}
