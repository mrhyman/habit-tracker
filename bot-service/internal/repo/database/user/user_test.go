package user

import (
	"context"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"main/internal/domain"
	"main/internal/repo/database"
	"main/pkg"
	"os"
	"sync"
	"testing"
	"time"
)

var (
	startContainer sync.Once
	pgc            *postgres.PostgresContainer
	db             *database.DB
	testNowUtc     = time.Now().Truncate(time.Microsecond).UTC()
	uuidGenerator  = pkg.FakeUUIDGenerator{FixedUUID: uuid.NewString()}
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
		ctx := context.Background()
		repo := NewRepo(db.Pool)
		user, _ := domain.NewUser(uuidGenerator, uuid.New(), uuid.New().String(), testNowUtc, nil, nil)
		err := repo.CreateUser(ctx, user)
		dbRecord, err := repo.GetUserByID(ctx, user.Id)
		require.NoError(t, err)
		require.Equal(t, &domain.User{
			Id:             user.Id,
			Nickname:       user.Nickname,
			CreatedAt:      user.CreatedAt,
			Birthday:       nil,
			ActiveHabitIds: nil,
			AggregateRoot:  domain.AggregateRoot{Events: nil},
		}, dbRecord)
	})
}
