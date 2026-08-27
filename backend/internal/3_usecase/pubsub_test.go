package usecase

import (
	"context"
	"errors"
	"strings"
	"testing"
)

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
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			gatewayExternal := &fakeGatewayExternal{publishTestTopicErr: tt.gatewayErr}
			useCase := NewUseCase(nil, &fakeGatewayDB{}, gatewayExternal)

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
			if !gatewayExternal.publishTestTopicCalled {
				t.Fatal("gateway should be called")
			}
		})
	}
}
