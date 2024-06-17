package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"main/internal/domain"
	"main/internal/usecase/createuser"
	"main/internal/usecase/getuserbyid"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHttpHandler_CreateUser(t *testing.T) {
	t.Parallel()
	var (
		userId    = uuid.New()
		nickname  = uuid.NewString()
		cmd       = createuser.Command{UserId: userId, UserNickname: nickname}
		validUser = domain.User{Id: userId, Nickname: nickname}
		someError = errors.New("some error")
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
				jsonStr, _ := json.Marshal(validUser)
				req, _ := http.NewRequest("POST", "/createUser", bytes.NewBuffer(jsonStr))
				return args{
					req: req,
				}
			},
			setMocks: func(m mocks) {
				m.createUserMock.HandleMock.When(cmd).Then(nil)

			},
			wantCode: http.StatusOK,
			wantBody: fmt.Sprintf("Person ID: %s", validUser.Id),
		},
		{
			name: "400 user decode error",
			args: func(*testing.T) args {
				jsonStr, _ := json.Marshal("{}")
				req, _ := http.NewRequest("POST", "/createUser", bytes.NewBuffer(jsonStr))
				return args{
					req: req,
				}
			},
			setMocks: func(m mocks) {},
			wantCode: http.StatusBadRequest,
			wantBody: "json: cannot unmarshal string into Go value of type handler.UserModel\n",
		},
		{
			name: "400 command create error",
			args: func(*testing.T) args {
				jsonStr := []byte(`{"Id":"","Nickname":"John Doe"}`)
				req, _ := http.NewRequest("POST", "/createUser", bytes.NewBuffer(jsonStr))
				return args{
					req: req,
				}
			},
			setMocks: func(m mocks) {},
			wantCode: http.StatusBadRequest,
			wantBody: fmt.Sprintf("%s\n", getuserbyid.ErrInvalidUserID.Error()),
		},
		{
			name: "500 command handle error",
			args: func(*testing.T) args {
				jsonStr, _ := json.Marshal(validUser)
				req, _ := http.NewRequest("POST", "/createUser", bytes.NewBuffer(jsonStr))
				return args{
					req: req,
				}
			},
			setMocks: func(m mocks) {
				m.createUserMock.HandleMock.When(cmd).Then(someError)
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
			mux := sut.SetupMux()
			tt.setMocks(mocks)
			resp := httptest.NewRecorder()
			mux.ServeHTTP(resp, tt.args(t).req)

			require.Equal(t, tt.wantCode, resp.Result().StatusCode)
			require.Equal(t, tt.wantBody, resp.Body.String())
		})
	}
}
