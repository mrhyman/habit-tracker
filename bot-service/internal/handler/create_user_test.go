package handler

import (
	"github.com/gojuno/minimock/v3"
	"github.com/google/uuid"
	"main/internal/usecase/createuser"
	"main/mocks/handler/http"
	"testing"
	"time"
)

type mocks struct {
	createUserMock *http.ICreateUserMock
	getUserMock    *http.IGetUserMock
}

func setup(t *testing.T) (*HttpHandler, mocks) {
	mc := minimock.NewController(t)

	createUserMock := http.NewICreateUserMock(mc)
	getUserMock := http.NewIGetUserMock(mc)

	sut := New(createUserMock, getUserMock)

	return sut, mocks{
		createUserMock: createUserMock,
		getUserMock:    getUserMock,
	}
}

func TestHttpHandler_CreateUser(t *testing.T) {
	t.Run("case name", func(t *testing.T) {
		t.Parallel()
		sut, sutMocks := setup(t)
		cmd := createuser.Command{
			UserId:            uuid.MustParse("c3c5f0bc-b027-46b4-b3cc-ce9675b165d5"),
			UserNickname:      "Gennadiy",
			UserCreatedAt:     time.Now(),
			UserBirthday:      nil,
			UserActiveHabitId: nil,
		}
		sutMocks.createUserMock.
			HandleMock.
			Expect(cmd).
			Return(nil)
		defer sutMocks.createUserMock.MinimockFinish()

		//что-то идет не так, как замокать HTTP-handler?
		got := sut.CreateUser().ServeHTTP()
	})
}
