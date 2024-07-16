package activatehabit

import (
	"context"
	"github.com/google/uuid"
	"main/internal/domain"
	"time"
)

func (ch CommandHandler) Handle(ctx context.Context, cmd Command) error {
	user, err := ch.userRepo.GetUserByID(cmd.UserId)
	if err != nil {
		return err
	}

	if err = ch.userRepo.ActivateHabit(cmd.UserId, cmd.HabitId); err != nil {
		return err
	}
	user.AddEvent(domain.NewHabitActivatedEvent(
		uuid.NewString(),
		time.Now().UTC(),
		user.Id,
		cmd.HabitId,
	))

	return ch.eventRouter.RouteAllEvents(ctx, user.PopAllEvents())
}
