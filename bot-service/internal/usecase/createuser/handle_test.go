package createuser

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"main/internal/domain"
	createUser "main/mocks/usecase/createuser"
	"testing"
	"time"
)

type mocks struct {
	userRepoMock    *createUser.IUserRepoMock
	eventRouterMock *createUser.IEventRouterMock
}

func setup(t *testing.T) (*CommandHandler, mocks) {
	userRepoMock := createUser.NewIUserRepoMock(t)
	eventRouterMock := createUser.NewIEventRouterMock(t)

	sut := NewCommandHandler(userRepoMock, eventRouterMock)

	return sut, mocks{
		userRepoMock:    userRepoMock,
		eventRouterMock: eventRouterMock,
	}
}

func TestUC_Handle(t *testing.T) {
	t.Parallel()

	var (
		userId    = uuid.New()
		nickname  = uuid.NewString()
		timestamp = time.Now()
		cmd       = Command{UserId: userId, UserNickname: nickname}
		validUser = domain.User{Id: userId, Nickname: nickname, CreatedAt: timestamp}
		someError = errors.New("some error")
	)

	type args struct {
		cmd Command
	}

	tests := []struct {
		name     string
		args     args
		setMocks func(m mocks)
		wantErr  string
	}{
		{
			name: "success",
			args: args{
				cmd: cmd,
			},
			setMocks: func(m mocks) {
				m.userRepoMock.GetUserByIDMock.When(cmd.UserId).Then(nil, domain.ErrUserNotFound)
				m.userRepoMock.CreateUserMock.Set(func(user *domain.User) (err error) {
					require.Equal(t, validUser.Id, user.Id)
					m.eventRouterMock.RouteAllEventsMock.When(context.Background(), user.Events).Then(nil)
					return nil
				})
			},
			wantErr: "",
		},
		{
			name: "user create error",
			args: args{
				cmd: cmd,
			},
			setMocks: func(m mocks) {
				m.userRepoMock.GetUserByIDMock.When(cmd.UserId).Then(nil, domain.ErrUserNotFound)
				m.userRepoMock.CreateUserMock.Set(func(user *domain.User) (err error) {
					require.Equal(t, validUser.Id, user.Id)
					return someError
				})
			},
			wantErr: someError.Error(),
		},
		{
			name: "user already exists error",
			args: args{
				cmd: cmd,
			},
			setMocks: func(m mocks) {
				m.userRepoMock.GetUserByIDMock.When(cmd.UserId).Then(&validUser, nil)
			},
			wantErr: domain.ErrUserAlreadyExists.Error(),
		},
	}

	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ctx := context.Background()
			sut, sutMocks := setup(t)
			tt.setMocks(sutMocks)

			err := sut.Handle(ctx, tt.args.cmd)

			if tt.wantErr == "" {
				require.NoError(t, err)
			} else {
				require.EqualError(t, err, tt.wantErr)
			}
		})
	}
}
