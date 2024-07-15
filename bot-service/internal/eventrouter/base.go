package eventrouter

import (
	"main/mocks/eventrouter"
	routerMock "main/mocks/eventrouter"
	"testing"
)

type mocks struct {
	userEventRepoMock *eventrouter.IUserEventRepoMock
}

func setup(t *testing.T) (*Service, mocks) {
	userEventRepoMock := routerMock.NewIUserEventRepoMock(t)

	sut := NewService(userEventRepoMock)

	return sut, mocks{
		userEventRepoMock: userEventRepoMock,
	}
}
