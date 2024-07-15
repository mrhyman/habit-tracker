package database

import (
	"github.com/google/uuid"
	"main/internal/domain"
	"time"
)

type UserRecord struct {
	Id            uuid.UUID  `db:"id"`
	Nickname      string     `db:"nickname"`
	CreatedAt     time.Time  `db:"created_at"`
	Birthday      *time.Time `db:"birthday"`
	ActiveHabitId *uuid.UUID `db:"active_habit_id,omitempty"`
}

func userFromDomain(user *domain.User) UserRecord {
	return UserRecord{
		Id:            user.Id,
		Nickname:      user.Nickname,
		CreatedAt:     user.CreatedAt,
		Birthday:      user.Birthday,
		ActiveHabitId: user.ActiveHabitId,
	}
}
