package repository

import (
	"github.com/google/uuid"
	"time"
)

type UserRecord struct {
	Id            uuid.UUID  `pg:"id"`
	Nickname      string     `pg:"nickname"`
	CreatedAt     time.Time  `pg:"created_at"`
	Birthday      *time.Time `pg:"birthday"`
	ActiveHabitId *uuid.UUID `pg:"active_habit_id,omitempty"`
}

func (ur UserRecord) ToDTO() UserDTO {
	return UserDTO{
		Id:            ur.Id,
		Nickname:      ur.Nickname,
		CreatedAt:     ur.CreatedAt,
		Birthday:      ur.Birthday,
		ActiveHabitId: ur.ActiveHabitId,
	}
}

func (ur UserRecord) FromDTO(dto UserDTO) UserRecord {
	return UserRecord{
		Id:            dto.Id,
		Nickname:      dto.Nickname,
		CreatedAt:     dto.CreatedAt,
		Birthday:      dto.Birthday,
		ActiveHabitId: dto.ActiveHabitId,
	}
}
