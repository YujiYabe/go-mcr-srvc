package usecase

import (
	"context"
	"errors"
	"strings"
	"testing"

	domain "backend/internal/4_domain"
)

func TestGetUserListWrapsGatewayError(t *testing.T) {
	t.Parallel()

	gatewayErr := errors.New("db unavailable")
	gatewayDB := &fakeGatewayDB{getListErr: gatewayErr}
	useCase := NewUseCase(nil, gatewayDB, &fakeGatewayExternal{})

	_, err := useCase.GetUserList(context.Background())
	if !errors.Is(err, gatewayErr) {
		t.Fatalf("expected wrapped gateway error, got: %v", err)
	}
	if !strings.Contains(err.Error(), "GetUserList") {
		t.Fatalf("expected usecase name in error, got: %v", err)
	}
	if !gatewayDB.getListCalled {
		t.Fatal("gateway should be called")
	}
}

func TestGetUserListByConditionRequiresCondition(t *testing.T) {
	gatewayDB := &fakeGatewayDB{}
	useCase := NewUseCase(nil, gatewayDB, &fakeGatewayExternal{})
	user := newTestUser(t, nil, nil, nil)

	_, err := useCase.GetUserListByCondition(context.Background(), user)
	if err == nil {
		t.Fatal("expected condition error")
	}
	if !strings.Contains(err.Error(), "user search condition is required") {
		t.Fatalf("unexpected error: %v", err)
	}
	if gatewayDB.getByConditionCalled {
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

			gatewayDB := &fakeGatewayDB{getByConditionErr: tt.gatewayErr}
			useCase := NewUseCase(nil, gatewayDB, &fakeGatewayExternal{})
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
			if !gatewayDB.getByConditionCalled {
				t.Fatal("gateway should be called")
			}
		})
	}
}

func TestUpdateUserRunsInTransaction(t *testing.T) {
	gatewayDB := &fakeGatewayDB{}
	useCase := NewUseCase(nil, gatewayDB, &fakeGatewayExternal{})
	user := newTestUser(t, intPointer(1), stringPointer("name"), stringPointer("test@example.com"))

	if err := useCase.UpdateUser(context.Background(), user); err != nil {
		t.Fatalf("expected update success, got: %v", err)
	}
	if !gatewayDB.runInTransactionCalled {
		t.Fatal("transaction boundary was not used")
	}
	if !gatewayDB.updateUserCalled {
		t.Fatal("update gateway was not called")
	}
}

func TestUpdateUserRequiresIdentity(t *testing.T) {
	gatewayDB := &fakeGatewayDB{}
	useCase := NewUseCase(nil, gatewayDB, &fakeGatewayExternal{})
	user := newTestUser(t, nil, stringPointer("name"), stringPointer("test@example.com"))

	err := useCase.UpdateUser(context.Background(), user)
	if err == nil {
		t.Fatal("expected identity error")
	}
	if !strings.Contains(err.Error(), "user identity is required") {
		t.Fatalf("unexpected error: %v", err)
	}
	if gatewayDB.runInTransactionCalled {
		t.Fatal("transaction should not run when user lifecycle state is invalid")
	}
	if gatewayDB.updateUserCalled {
		t.Fatal("gateway should not be called when user lifecycle state is invalid")
	}
}

func TestUpdateUserWrapsGatewayError(t *testing.T) {
	t.Parallel()

	gatewayErr := errors.New("update failed")
	gatewayDB := &fakeGatewayDB{updateUserErr: gatewayErr}
	useCase := NewUseCase(nil, gatewayDB, &fakeGatewayExternal{})
	user := newTestUser(t, intPointer(1), stringPointer("name"), stringPointer("test@example.com"))

	err := useCase.UpdateUser(context.Background(), user)
	if !errors.Is(err, gatewayErr) {
		t.Fatalf("expected wrapped gateway error, got: %v", err)
	}
	if !strings.Contains(err.Error(), "UpdateUser") {
		t.Fatalf("expected usecase name in error, got: %v", err)
	}
	if !gatewayDB.runInTransactionCalled {
		t.Fatal("transaction boundary should be used")
	}
	if !gatewayDB.updateUserCalled {
		t.Fatal("gateway should be called")
	}
}

func TestUpdateUserProfileWithPrimaryEmploymentRunsMultipleUpdatesInTransaction(t *testing.T) {
	gatewayDB := &fakeGatewayDB{}
	useCase := NewUseCase(domain.NewDomain(), gatewayDB, &fakeGatewayExternal{})
	user := newTestUser(t, intPointer(1), stringPointer("name"), stringPointer("test@example.com"))
	employment := newTestUserEmployment(t, 1, true)

	if err := useCase.UpdateUserProfileWithPrimaryEmployment(context.Background(), user, employment); err != nil {
		t.Fatalf("expected update success, got: %v", err)
	}
	if !gatewayDB.runInTransactionCalled {
		t.Fatal("transaction boundary was not used")
	}
	if !gatewayDB.updateUserCalled {
		t.Fatal("user update gateway was not called")
	}
	if !gatewayDB.updateEmploymentCalled {
		t.Fatal("employment update gateway was not called")
	}
	if strings.Join(gatewayDB.calls, ",") != "update_user,update_user_employment" {
		t.Fatalf("unexpected call order: %v", gatewayDB.calls)
	}
}

func TestUpdateUserProfileWithPrimaryEmploymentRejectsDifferentUser(t *testing.T) {
	gatewayDB := &fakeGatewayDB{}
	useCase := NewUseCase(domain.NewDomain(), gatewayDB, &fakeGatewayExternal{})
	user := newTestUser(t, intPointer(1), stringPointer("name"), stringPointer("test@example.com"))
	employment := newTestUserEmployment(t, 2, true)

	err := useCase.UpdateUserProfileWithPrimaryEmployment(context.Background(), user, employment)
	if err == nil {
		t.Fatal("expected ownership error")
	}
	if !strings.Contains(err.Error(), "must belong to the user") {
		t.Fatalf("unexpected error: %v", err)
	}
	if gatewayDB.runInTransactionCalled {
		t.Fatal("transaction should not run when employment does not belong to user")
	}
	if gatewayDB.updateUserCalled || gatewayDB.updateEmploymentCalled {
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

			gatewayDB := &fakeGatewayDB{
				updateUserErr:           tt.updateUserErr,
				updateUserEmploymentErr: tt.updateUserEmploymentErr,
			}
			useCase := NewUseCase(domain.NewDomain(), gatewayDB, &fakeGatewayExternal{})
			user := newTestUser(t, intPointer(1), stringPointer("name"), stringPointer("test@example.com"))
			employment := newTestUserEmployment(t, 1, true)

			err := useCase.UpdateUserProfileWithPrimaryEmployment(context.Background(), user, employment)
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("expected wrapped gateway error, got: %v", err)
			}
			if !strings.Contains(err.Error(), "UpdateUserProfileWithPrimaryEmployment") {
				t.Fatalf("expected usecase name in error, got: %v", err)
			}
			if !gatewayDB.runInTransactionCalled {
				t.Fatal("transaction boundary should be used")
			}
			if strings.Join(gatewayDB.calls, ",") != tt.wantCalls {
				t.Fatalf("unexpected call order: %v", gatewayDB.calls)
			}
		})
	}
}
