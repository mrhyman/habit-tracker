package createuser

import (
	"context"
	"errors"
	"main/pkg"

	"main/internal/domain"
)

func (ch CommandHandler) Handle(ctx context.Context, cmd Command) error {
	user, err := ch.userRepo.GetUserByID(cmd.UserId)
	if err != nil && !errors.Is(err, domain.ErrUserNotFound) {
		return err
	}
	if user != nil {
		return domain.ErrUserAlreadyExists
	}

	user, err = domain.NewUser(
		pkg.RealUUIDGenerator{},
		cmd.UserId,
		cmd.UserNickname,
		cmd.UserCreatedAt,
		cmd.UserBirthday,
		cmd.UserActiveHabitIds,
	)
	if err != nil {
		return err
	}
	if err = ch.userRepo.CreateUser(user); err != nil {
		return err
	}

	if user.IsAdult() {
		adultUserInc.Inc()
	}

	return ch.eventRouter.RouteAllEvents(ctx, user.PopAllEvents())
}
