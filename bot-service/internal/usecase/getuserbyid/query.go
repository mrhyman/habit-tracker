package getuserbyid

import (
	"errors"
	"github.com/google/uuid"
)

var (
	ErrEmptyUserID   = errors.New("empty user id")
	ErrInvalidUserID = errors.New("user ID should be a valid UUID")
)

type Query struct {
	UserID uuid.UUID
}

func NewQuery(userId string) (Query, error) {
	if userId == "" {
		return Query{}, ErrEmptyUserID
	}

	userUuid, err := uuid.Parse(userId)
	if err != nil {
		return Query{}, ErrInvalidUserID
	}

	return Query{
		UserID: userUuid,
	}, nil
}
