package type_object

import (
	"strings"
	"testing"
)

func TestNewClientIDAcceptsValidValue(
	t *testing.T,
) {
	t.Parallel()

	value := "client-id"

	clientID, err := NewClientID(&value)
	if err != nil {
		t.Fatalf("expected valid client id, got error: %v", err)
	}
	if clientID.GetValue() != value {
		t.Fatalf("expected %q, got %q", value, clientID.GetValue())
	}
}

func TestNewClientIDRejectsTooLongValue(
	t *testing.T,
) {
	t.Parallel()

	value := strings.Repeat("a", 100)

	_, err := NewClientID(&value)
	if err == nil {
		t.Fatal("expected too long client id to return an error")
	}
}
