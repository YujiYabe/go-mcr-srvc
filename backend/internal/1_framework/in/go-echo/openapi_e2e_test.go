package goEcho

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"

	"backend/internal/1_framework/in/go-echo/openapi"
	groupObject "backend/internal/4_domain/group_object"
	typeObject "backend/internal/4_domain/type_object"
)

type fakeController struct {
	getUserListByConditionCalled bool
	publishTestTopicCalled       bool
	lastReqUser                  groupObject.User
	userList                     groupObject.UserList
}

func newE2ETestEcho(t *testing.T, controller *fakeController) *echo.Echo {
	t.Helper()

	echoEcho := NewEcho()
	openapi.RegisterHandlers(
		echoEcho,
		&ServerInterfaceImpl{
			Controller: controller,
		},
	)

	return echoEcho
}

func TestOpenAPIE2E_V1UsersGet(t *testing.T) {
	name := "Alice"
	email := "alice@example.com"
	id := 1
	userList, err := groupObject.NewUserList(
		&groupObject.NewUserListArgs{
			Content: []groupObject.NewUserArgs{
				{
					ID:    &id,
					Name:  &name,
					Email: &email,
				},
			},
		},
	)
	if err != nil {
		t.Fatalf("failed to build user list: %v", err)
	}

	controller := &fakeController{
		userList: userList,
	}
	echoEcho := newE2ETestEcho(t, controller)

	req := httptest.NewRequestWithContext(
		context.Background(),
		http.MethodGet,
		"/v1/users?name=Alice&email=alice@example.com",
		nil,
	)
	req.Header.Set(echo.HeaderAccept, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	echoEcho.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d: %s", http.StatusOK, rec.Code, rec.Body.String())
	}
	if !controller.getUserListByConditionCalled {
		t.Fatal("expected GetUserListByCondition to be called")
	}
	if got := controller.lastReqUser.Name().GetValue(); got != name {
		t.Fatalf("expected query name %q, got %q", name, got)
	}
	if got := controller.lastReqUser.Email().GetValue(); got != email {
		t.Fatalf("expected query email %q, got %q", email, got)
	}

	var body []struct {
		ID    int    `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}
	if err := json.Unmarshal(rec.Body.Bytes(), &body); err != nil {
		t.Fatalf("failed to decode response body: %v", err)
	}
	if len(body) != 1 {
		t.Fatalf("expected 1 user, got %d", len(body))
	}
	if body[0].ID != id || body[0].Name != name || body[0].Email != email {
		t.Fatalf("unexpected response body: %+v", body[0])
	}
}

func TestOpenAPIE2E_V1UsersPost(t *testing.T) {
	controller := &fakeController{}
	echoEcho := newE2ETestEcho(t, controller)

	reqBody := bytes.NewBufferString(`{"name":"Bob","email":"bob@example.com"}`)
	req := httptest.NewRequestWithContext(
		context.Background(),
		http.MethodPost,
		"/v1/users",
		reqBody,
	)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	echoEcho.ServeHTTP(rec, req)

	if rec.Code != http.StatusCreated {
		t.Fatalf("expected status %d, got %d: %s", http.StatusCreated, rec.Code, rec.Body.String())
	}

	var body struct {
		ID    int    `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}
	if err := json.Unmarshal(rec.Body.Bytes(), &body); err != nil {
		t.Fatalf("failed to decode response body: %v", err)
	}
	if body.ID != 3 || body.Name != "Bob" || body.Email != "bob@example.com" {
		t.Fatalf("unexpected response body: %+v", body)
	}
}

func TestOpenAPIE2E_V1HealthGet(t *testing.T) {
	controller := &fakeController{}
	echoEcho := newE2ETestEcho(t, controller)

	req := httptest.NewRequestWithContext(
		context.Background(),
		http.MethodGet,
		"/v1/health",
		nil,
	)
	rec := httptest.NewRecorder()

	echoEcho.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d: %s", http.StatusOK, rec.Code, rec.Body.String())
	}
	if rec.Body.String() != "OK" {
		t.Fatalf("expected body %q, got %q", "OK", rec.Body.String())
	}
}

func TestOpenAPIE2E_V1ToPubsubGet(t *testing.T) {
	controller := &fakeController{}
	echoEcho := newE2ETestEcho(t, controller)

	req := httptest.NewRequestWithContext(
		context.Background(),
		http.MethodGet,
		"/v1/to-pubsub",
		nil,
	)
	rec := httptest.NewRecorder()

	echoEcho.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d: %s", http.StatusOK, rec.Code, rec.Body.String())
	}
	if !controller.publishTestTopicCalled {
		t.Fatal("expected PublishTestTopic to be called")
	}
}

func (receiver *fakeController) GetUserList(
	_ context.Context,
) (
	userList groupObject.UserList,
	err error,
) {
	return groupObject.UserList{}, nil
}

func (receiver *fakeController) GetUserListByCondition(
	_ context.Context,
	reqUser groupObject.User,
) (
	resUserList groupObject.UserList,
	err error,
) {
	receiver.getUserListByConditionCalled = true
	receiver.lastReqUser = reqUser

	return receiver.userList, nil
}

func (receiver *fakeController) FetchAccessToken(
	_ context.Context,
	_ groupObject.Credential,
) (
	accessToken typeObject.AccessToken,
	err error,
) {
	return typeObject.AccessToken{}, nil
}

func (receiver *fakeController) GetUserListViaGRPC(
	_ context.Context,
	_ groupObject.User,
) (
	resUserList groupObject.UserList,
	err error,
) {
	return groupObject.UserList{}, nil
}

func (receiver *fakeController) UpdateUser(
	_ context.Context,
	_ groupObject.User,
) error {
	return nil
}

func (receiver *fakeController) UpdateUserProfileWithPrimaryEmployment(
	_ context.Context,
	_ groupObject.User,
	_ groupObject.UserEmployment,
) error {
	return nil
}

func (receiver *fakeController) PublishTestTopic(
	_ context.Context,
) error {
	receiver.publishTestTopicCalled = true

	return nil
}
