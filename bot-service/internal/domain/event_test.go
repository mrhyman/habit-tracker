package domain

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"main/pkg"
	"testing"
	"time"
)

func TestNewUserCreatedEvent(t *testing.T) {
	t.Parallel()

	var (
		validUserId        = uuid.New()
		validNickname      = "Some_userName"
		validBirthday      = testNowUtc.AddDate(-10, 0, 0)
		validActiveHabitId = uuid.New()
		uuidGenerator      = pkg.FakeUUIDGenerator{FixedUUID: "82a89365-2a1e-4ed9-a373-057a7694d458"}
	)

	type args struct {
		eventID       string
		now           time.Time
		userID        uuid.UUID
		nickname      string
		createdAt     time.Time
		birthday      *time.Time
		activeHabitId *uuid.UUID
	}
	tests := []struct {
		name string
		args args
		want *UserCreatedEvent
	}{
		{
			name: "success_all_args",
			args: args{
				eventID:       uuidGenerator.NewString(),
				now:           testNowUtc,
				userID:        validUserId,
				nickname:      validNickname,
				createdAt:     testNowUtc,
				birthday:      &validBirthday,
				activeHabitId: &validActiveHabitId,
			},
			want: &UserCreatedEvent{
				EventBase: EventBase{
					id:         uuidGenerator.NewString(),
					happenedAt: testNowUtc,
				},
				UserID:        validUserId,
				Nickname:      validNickname,
				CreatedAt:     testNowUtc,
				Birthday:      &validBirthday,
				ActiveHabitId: &validActiveHabitId,
			},
		},
		{
			name: "success_a_few_args",
			args: args{
				eventID:   uuidGenerator.NewString(),
				now:       testNowUtc,
				userID:    validUserId,
				nickname:  validNickname,
				createdAt: testNowUtc,
			},
			want: &UserCreatedEvent{
				EventBase: EventBase{
					id:         uuidGenerator.NewString(),
					happenedAt: testNowUtc,
				},
				UserID:        validUserId,
				Nickname:      validNickname,
				CreatedAt:     testNowUtc,
				Birthday:      nil,
				ActiveHabitId: nil,
			},
		},
	}

	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			setup()

			got := NewUserCreatedEvent(
				uuidGenerator.NewString(),
				testNowUtc,
				tt.args.userID,
				tt.args.nickname,
				tt.args.createdAt,
				tt.args.birthday,
				tt.args.activeHabitId,
			)

			require.Equal(t, tt.want, got)
		})
	}
}

func TestNewHabitActivatedEvent(t *testing.T) {
	t.Parallel()

	var (
		validUserId   = uuid.New()
		validHabitId  = uuid.New()
		uuidGenerator = pkg.FakeUUIDGenerator{FixedUUID: uuid.NewString()}
	)

	type args struct {
		eventID string
		now     time.Time
		userID  uuid.UUID
		habitId uuid.UUID
	}
	tests := []struct {
		name string
		args args
		want *HabitActivatedEvent
	}{
		{
			name: "success_all_args",
			args: args{
				eventID: uuidGenerator.NewString(),
				now:     testNowUtc,
				userID:  validUserId,
				habitId: validHabitId,
			},
			want: &HabitActivatedEvent{
				EventBase: EventBase{
					id:         uuidGenerator.NewString(),
					happenedAt: testNowUtc,
				},
				UserID:  validUserId,
				HabitId: validHabitId,
			},
		},
	}

	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			setup()

			got := NewHabitActivatedEvent(
				uuidGenerator.NewString(),
				testNowUtc,
				tt.args.userID,
				tt.args.habitId,
			)

			require.Equal(t, tt.want, got)
		})
	}
}
