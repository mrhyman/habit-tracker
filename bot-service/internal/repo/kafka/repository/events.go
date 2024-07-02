package repository

import (
	"github.com/google/uuid"
	"main/internal/domain"
	"time"
)

type UserCreatedEvent struct {
	Id            uuid.UUID  `json:"id"`
	Nickname      string     `json:"nickname"`
	CreatedAt     time.Time  `json:"created_at"`
	Birthday      *time.Time `json:"birthday"`
	ActiveHabitId *uuid.UUID `json:"active_habit_id,omitempty"`
}

func eventFromDomain(user *domain.User) UserCreatedEvent {
	return UserCreatedEvent{
		Id:            user.Id,
		Nickname:      user.Nickname,
		CreatedAt:     user.CreatedAt,
		Birthday:      user.Birthday,
		ActiveHabitId: user.ActiveHabitId,
	}
}
