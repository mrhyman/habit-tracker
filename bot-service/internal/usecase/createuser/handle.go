package createuser

import (
	"context"
	"errors"
	"main/internal/domain"
)

func (ch CommandHandler) Handle(ctx context.Context, cmd Command) error {
	user, err := ch.userRepo.GetUserByID(ctx, cmd.User.Id)
	if err != nil && !errors.Is(err, domain.ErrUserNotFound) {
		return err
	}

	user, err = domain.NewUser(
		cmd.User.Id,
		cmd.User.Nickname,
		cmd.User.Birthday,
		cmd.User.ActiveHabitId,
	)
	if err != nil {
		return err
	}
	if err = ch.userRepo.CreateUser(ctx, user); err != nil {
		return err
	}

	return nil
}
