package database

import (
	"context"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"main/internal/domain"
	"os"
	"sync"
	"testing"
	"time"
)

var (
	startContainer sync.Once
	pgc            *postgres.PostgresContainer
	db             *DB
	testNowUtc     = time.Now().Truncate(time.Microsecond).UTC()
)

func TestMain(m *testing.M) {
	startContainer.Do(func() {
		pgc, db = StartDbContainer("init.sh")
	})
	defer StopDbContainer(pgc, db)

	code := m.Run()

	os.Exit(code)
}

func TestCreateUser(t *testing.T) {
	t.Parallel()
	t.Run("Integration_Create_And_Get_User", func(t *testing.T) {
		t.Parallel()

		repo := NewUserRepository(context.Background(), db.Pool)
		user, _ := domain.NewUser(uuid.New(), uuid.New().String(), testNowUtc, nil, nil)
		err := repo.CreateUser(user)
		dbRecord, err := repo.GetUserByID(user.Id)
		require.NoError(t, err)
		require.Equal(t, user, dbRecord)
	})
}
