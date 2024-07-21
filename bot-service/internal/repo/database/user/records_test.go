package user

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"main/internal/domain"
	"testing"
	"time"
)

func TestRecordFromDomain(t *testing.T) {
	t.Parallel()

	t.Run("Record_from_domain", func(t *testing.T) {
		t.Parallel()

		user := domain.User{Id: uuid.New(), Nickname: uuid.New().String(), CreatedAt: time.Now().UTC()}

		record := userFromDomain(&user)
		require.Equal(t, record, Record{Id: user.Id, Nickname: user.Nickname, CreatedAt: user.CreatedAt})
	})
}
