package type_object

import (
	"strings"
	"testing"
)

func TestNewPermissionAcceptsValidValue(t *testing.T) {
	t.Parallel()

	value := "user:read"

	permission, err := NewPermission(&value)
	if err != nil {
		t.Fatalf("expected valid permission, got error: %v", err)
	}
	if permission.GetValue() != value {
		t.Fatalf("expected %q, got %q", value, permission.GetValue())
	}
}

func TestNewPermissionRejectsInvalidValues(t *testing.T) {
	t.Parallel()

	values := []string{
		"",
		strings.Repeat("a", 51),
	}

	for _, value := range values {
		t.Run(value, func(t *testing.T) {
			t.Parallel()

			_, err := NewPermission(&value)
			if err == nil {
				t.Fatal("expected invalid permission to return an error")
			}
		})
	}
}
