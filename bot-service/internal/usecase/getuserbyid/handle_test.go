package getuserbyid

import (
	"errors"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/require"
	"main/internal/domain"
	"main/mocks/usecase/getuserbyid"
	"testing"
	"time"
)

type mocks struct {
	userRepoMock *getuserbyid.IUserRepoMock
}

func setup(t *testing.T) (*QueryHandler, mocks) {
	userRepoMock := getuserbyid.NewIUserRepoMock(t)

	sut := NewQueryHandler(userRepoMock)

	return sut, mocks{
		userRepoMock: userRepoMock,
	}
}
func TestUC_Handle(t *testing.T) {
	t.Parallel()

	var (
		validUuid  = uuid.New()
		validQuery = Query{UserID: validUuid}
		validUser  = domain.User{Id: validUuid, Nickname: "someName", CreatedAt: time.Now().UTC()}
		someErr    = errors.New("some error")
	)

	type args struct {
		query Query
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
				query: validQuery,
			},
			setMocks: func(m mocks) {
				m.userRepoMock.GetUserByIDMock.Set(func(userID uuid.UUID) (u *domain.User, err error) {
					return &validUser, nil
				})
			},
			wantErr: "",
		},
		{
			name: "not found",
			args: args{
				query: validQuery,
			},
			setMocks: func(m mocks) {
				m.userRepoMock.GetUserByIDMock.Set(func(userID uuid.UUID) (u *domain.User, err error) {
					return nil, pgx.ErrNoRows
				})
			},
			wantErr: domain.ErrUserNotFound.Error(),
		},
		{
			name: "some pgx error",
			args: args{
				query: validQuery,
			},
			setMocks: func(m mocks) {
				m.userRepoMock.GetUserByIDMock.Set(func(userID uuid.UUID) (u *domain.User, err error) {
					return nil, someErr
				})
			},
			wantErr: someErr.Error(),
		},
	}

	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			sut, sutMocks := setup(t)
			tt.setMocks(sutMocks)

			user, err := sut.Handle(nil, tt.args.query)

			if tt.wantErr == "" {
				require.NoError(t, err)
				require.Equal(t, &validUser, user)
			} else {
				require.EqualError(t, err, tt.wantErr)
			}
		})
	}

}
