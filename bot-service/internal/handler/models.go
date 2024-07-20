package handler

import (
	"github.com/google/uuid"
	"main/internal/domain"
	"time"
)

type UserModel struct {
	Id             string      `json:"id" format:"uuid"`
	Nickname       string      `json:"nickname"`
	CreatedAt      time.Time   `json:"createdAt" format:"datetime"`
	Birthday       *time.Time  `json:"birthday" format:"datetime"`
	ActiveHabitIds []uuid.UUID `json:"active_habit_ids" format:"uuid[]"`
}

type HabitActivationModel struct {
	UserId  string `json:"user_id" format:"uuid"`
	HabitId string `json:"active_habit_id" format:"uuid"`
}

func UserFromDomain(user *domain.User) UserModel {
	return UserModel{
		Id:             user.Id.String(),
		Nickname:       user.Nickname,
		CreatedAt:      user.CreatedAt,
		Birthday:       user.Birthday,
		ActiveHabitIds: user.ActiveHabitIds,
	}
}
