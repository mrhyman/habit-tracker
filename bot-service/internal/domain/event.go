package domain

import (
	"time"

	"github.com/google/uuid"
)

type UserCreatedEvent struct {
	EventBase
	UserID        uuid.UUID
	Nickname      string
	CreatedAt     time.Time
	Birthday      *time.Time
	ActiveHabitId *uuid.UUID
}

func NewUserCreatedEvent(
	eventID string,
	now time.Time,
	userID uuid.UUID,
	nickname string,
	createdAt time.Time,
	birthday *time.Time,
	activeHabitId *uuid.UUID,
) *UserCreatedEvent {
	return &UserCreatedEvent{
		EventBase: EventBase{
			id:         eventID,
			happenedAt: now,
		},
		UserID:        userID,
		Nickname:      nickname,
		CreatedAt:     createdAt,
		Birthday:      birthday,
		ActiveHabitId: activeHabitId,
	}
}
