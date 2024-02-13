package repository

import (
	"time"
)

type UserModel struct {
	Id            string     `pg:"id"`
	Nickname      string     `pg:"nickname"`
	CreatedAt     *time.Time `pg:"created_at"`
	Birthday      *time.Time `pg:"birthday"`
	ActiveHabitId int64      `pg:"active_habit_id"`
}

type HabitModel struct {
	Id         string     `pg:"id"`
	CreatedAt  *time.Time `pg:"created_at"`
	Name       string     `pg:"name"`
	OwnerId    int64      `pg:"owner_id"`
	IsActive   bool       `pg:"active"`
	ScheduleId int64      `pg:"schedule_id"`
}

type ScheduleModel struct {
	Id        string     `pg:"id"`
	CreatedAt *time.Time `pg:"created_at"`
	IsActive  bool       `pg:"active"`
	Schedule  string     `pg:"schedule"`
}
