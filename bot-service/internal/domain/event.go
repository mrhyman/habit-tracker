package domain

import (
	"time"

	"github.com/google/uuid"
)

// Event контракт доменного события.
type Event interface {
	ID() string
	HappenedAt() time.Time
}

// EventBase базовое доменное события.
type EventBase struct {
	id         string
	happenedAt time.Time
}

func (e EventBase) ID() string {
	return e.id
}

func (e EventBase) HappenedAt() time.Time {
	return e.happenedAt
}

// AggregateRoot базовый агрегат.
type AggregateRoot struct {
	Events []Event
}

// PopAllEvents возвращает все доменные события и очищает их.
func (a *AggregateRoot) PopAllEvents() []Event {
	res := a.Events
	a.Events = nil
	return res
}

func (a *AggregateRoot) AddEvent(e Event) {
	a.Events = append(a.Events, e)
}

type UserCreatedEvent struct {
	EventBase
	UserID         uuid.UUID
	Nickname       string
	CreatedAt      time.Time
	Birthday       *time.Time
	ActiveHabitIds []uuid.UUID
}

type UserUpdatedEvent struct {
	EventBase
	UserID uuid.UUID
}

type HabitActivatedEvent struct {
	EventBase
	UserID  uuid.UUID
	HabitId uuid.UUID
}

func NewUserCreatedEvent(
	eventID string,
	now time.Time,
	userID uuid.UUID,
	nickname string,
	createdAt time.Time,
	birthday *time.Time,
	activeHabitIds []uuid.UUID,
) *UserCreatedEvent {
	return &UserCreatedEvent{
		EventBase: EventBase{
			id:         eventID,
			happenedAt: now,
		},
		UserID:         userID,
		Nickname:       nickname,
		CreatedAt:      createdAt,
		Birthday:       birthday,
		ActiveHabitIds: activeHabitIds,
	}
}

func NewHabitActivatedEvent(
	eventID string,
	now time.Time,
	userID uuid.UUID,
	habitId uuid.UUID,
) *HabitActivatedEvent {
	return &HabitActivatedEvent{
		EventBase: EventBase{
			id:         eventID,
			happenedAt: now,
		},
		UserID:  userID,
		HabitId: habitId,
	}
}
