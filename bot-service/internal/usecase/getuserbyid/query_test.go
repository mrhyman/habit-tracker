package getuserbyid

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewQuery(t *testing.T) {
	t.Parallel()

	var (
		validUserId = uuid.New()
	)

	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		want    Query
		wantErr string
	}{
		{
			name: "valid_user_id",
			args: args{
				id: validUserId.String(),
			},
			want:    Query{validUserId},
			wantErr: "",
		},
		{
			name: "empty_user_id",
			args: args{
				id: "",
			},
			want:    Query{},
			wantErr: ErrEmptyUserID.Error(),
		},
		{
			name: "invalid_user_id",
			args: args{
				id: "2222",
			},
			want:    Query{},
			wantErr: ErrInvalidUserID.Error(),
		},
	}

	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := NewQuery(tt.args.id)

			require.Equal(t, tt.want, got)
			if tt.wantErr == "" {
				require.NoError(t, err)
			} else {
				require.EqualError(t, err, tt.wantErr)
			}
		})
	}
}
