package createuser

import (
	"errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"main/internal/domain"
	createUser "main/mocks/usecase/createuser"
	"testing"
	"time"
)

type mocks struct {
	userRepoMock *createUser.IUserRepoMock
}

func setup(t *testing.T) (*CommandHandler, mocks) {
	userRepoMock := createUser.NewIUserRepoMock(t)

	sut := NewCommandHandler(userRepoMock)

	return sut, mocks{
		userRepoMock: userRepoMock,
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
			name: "sucsess",
			args: args{
				cmd: cmd,
			},
			setMocks: func(m mocks) {
				m.userRepoMock.GetUserByIDMock.When(cmd.UserId).Then(&validUser, nil)
				m.userRepoMock.CreateUserMock.Set(func(user *domain.User) (err error) {
					require.Equal(t, validUser.Id, user.Id)
					return nil
				})
			},
			wantErr: "",
		},
		{
			name: "user search error",
			args: args{
				cmd: cmd,
			},
			setMocks: func(m mocks) {
				m.userRepoMock.GetUserByIDMock.When(cmd.UserId).Then(nil, someError)
			},
			wantErr: someError.Error(),
		},
		{
			name: "user create error",
			args: args{
				cmd: cmd,
			},
			setMocks: func(m mocks) {
				m.userRepoMock.GetUserByIDMock.When(cmd.UserId).Then(&validUser, nil)
				m.userRepoMock.CreateUserMock.Set(func(user *domain.User) (err error) {
					require.Equal(t, validUser.Id, user.Id)
					return someError
				})
			},
			wantErr: someError.Error(),
		},
	}

	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			sut, sutMocks := setup(t)
			tt.setMocks(sutMocks)

			err := sut.Handle(tt.args.cmd)

			if tt.wantErr == "" {
				require.NoError(t, err)
			} else {
				require.EqualError(t, err, tt.wantErr)
			}
		})
	}
}
