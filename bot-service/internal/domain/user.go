package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrInvalidUserID     = errors.New("user ID should be a valid UUID")
	ErrInvalidUserName   = errors.New("user name should not be empty")
	ErrUserAlreadyExists = errors.New("user with such id already exists")
)

type User struct {
	AggregateRoot
	Id            uuid.UUID
	Nickname      string
	CreatedAt     time.Time
	Birthday      *time.Time
	ActiveHabitId *uuid.UUID
}

func NewUser(
	userID uuid.UUID,
	userName string,
	createdAt time.Time,
	birthday *time.Time,
	activeHabitId *uuid.UUID,
) (*User, error) {
	nowUTC := timeNowFn()
	created := createdAt.UTC()
	if uuid.Nil == userID {
		return nil, ErrInvalidUserID
	}

	if userName == "" {
		return nil, ErrInvalidUserName
	}

	if time.Time.IsZero(created) {
		created = nowUTC
	}

	user := &User{
		Id:            userID,
		Nickname:      userName,
		CreatedAt:     created.Truncate(time.Microsecond),
		Birthday:      birthday,
		ActiveHabitId: activeHabitId,
	}

	user.addEvent(NewUserCreatedEvent(
		uuid.NewString(),
		nowUTC,
		user.Id,
		user.Nickname,
		user.CreatedAt,
		user.Birthday,
		user.ActiveHabitId,
	))

	return user, nil
}

func (u *User) IsAdult() bool {
	if u.Birthday != nil {
		return time.Since(*u.Birthday).Hours()/24/365 >= 18
	}
	return false
}
