package type_object

import (
	"strings"
	"testing"
)

func TestNewTenantIDAcceptsValidValue(
	t *testing.T,
) {
	t.Parallel()

	value := "tenant-1"

	tenantID, err := NewTenantID(&value)
	if err != nil {
		t.Fatalf("expected valid tenant id, got error: %v", err)
	}
	if tenantID.GetValue() != value {
		t.Fatalf("expected %q, got %q", value, tenantID.GetValue())
	}
}

func TestNewTenantIDRejectsInvalidValues(
	t *testing.T,
) {
	t.Parallel()

	values := []string{
		"",
		strings.Repeat("a", 100),
	}

	for _, value := range values {
		t.Run(value, func(t *testing.T) {
			t.Parallel()

			_, err := NewTenantID(&value)
			if err == nil {
				t.Fatal("expected invalid tenant id to return an error")
			}
		})
	}
}
