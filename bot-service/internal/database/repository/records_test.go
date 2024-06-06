package repository

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
		require.Equal(t, record, UserRecord{Id: user.Id, Nickname: user.Nickname, CreatedAt: user.CreatedAt})
	})
}

func TestUserFromRecord(t *testing.T) {
	t.Parallel()

	t.Run("User_from_record", func(t *testing.T) {
		t.Parallel()

		record := UserRecord{Id: uuid.New(), Nickname: uuid.New().String(), CreatedAt: time.Now().UTC()}

		user := record.toUser()
		require.Equal(t, *user, domain.User{Id: record.Id, Nickname: record.Nickname, CreatedAt: record.CreatedAt})
	})
}
