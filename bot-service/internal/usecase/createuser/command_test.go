package createuser

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestNewCommand(t *testing.T) {
	t.Parallel()

	var (
		userId    = uuid.New()
		nickname  = uuid.NewString()
		timestamp = time.Now()
	)

	type args struct {
		id            string
		nickname      string
		birthday      *time.Time
		activeHabitId *uuid.UUID
	}
	tests := []struct {
		name    string
		args    args
		want    Command
		wantErr string
	}{
		{
			name: "success_all_fields",
			args: args{
				id:            userId.String(),
				nickname:      nickname,
				birthday:      &timestamp,
				activeHabitId: &userId,
			},
			want:    Command{userId, nickname, &timestamp, &userId},
			wantErr: "",
		},
		{
			name: "success_a_few_fields",
			args: args{
				id:            userId.String(),
				nickname:      nickname,
				birthday:      nil,
				activeHabitId: nil,
			},
			want:    Command{userId, nickname, nil, nil},
			wantErr: "",
		},
		{
			name: "user_id_is_nil",
			args: args{
				nickname:      nickname,
				birthday:      nil,
				activeHabitId: nil,
			},
			want:    Command{},
			wantErr: ErrInvalidUserID.Error(),
		},
		{
			name: "user_id_is_empty",
			args: args{
				id:            "",
				nickname:      nickname,
				birthday:      nil,
				activeHabitId: nil,
			},
			want:    Command{},
			wantErr: ErrInvalidUserID.Error(),
		},
		{
			name: "user_id_wrong_format",
			args: args{
				id:            "112233",
				nickname:      nickname,
				birthday:      nil,
				activeHabitId: nil,
			},
			want:    Command{},
			wantErr: ErrInvalidUserID.Error(),
		},
		{
			name: "nickname_is_empty",
			args: args{
				id:            userId.String(),
				nickname:      "",
				birthday:      nil,
				activeHabitId: nil,
			},
			want:    Command{},
			wantErr: ErrInvalidUserName.Error(),
		},
	}

	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := NewCommand(tt.args.id, tt.args.nickname, tt.args.birthday, tt.args.activeHabitId)

			require.Equal(t, tt.want, got)
			if tt.wantErr == "" {
				require.NoError(t, err)
			} else {
				require.EqualError(t, err, tt.wantErr)
			}
		})
	}
}
