package handler

import (
	"github.com/google/uuid"
	"main/internal/domain"
	"time"
)

type UserModel struct {
	Id            uuid.UUID  `json:"id"`
	Nickname      string     `json:"nickname"`
	CreatedAt     time.Time  `json:"created_at"`
	Birthday      *time.Time `json:"birthday"`
	ActiveHabitId *uuid.UUID `json:"active_habit_id,omitempty"`
}

func (um UserModel) toUser() *domain.User {
	return &domain.User{
		Id:            um.Id,
		Nickname:      um.Nickname,
		CreatedAt:     um.CreatedAt,
		Birthday:      um.Birthday,
		ActiveHabitId: um.ActiveHabitId,
	}
}

func (um UserModel) userFromDomain(user *domain.User) UserModel {
	return UserModel{
		Id:            user.Id,
		Nickname:      user.Nickname,
		CreatedAt:     user.CreatedAt,
		Birthday:      user.Birthday,
		ActiveHabitId: user.ActiveHabitId,
	}
}
