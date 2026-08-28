package usecase

import (
	"context"
	"errors"
	"strings"
	"testing"
)

func TestValidationWordRuleUpdatesCallGateway(t *testing.T) {
	tests := []struct {
		name       string
		run        func(ToUseCase) error
		assertCall func(*fakeGatewayDB) bool
	}{
		{
			name: "add",
			run: func(useCase ToUseCase) error {
				return useCase.AddValidationWord(context.Background(), "name", true, "root")
			},
			assertCall: func(gatewayDB *fakeGatewayDB) bool {
				return gatewayDB.addValidationWordCalled
			},
		},
		{
			name: "update",
			run: func(useCase ToUseCase) error {
				return useCase.UpdateValidationWord(context.Background(), "name", true, "root", "admin")
			},
			assertCall: func(gatewayDB *fakeGatewayDB) bool {
				return gatewayDB.updateValidationWordCalled
			},
		},
		{
			name: "delete",
			run: func(useCase ToUseCase) error {
				return useCase.DeleteValidationWord(context.Background(), "name", true, "root")
			},
			assertCall: func(gatewayDB *fakeGatewayDB) bool {
				return gatewayDB.deleteValidationWordCalled
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gatewayDB := &fakeGatewayDB{}
			useCase := NewUseCase(nil, gatewayDB, &fakeGatewayExternal{})

			if err := tt.run(useCase); err != nil {
				t.Fatalf("expected success, got: %v", err)
			}
			if !tt.assertCall(gatewayDB) {
				t.Fatal("gateway should be called")
			}
		})
	}
}

func TestValidationWordRuleUpdatesWrapGatewayError(t *testing.T) {
	gatewayErr := errors.New("gateway failed")
	gatewayDB := &fakeGatewayDB{validationWordUpdateErr: gatewayErr}
	useCase := NewUseCase(nil, gatewayDB, &fakeGatewayExternal{})

	err := useCase.AddValidationWord(context.Background(), "name", true, "root")
	if !errors.Is(err, gatewayErr) {
		t.Fatalf("expected wrapped gateway error, got: %v", err)
	}
	if !strings.Contains(err.Error(), "AddValidationWord") {
		t.Fatalf("expected usecase name in error, got: %v", err)
	}
}
