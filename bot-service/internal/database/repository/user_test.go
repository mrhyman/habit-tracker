package repository

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"main/internal/database"
	"main/internal/domain"
	"os"
	"sync"
	"testing"
)

var (
	startContainer sync.Once
	pgc            *postgres.PostgresContainer
	db             *database.DB
)

func TestMain(m *testing.M) {
	startContainer.Do(func() {
		pgc, db = database.StartDbContainer("init.sh")
	})
	defer database.StopDbContainer(pgc, db)

	code := m.Run()

	os.Exit(code)
}

func TestCreateUser(t *testing.T) {
	t.Parallel()
	t.Run("Integration_Create_And_Get_User", func(t *testing.T) {
		t.Parallel()

		repo := NewUserRepository(db.Pool)
		user, _ := domain.NewUser(uuid.New(), uuid.New().String(), nil, nil)
		err := repo.CreateUser(user)
		dbRecord, err := repo.GetUserByID(user.Id)
		require.NoError(t, err)
		require.Equal(t, user, dbRecord)
	})
}
