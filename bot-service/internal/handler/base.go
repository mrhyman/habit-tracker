package handler

import (
	httpMock "main/mocks/handler/http"
	"testing"
)

type mocks struct {
	createUserMock  *httpMock.ICreateUserMock
	getUserByIdMock *httpMock.IGetUserMock
}

func setup(t *testing.T) (*HttpHandler, mocks) {
	createUserMock := httpMock.NewICreateUserMock(t)
	getUserByIdMock := httpMock.NewIGetUserMock(t)

	sut := New(createUserMock, getUserByIdMock)

	return sut, mocks{
		createUserMock:  createUserMock,
		getUserByIdMock: getUserByIdMock,
	}
}
