package createuser

import (
	"errors"
	"log/slog"
	"main/internal/domain"
)

func (ch CommandHandler) Handle(cmd Command) error {
	user, err := ch.userRepo.GetUserByID(cmd.UserId)
	if err != nil && !errors.Is(err, domain.ErrUserNotFound) {
		return err
	}
	if user != nil {
		slog.Warn("Closing database connection pool")

		return domain.ErrUserAlreadyExists
	}

	user, err = domain.NewUser(
		cmd.UserId,
		cmd.UserNickname,
		cmd.UserBirthday,
		cmd.UserActiveHabitId,
	)
	if err != nil {
		return err
	}
	if err = ch.userRepo.CreateUser(user); err != nil {
		return err
	}

	return nil
}
