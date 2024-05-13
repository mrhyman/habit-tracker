package domain

import "errors"

var (
	ErrUserNotFound = errors.New("can't find user based on the specified parameters")
	ErrIdInvalid    = errors.New("invalid id provided")
)
