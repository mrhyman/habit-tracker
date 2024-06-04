package handler

import (
	"github.com/google/uuid"
	"main/internal/domain"
	"time"
)

type UserModel struct {
	Id            string     `json:"id"`
	Nickname      string     `json:"nickname"`
	CreatedAt     time.Time  `json:"createdAt"`
	Birthday      *time.Time `json:"birthday"`
	ActiveHabitId *uuid.UUID `json:"active_habit_id,omitempty"`
}

func (um UserModel) toUser() *domain.User {
	uuidId, _ := uuid.Parse(um.Id)
	return &domain.User{
		Id:            uuidId,
		Nickname:      um.Nickname,
		CreatedAt:     um.CreatedAt,
		Birthday:      um.Birthday,
		ActiveHabitId: um.ActiveHabitId,
	}
}

func UserFromDomain(user *domain.User) UserModel {
	return UserModel{
		Id:            user.Id.String(),
		Nickname:      user.Nickname,
		CreatedAt:     user.CreatedAt,
		Birthday:      user.Birthday,
		ActiveHabitId: user.ActiveHabitId,
	}
}
