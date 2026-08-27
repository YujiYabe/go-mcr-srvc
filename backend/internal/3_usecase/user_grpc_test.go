package usecase

import (
	"context"
	"errors"
	"strings"
	"testing"

	groupObject "backend/internal/4_domain/group_object"
)

func TestGetUserListViaGRPCCases(t *testing.T) {
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
			wantErrText:     "GetUserListViaGRPC",
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

			gatewayExternal := &fakeGatewayExternal{viaGRPCErr: tt.gatewayErr}
			useCase := NewUseCase(nil, &fakeGatewayDB{}, gatewayExternal)

			_, err := useCase.GetUserListViaGRPC(context.Background(), tt.buildUserInsideCase(t))
			if tt.wantErr == nil && tt.wantErrText == "" && err != nil {
				t.Fatalf("expected success, got: %v", err)
			}
			if tt.wantErr != nil && !errors.Is(err, tt.wantErr) {
				t.Fatalf("expected wrapped gateway error, got: %v", err)
			}
			if tt.wantErrText != "" && (err == nil || !strings.Contains(err.Error(), tt.wantErrText)) {
				t.Fatalf("expected error to contain %q, got: %v", tt.wantErrText, err)
			}
			if gatewayExternal.viaGRPCCalled != tt.wantGatewayCall {
				t.Fatalf("gateway call = %v, want %v", gatewayExternal.viaGRPCCalled, tt.wantGatewayCall)
			}
		})
	}
}
