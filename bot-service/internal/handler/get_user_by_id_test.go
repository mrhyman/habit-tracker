package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"main/internal/domain"
	"main/internal/usecase/getuserbyid"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHttpHandler_GetUserById(t *testing.T) {
	t.Parallel()

	var (
		userId    = uuid.New()
		nickname  = uuid.NewString()
		validUser = domain.User{Id: userId, Nickname: nickname}
		query     = getuserbyid.Query{UserID: userId}
		body, _   = json.Marshal(UserFromDomain(&validUser))
		someError = fmt.Errorf("some error")
		ctx       = context.Background()
	)

	type args struct {
		req *http.Request
	}

	tests := []struct {
		name     string
		args     func(t *testing.T) args
		wantCode int
		setMocks func(m mocks)
		wantBody string
	}{
		{
			name: "200 ok",
			args: func(*testing.T) args {
				req, _ := http.NewRequest("GET", "/getUser", nil)
				q := req.URL.Query()
				q.Add("id", userId.String())

				req.URL.RawQuery = q.Encode()
				return args{
					req: req,
				}
			},
			setMocks: func(m mocks) {
				m.getUserByIdMock.HandleMock.When(ctx, query).Then(&validUser, nil)
			},
			wantCode: http.StatusOK,
			wantBody: string(body),
		},
		{
			name: "400 create query error",
			args: func(*testing.T) args {
				req, _ := http.NewRequest("GET", "/getUser", nil)
				q := req.URL.Query()
				q.Add("id", "some_string")

				req.URL.RawQuery = q.Encode()
				return args{
					req: req,
				}
			},
			setMocks: func(m mocks) {},
			wantCode: http.StatusBadRequest,
			wantBody: "invalid argument\nuser ID should be a valid UUID\n",
		},
		{
			name: "404 user not found",
			args: func(*testing.T) args {
				req, _ := http.NewRequest("GET", "/getUser", nil)
				q := req.URL.Query()
				q.Add("id", userId.String())

				req.URL.RawQuery = q.Encode()
				return args{
					req: req,
				}
			},
			setMocks: func(m mocks) {
				m.getUserByIdMock.HandleMock.When(ctx, query).Then(nil, domain.ErrUserNotFound)
			},
			wantCode: http.StatusNotFound,
			wantBody: fmt.Sprintf("%s\n", domain.ErrUserNotFound.Error()),
		},
		{
			name: "500 query handle error",
			args: func(*testing.T) args {
				req, _ := http.NewRequest("GET", "/getUser", nil)
				q := req.URL.Query()
				q.Add("id", userId.String())

				req.URL.RawQuery = q.Encode()
				return args{
					req: req,
				}
			},
			setMocks: func(m mocks) {
				m.getUserByIdMock.HandleMock.When(ctx, query).Then(nil, someError)
			},
			wantCode: http.StatusInternalServerError,
			wantBody: fmt.Sprintf("%s\n", someError.Error()),
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			sut, mocks := setup(t)
			tt.setMocks(mocks)

			resp := httptest.NewRecorder()
			sut.GetUserById().ServeHTTP(resp, tt.args(t).req)
			require.Equal(t, tt.wantCode, resp.Result().StatusCode)
			require.Equal(t, tt.wantBody, resp.Body.String())
		})
	}
}
