package createuser

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestNewCommand(t *testing.T) {
	var (
		userId    = uuid.New()
		nickname  = uuid.NewString()
		timestamp = time.Now()
	)

	type args struct {
		id            uuid.UUID
		nickname      string
		createdAt     time.Time
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
			name: "valid_user_id",
			args: args{
				id:            userId,
				nickname:      nickname,
				createdAt:     timestamp,
				birthday:      nil,
				activeHabitId: nil,
			},
			want:    Command{userId, nickname, timestamp, nil, nil},
			wantErr: "",
		},
	}

	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := NewCommand(tt.args.id, tt.args.nickname, tt.args.createdAt, tt.args.birthday, tt.args.activeHabitId)

			require.Equal(t, tt.want, got)
			if tt.wantErr == "" {
				require.NoError(t, err)
			} else {
				require.EqualError(t, err, tt.wantErr)
			}
		})
	}
}
