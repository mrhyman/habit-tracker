package domain

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"main/utils"
	"testing"
	"time"
)

func TestCreateUser(t *testing.T) {
	t.Parallel()

	var (
		validUserId        = uuid.New()
		validNickname      = "Some_userName"
		validBirthday      = testNowUtc.AddDate(-10, 0, 0)
		validActiveHabitId = uuid.New()
		uuidGenerator      = utils.FakeUUIDGenerator{FixedUUID: "82a89365-2a1e-4ed9-a373-057a7694d458"}
	)

	type args struct {
		Id            uuid.UUID
		Nickname      string
		CreatedAt     time.Time
		Birthday      *time.Time
		ActiveHabitId *uuid.UUID
	}
	tests := []struct {
		name    string
		args    args
		want    *User
		wantErr string
	}{
		{
			name: "success_all_args",
			args: args{
				Id:            validUserId,
				Nickname:      validNickname,
				Birthday:      &validBirthday,
				ActiveHabitId: &validActiveHabitId,
			},
			want: &User{
				Id:            validUserId,
				Nickname:      validNickname,
				CreatedAt:     testNowUtc,
				Birthday:      &validBirthday,
				ActiveHabitId: &validActiveHabitId,
				AggregateRoot: AggregateRoot{events: []Event{NewUserCreatedEvent(
					uuidGenerator.NewString(),
					testNowUtc,
					validUserId,
					validNickname,
					testNowUtc,
					&validBirthday,
					&validActiveHabitId,
				)}},
			},
			wantErr: "",
		},
		{
			name: "success_a_few_args",
			args: args{
				Id:            validUserId,
				Nickname:      validNickname,
				Birthday:      nil,
				ActiveHabitId: nil,
			},
			want: &User{
				Id:            validUserId,
				Nickname:      validNickname,
				CreatedAt:     testNowUtc,
				Birthday:      nil,
				ActiveHabitId: nil,
				AggregateRoot: AggregateRoot{events: []Event{NewUserCreatedEvent(
					uuidGenerator.NewString(),
					testNowUtc,
					validUserId,
					validNickname,
					testNowUtc,
					nil,
					nil,
				)}},
			},
			wantErr: "",
		},
		{
			name: "empty_userId_error",
			args: args{
				Nickname:      validNickname,
				Birthday:      nil,
				ActiveHabitId: nil,
			},
			want:    nil,
			wantErr: ErrInvalidUserID.Error(),
		},
		{
			name: "invalid_userId_error",
			args: args{
				Nickname:      "",
				Birthday:      nil,
				ActiveHabitId: nil,
			},
			want:    nil,
			wantErr: ErrInvalidUserID.Error(),
		},
		{
			name: "empty_username_error",
			args: args{
				Id:            validUserId,
				Nickname:      "",
				Birthday:      nil,
				ActiveHabitId: nil,
			},
			want:    nil,
			wantErr: ErrInvalidUserName.Error(),
		},
	}

	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			setup()

			got, err := NewUser(uuidGenerator, tt.args.Id, tt.args.Nickname, tt.args.CreatedAt, tt.args.Birthday, tt.args.ActiveHabitId)

			if tt.wantErr == "" {
				require.NoError(t, err)
				require.Equal(t, tt.want, got)
			} else {
				require.EqualError(t, err, tt.wantErr)
				require.Nil(t, got)
			}

		})
	}
}
