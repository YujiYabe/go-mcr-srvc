package usecase

import (
	"context"
	"errors"
	"strings"
	"testing"

	domain "backend/internal/4_domain"
	groupObject "backend/internal/4_domain/group_object"
	typeObject "backend/internal/4_domain/type_object"
)

type fakeGatewayDB struct {
	runInTransactionCalled  bool
	updateUserCalled        bool
	updateEmploymentCalled  bool
	getByConditionCalled    bool
	calls                   []string
	updateUserErr           error
	updateUserEmploymentErr error
}

func (receiver *fakeGatewayDB) RunInTransaction(
	ctx context.Context,
	fn func(context.Context) error,
) error {
	receiver.runInTransactionCalled = true
	return fn(ctx)
}

func (receiver *fakeGatewayDB) GetUserList(
	ctx context.Context,
) (
	groupObject.UserList,
	error,
) {
	return groupObject.UserList{}, nil
}

func (receiver *fakeGatewayDB) GetUserListByCondition(
	ctx context.Context,
	reqUser groupObject.User,
) (
	groupObject.UserList,
	error,
) {
	receiver.getByConditionCalled = true
	return groupObject.UserList{}, nil
}

func (receiver *fakeGatewayDB) UpdateUser(
	ctx context.Context,
	newUser groupObject.User,
) error {
	receiver.updateUserCalled = true
	receiver.calls = append(receiver.calls, "update_user")
	return receiver.updateUserErr
}

func (receiver *fakeGatewayDB) UpdateUserEmployment(
	ctx context.Context,
	userEmployment groupObject.UserEmployment,
) error {
	receiver.updateEmploymentCalled = true
	receiver.calls = append(receiver.calls, "update_user_employment")
	return receiver.updateUserEmploymentErr
}

type fakeGatewayExternal struct {
	fetchAccessTokenCalled bool
}

func (receiver *fakeGatewayExternal) FetchAccessToken(
	ctx context.Context,
	credential groupObject.Credential,
) (
	typeObject.AccessToken,
	error,
) {
	receiver.fetchAccessTokenCalled = true
	return typeObject.NewAccessToken(stringPointer("access-token"))
}

func (receiver *fakeGatewayExternal) ViaGRPC(
	ctx context.Context,
	reqUser groupObject.User,
) (
	groupObject.UserList,
	error,
) {
	return groupObject.UserList{}, nil
}

func (receiver *fakeGatewayExternal) PublishTestTopic(
	ctx context.Context,
) error {
	return nil
}

func TestGetUserListByConditionRequiresCondition(t *testing.T) {
	dbGateway := &fakeGatewayDB{}
	useCase := NewUseCase(nil, dbGateway, &fakeGatewayExternal{})
	user := newTestUser(t, nil, nil, nil)

	_, err := useCase.GetUserListByCondition(context.Background(), user)
	if err == nil {
		t.Fatal("expected condition error")
	}
	if !strings.Contains(err.Error(), "user search condition is required") {
		t.Fatalf("unexpected error: %v", err)
	}
	if dbGateway.getByConditionCalled {
		t.Fatal("gateway should not be called when condition is empty")
	}
}

func TestFetchAccessTokenRequiresCredential(t *testing.T) {
	externalGateway := &fakeGatewayExternal{}
	useCase := NewUseCase(nil, &fakeGatewayDB{}, externalGateway)
	credential := newTestCredential(t, "", "secret")

	_, err := useCase.FetchAccessToken(context.Background(), credential)
	if err == nil {
		t.Fatal("expected credential error")
	}
	if !strings.Contains(err.Error(), "client id is required") {
		t.Fatalf("unexpected error: %v", err)
	}
	if externalGateway.fetchAccessTokenCalled {
		t.Fatal("gateway should not be called when credential is invalid")
	}
}

func TestUpdateUserRunsInTransaction(t *testing.T) {
	dbGateway := &fakeGatewayDB{}
	useCase := NewUseCase(nil, dbGateway, &fakeGatewayExternal{})
	user := newTestUser(t, intPointer(1), stringPointer("name"), stringPointer("test@example.com"))

	if err := useCase.UpdateUser(context.Background(), user); err != nil {
		t.Fatalf("expected update success, got: %v", err)
	}
	if !dbGateway.runInTransactionCalled {
		t.Fatal("transaction boundary was not used")
	}
	if !dbGateway.updateUserCalled {
		t.Fatal("update gateway was not called")
	}
}

func TestUpdateUserRequiresIdentity(t *testing.T) {
	dbGateway := &fakeGatewayDB{}
	useCase := NewUseCase(nil, dbGateway, &fakeGatewayExternal{})
	user := newTestUser(t, nil, stringPointer("name"), stringPointer("test@example.com"))

	err := useCase.UpdateUser(context.Background(), user)
	if err == nil {
		t.Fatal("expected identity error")
	}
	if !strings.Contains(err.Error(), "user identity is required") {
		t.Fatalf("unexpected error: %v", err)
	}
	if dbGateway.runInTransactionCalled {
		t.Fatal("transaction should not run when user lifecycle state is invalid")
	}
	if dbGateway.updateUserCalled {
		t.Fatal("gateway should not be called when user lifecycle state is invalid")
	}
}

func TestUpdateUserProfileWithPrimaryEmploymentRunsMultipleUpdatesInTransaction(t *testing.T) {
	dbGateway := &fakeGatewayDB{}
	useCase := NewUseCase(domain.NewDomain(), dbGateway, &fakeGatewayExternal{})
	user := newTestUser(t, intPointer(1), stringPointer("name"), stringPointer("test@example.com"))
	employment := newTestUserEmployment(t, 1, true)

	if err := useCase.UpdateUserProfileWithPrimaryEmployment(context.Background(), user, employment); err != nil {
		t.Fatalf("expected update success, got: %v", err)
	}
	if !dbGateway.runInTransactionCalled {
		t.Fatal("transaction boundary was not used")
	}
	if !dbGateway.updateUserCalled {
		t.Fatal("user update gateway was not called")
	}
	if !dbGateway.updateEmploymentCalled {
		t.Fatal("employment update gateway was not called")
	}
	if strings.Join(dbGateway.calls, ",") != "update_user,update_user_employment" {
		t.Fatalf("unexpected call order: %v", dbGateway.calls)
	}
}

func TestUpdateUserProfileWithPrimaryEmploymentRejectsDifferentUser(t *testing.T) {
	dbGateway := &fakeGatewayDB{}
	useCase := NewUseCase(domain.NewDomain(), dbGateway, &fakeGatewayExternal{})
	user := newTestUser(t, intPointer(1), stringPointer("name"), stringPointer("test@example.com"))
	employment := newTestUserEmployment(t, 2, true)

	err := useCase.UpdateUserProfileWithPrimaryEmployment(context.Background(), user, employment)
	if err == nil {
		t.Fatal("expected ownership error")
	}
	if !strings.Contains(err.Error(), "must belong to the user") {
		t.Fatalf("unexpected error: %v", err)
	}
	if dbGateway.runInTransactionCalled {
		t.Fatal("transaction should not run when employment does not belong to user")
	}
	if dbGateway.updateUserCalled || dbGateway.updateEmploymentCalled {
		t.Fatal("gateway should not be called when employment does not belong to user")
	}
}

func TestEnsureContextReadyReturnsCanceledError(t *testing.T) {
	useCase := NewUseCase(nil, &fakeGatewayDB{}, &fakeGatewayExternal{})
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	_, err := useCase.GetUserList(ctx)
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("expected context.Canceled, got: %v", err)
	}
}

func newTestUser(
	t *testing.T,
	id *int,
	name *string,
	email *string,
) groupObject.User {
	t.Helper()

	user, err := groupObject.NewUser(&groupObject.NewUserArgs{
		ID:    id,
		Name:  name,
		Email: email,
	})
	if err != nil {
		t.Fatalf("failed to create user: %v", err)
	}

	return *user
}

func newTestCredential(
	t *testing.T,
	clientID string,
	clientSecret string,
) groupObject.Credential {
	t.Helper()

	credential, err := groupObject.NewCredential(&groupObject.NewCredentialArgs{
		ClientID:     &clientID,
		ClientSecret: &clientSecret,
	})
	if err != nil {
		t.Fatalf("failed to create credential: %v", err)
	}

	return *credential
}

func newTestUserEmployment(
	t *testing.T,
	userID int,
	isPrimary bool,
) groupObject.UserEmployment {
	t.Helper()

	employeeCode := "EMP001"
	employmentType := "full_time"
	userEmployment, err := groupObject.NewUserEmployment(&groupObject.NewUserEmploymentArgs{
		UserID:         intPointer(userID),
		CompanyID:      intPointer(1),
		DepartmentID:   intPointer(2),
		PositionID:     intPointer(3),
		EmployeeCode:   &employeeCode,
		EmploymentType: &employmentType,
		IsPrimary:      &isPrimary,
	})
	if err != nil {
		t.Fatalf("failed to create user employment: %v", err)
	}

	return *userEmployment
}

func intPointer(value int) *int {
	return &value
}

func stringPointer(value string) *string {
	return &value
}
