package handler

import (
	httpMock "main/mocks/handler/http"
	"testing"
)

type mocks struct {
	createUserMock    *httpMock.ICreateUserMock
	getUserByIdMock   *httpMock.IGetUserMock
	activateHabitMock *httpMock.IActivateHabitMock
}

func setup(t *testing.T) (*HttpHandler, mocks) {
	createUserMock := httpMock.NewICreateUserMock(t)
	getUserByIdMock := httpMock.NewIGetUserMock(t)
	activateHabitMock := httpMock.NewIActivateHabitMock(t)

	sut := New(createUserMock, getUserByIdMock, activateHabitMock)

	return sut, mocks{
		createUserMock:  createUserMock,
		getUserByIdMock: getUserByIdMock,
	}
}
