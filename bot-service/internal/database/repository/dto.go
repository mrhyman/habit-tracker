package repository

import (
	"github.com/google/uuid"
	"time"
)

type UserDTO struct {
	Id            uuid.UUID  `json:"userId"`
	Nickname      string     `json:"nickname"`
	CreatedAt     time.Time  `json:"createdAt"`
	Birthday      *time.Time `json:"birthday"`
	ActiveHabitId *uuid.UUID `json:"activeHabitId,omitempty"`
}
