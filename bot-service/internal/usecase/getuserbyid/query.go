package getuserbyid

import (
	"errors"
	"github.com/google/uuid"
)

var (
	ErrInvalidUserID = errors.New("user ID should be a valid UUID")
)

type Query struct {
	UserID uuid.UUID
}

func NewQuery(userID uuid.UUID) (Query, error) {
	if uuid.Nil == userID {
		return Query{}, ErrInvalidUserID
	}

	return Query{
		UserID: userID,
	}, nil
}
