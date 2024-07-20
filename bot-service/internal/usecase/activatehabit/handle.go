package activatehabit

import (
	"context"
)

func (ch CommandHandler) Handle(ctx context.Context, cmd Command) error {
	user, err := ch.userRepo.GetUserByID(ctx, cmd.UserId)
	if err != nil {
		return err
	}

	u, err := user.ActivateHabit(cmd.HabitId)
	if err != nil {
		return err
	}
	if err = ch.userRepo.ActivateHabit(ctx, cmd.UserId, cmd.HabitId); err != nil {
		return err
	}
	events := u.PopAllEvents()

	for _, event := range events {
		if err = ch.eventRepo.CreateEvent(ctx, &event); err != nil {
			return err
		}
	}

	return ch.eventRouter.RouteAllEvents(ctx, events)
}
