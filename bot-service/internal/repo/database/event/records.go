package event

import (
	"main/internal/domain"
	"time"
)

type Type int

const (
	UserCreated Type = iota
	HabitActivated
)

func (et Type) String() string {
	return [...]string{"UserCreated", "HabitActivated"}[et]
}

type Status int

const (
	created Status = iota
	pending
	processed
	failed
)

func (es Status) String() string {
	return [...]string{"created", "pending", "processed", "failed"}[es]
}

type Record struct {
	Id        string       `db:"id"`
	EventType Type         `db:"event_type"`
	CreatedAt time.Time    `db:"created_at"`
	Payload   domain.Event `db:"payload"`
	Status    Status       `db:"status"`
}

func FromDomain(ev domain.Event) Record {
	switch ev.(type) {
	case *domain.UserCreatedEvent:
		return Record{
			Id:        ev.ID(),
			EventType: UserCreated,
			CreatedAt: time.Now().UTC(),
			Payload:   ev,
			Status:    created,
		}
	case *domain.HabitActivatedEvent:
		return Record{
			Id:        ev.ID(),
			EventType: HabitActivated,
			CreatedAt: time.Now().UTC(),
			Payload:   ev,
			Status:    created,
		}
	default:
		return Record{}
	}
}
