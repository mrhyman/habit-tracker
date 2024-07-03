package domain

import "time"

var (
	timeNowFn = func() time.Time {
		return time.Now().UTC()
	}
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

// ID возвращает ID события.
func (e EventBase) ID() string {
	return e.id
}

// HappenedAt дата-время, когда событие произошло.
func (e EventBase) HappenedAt() time.Time {
	return e.happenedAt
}

// AggregateRoot базовый агрегат.
type AggregateRoot struct {
	events []Event
}

// PopAllEvents возвращает все доменные события и отчищает их.
func (a *AggregateRoot) PopAllEvents() []Event {
	res := a.events
	a.events = nil
	return res
}

func (a *AggregateRoot) addEvent(e Event) {
	a.events = append(a.events, e)
}
