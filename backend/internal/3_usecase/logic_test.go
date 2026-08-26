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
	getListCalled           bool
	getByConditionCalled    bool
	calls                   []string
	getListErr              error
	getByConditionErr       error
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
	receiver.getListCalled = true
	return groupObject.UserList{}, receiver.getListErr
}

func (receiver *fakeGatewayDB) GetUserListByCondition(
	ctx context.Context,
	reqUser groupObject.User,
) (
	groupObject.UserList,
	error,
) {
	receiver.getByConditionCalled = true
	return groupObject.UserList{}, receiver.getByConditionErr
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
	viaGRPCCalled          bool
	publishTestTopicCalled bool
	fetchAccessTokenErr    error
	viaGRPCErr             error
	publishTestTopicErr    error
}

func (receiver *fakeGatewayExternal) FetchAccessToken(
	ctx context.Context,
	credential groupObject.Credential,
) (
	typeObject.AccessToken,
	error,
) {
	receiver.fetchAccessTokenCalled = true
	if receiver.fetchAccessTokenErr != nil {
		return typeObject.AccessToken{}, receiver.fetchAccessTokenErr
	}
	return typeObject.NewAccessToken(stringPointer("access-token"))
}

func (receiver *fakeGatewayExternal) ViaGRPC(
	ctx context.Context,
	reqUser groupObject.User,
) (
	groupObject.UserList,
	error,
) {
	receiver.viaGRPCCalled = true
	return groupObject.UserList{}, receiver.viaGRPCErr
}

func (receiver *fakeGatewayExternal) PublishTestTopic(
	ctx context.Context,
) error {
	receiver.publishTestTopicCalled = true
	return receiver.publishTestTopicErr
}

func TestGetUserListWrapsGatewayError(t *testing.T) {
	t.Parallel()

	gatewayErr := errors.New("db unavailable")
	dbGateway := &fakeGatewayDB{getListErr: gatewayErr}
	useCase := NewUseCase(nil, dbGateway, &fakeGatewayExternal{})

	_, err := useCase.GetUserList(context.Background())
	if !errors.Is(err, gatewayErr) {
		t.Fatalf("expected wrapped gateway error, got: %v", err)
	}
	if !strings.Contains(err.Error(), "GetUserList") {
		t.Fatalf("expected usecase name in error, got: %v", err)
	}
	if !dbGateway.getListCalled {
		t.Fatal("gateway should be called")
	}
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

func TestGetUserListByConditionGatewayCases(t *testing.T) {
	t.Parallel()

	gatewayErr := errors.New("query failed")
	tests := []struct {
		name        string
		gatewayErr  error
		wantErr     error
		wantErrText string
	}{
		{
			name: "success",
		},
		{
			name:        "gateway error",
			gatewayErr:  gatewayErr,
			wantErr:     gatewayErr,
			wantErrText: "GetUserListByCondition",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			dbGateway := &fakeGatewayDB{getByConditionErr: tt.gatewayErr}
			useCase := NewUseCase(nil, dbGateway, &fakeGatewayExternal{})
			user := newTestUser(t, nil, stringPointer("alice"), nil)

			_, err := useCase.GetUserListByCondition(context.Background(), user)
			if tt.wantErr == nil && err != nil {
				t.Fatalf("expected success, got: %v", err)
			}
			if tt.wantErr != nil && !errors.Is(err, tt.wantErr) {
				t.Fatalf("expected wrapped gateway error, got: %v", err)
			}
			if tt.wantErrText != "" && !strings.Contains(err.Error(), tt.wantErrText) {
				t.Fatalf("expected error to contain %q, got: %v", tt.wantErrText, err)
			}
			if !dbGateway.getByConditionCalled {
				t.Fatal("gateway should be called")
			}
		})
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

func TestFetchAccessTokenGatewayCases(t *testing.T) {
	t.Parallel()

	gatewayErr := errors.New("auth service unavailable")
	tests := []struct {
		name        string
		gatewayErr  error
		wantErr     error
		wantErrText string
	}{
		{
			name: "success",
		},
		{
			name:        "gateway error",
			gatewayErr:  gatewayErr,
			wantErr:     gatewayErr,
			wantErrText: "FetchAccessToken",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			externalGateway := &fakeGatewayExternal{fetchAccessTokenErr: tt.gatewayErr}
			useCase := NewUseCase(nil, &fakeGatewayDB{}, externalGateway)
			credential := newTestCredential(t, "client-id", "secret")

			_, err := useCase.FetchAccessToken(context.Background(), credential)
			if tt.wantErr == nil && err != nil {
				t.Fatalf("expected success, got: %v", err)
			}
			if tt.wantErr != nil && !errors.Is(err, tt.wantErr) {
				t.Fatalf("expected wrapped gateway error, got: %v", err)
			}
			if tt.wantErrText != "" && !strings.Contains(err.Error(), tt.wantErrText) {
				t.Fatalf("expected error to contain %q, got: %v", tt.wantErrText, err)
			}
			if !externalGateway.fetchAccessTokenCalled {
				t.Fatal("gateway should be called")
			}
		})
	}
}

func TestViaGRPCCases(t *testing.T) {
	t.Parallel()

	gatewayErr := errors.New("grpc unavailable")
	tests := []struct {
		name                string
		gatewayErr          error
		wantErr             error
		wantErrText         string
		wantGatewayCall     bool
		buildUserInsideCase func(t *testing.T) groupObject.User
	}{
		{
			name:            "missing condition",
			wantErrText:     "user search condition is required",
			wantGatewayCall: false,
			buildUserInsideCase: func(t *testing.T) groupObject.User {
				return newTestUser(t, nil, nil, nil)
			},
		},
		{
			name:            "success",
			wantGatewayCall: true,
			buildUserInsideCase: func(t *testing.T) groupObject.User {
				return newTestUser(t, nil, nil, stringPointer("alice@example.com"))
			},
		},
		{
			name:            "gateway error",
			gatewayErr:      gatewayErr,
			wantErr:         gatewayErr,
			wantErrText:     "ViaGRPC",
			wantGatewayCall: true,
			buildUserInsideCase: func(t *testing.T) groupObject.User {
				return newTestUser(t, nil, stringPointer("alice"), nil)
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			externalGateway := &fakeGatewayExternal{viaGRPCErr: tt.gatewayErr}
			useCase := NewUseCase(nil, &fakeGatewayDB{}, externalGateway)

			_, err := useCase.ViaGRPC(context.Background(), tt.buildUserInsideCase(t))
			if tt.wantErr == nil && tt.wantErrText == "" && err != nil {
				t.Fatalf("expected success, got: %v", err)
			}
			if tt.wantErr != nil && !errors.Is(err, tt.wantErr) {
				t.Fatalf("expected wrapped gateway error, got: %v", err)
			}
			if tt.wantErrText != "" && (err == nil || !strings.Contains(err.Error(), tt.wantErrText)) {
				t.Fatalf("expected error to contain %q, got: %v", tt.wantErrText, err)
			}
			if externalGateway.viaGRPCCalled != tt.wantGatewayCall {
				t.Fatalf("gateway call = %v, want %v", externalGateway.viaGRPCCalled, tt.wantGatewayCall)
			}
		})
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

func TestUpdateUserWrapsGatewayError(t *testing.T) {
	t.Parallel()

	gatewayErr := errors.New("update failed")
	dbGateway := &fakeGatewayDB{updateUserErr: gatewayErr}
	useCase := NewUseCase(nil, dbGateway, &fakeGatewayExternal{})
	user := newTestUser(t, intPointer(1), stringPointer("name"), stringPointer("test@example.com"))

	err := useCase.UpdateUser(context.Background(), user)
	if !errors.Is(err, gatewayErr) {
		t.Fatalf("expected wrapped gateway error, got: %v", err)
	}
	if !strings.Contains(err.Error(), "UpdateUser") {
		t.Fatalf("expected usecase name in error, got: %v", err)
	}
	if !dbGateway.runInTransactionCalled {
		t.Fatal("transaction boundary should be used")
	}
	if !dbGateway.updateUserCalled {
		t.Fatal("gateway should be called")
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

func TestUpdateUserProfileWithPrimaryEmploymentGatewayErrors(t *testing.T) {
	t.Parallel()

	updateUserErr := errors.New("user update failed")
	updateEmploymentErr := errors.New("employment update failed")
	tests := []struct {
		name                    string
		updateUserErr           error
		updateUserEmploymentErr error
		wantErr                 error
		wantCalls               string
	}{
		{
			name:          "user update fails before employment update",
			updateUserErr: updateUserErr,
			wantErr:       updateUserErr,
			wantCalls:     "update_user",
		},
		{
			name:                    "employment update fails after user update",
			updateUserEmploymentErr: updateEmploymentErr,
			wantErr:                 updateEmploymentErr,
			wantCalls:               "update_user,update_user_employment",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			dbGateway := &fakeGatewayDB{
				updateUserErr:           tt.updateUserErr,
				updateUserEmploymentErr: tt.updateUserEmploymentErr,
			}
			useCase := NewUseCase(domain.NewDomain(), dbGateway, &fakeGatewayExternal{})
			user := newTestUser(t, intPointer(1), stringPointer("name"), stringPointer("test@example.com"))
			employment := newTestUserEmployment(t, 1, true)

			err := useCase.UpdateUserProfileWithPrimaryEmployment(context.Background(), user, employment)
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("expected wrapped gateway error, got: %v", err)
			}
			if !strings.Contains(err.Error(), "UpdateUserProfileWithPrimaryEmployment") {
				t.Fatalf("expected usecase name in error, got: %v", err)
			}
			if !dbGateway.runInTransactionCalled {
				t.Fatal("transaction boundary should be used")
			}
			if strings.Join(dbGateway.calls, ",") != tt.wantCalls {
				t.Fatalf("unexpected call order: %v", dbGateway.calls)
			}
		})
	}
}

func TestPublishTestTopicCases(t *testing.T) {
	t.Parallel()

	gatewayErr := errors.New("publish failed")
	tests := []struct {
		name        string
		gatewayErr  error
		wantErr     error
		wantErrText string
	}{
		{
			name: "success",
		},
		{
			name:        "gateway error",
			gatewayErr:  gatewayErr,
			wantErr:     gatewayErr,
			wantErrText: "PublishTestTopic",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			externalGateway := &fakeGatewayExternal{publishTestTopicErr: tt.gatewayErr}
			useCase := NewUseCase(nil, &fakeGatewayDB{}, externalGateway)

			err := useCase.PublishTestTopic(context.Background())
			if tt.wantErr == nil && err != nil {
				t.Fatalf("expected success, got: %v", err)
			}
			if tt.wantErr != nil && !errors.Is(err, tt.wantErr) {
				t.Fatalf("expected wrapped gateway error, got: %v", err)
			}
			if tt.wantErrText != "" && !strings.Contains(err.Error(), tt.wantErrText) {
				t.Fatalf("expected error to contain %q, got: %v", tt.wantErrText, err)
			}
			if !externalGateway.publishTestTopicCalled {
				t.Fatal("gateway should be called")
			}
		})
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
