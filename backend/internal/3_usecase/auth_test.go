package usecase

import (
	"context"
	"errors"
	"strings"
	"testing"
)

func TestFetchAccessTokenRequiresCredential(t *testing.T) {
	gatewayExternal := &fakeGatewayExternal{}
	useCase := NewUseCase(nil, &fakeGatewayDB{}, gatewayExternal)
	credential := newTestCredential(t, "", "secret")

	_, err := useCase.FetchAccessToken(context.Background(), credential)
	if err == nil {
		t.Fatal("expected credential error")
	}
	if !strings.Contains(err.Error(), "client id is required") {
		t.Fatalf("unexpected error: %v", err)
	}
	if gatewayExternal.fetchAccessTokenCalled {
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
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			gatewayExternal := &fakeGatewayExternal{fetchAccessTokenErr: tt.gatewayErr}
			useCase := NewUseCase(nil, &fakeGatewayDB{}, gatewayExternal)
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
			if !gatewayExternal.fetchAccessTokenCalled {
				t.Fatal("gateway should be called")
			}
		})
	}
}
