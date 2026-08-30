package type_object

import (
	"strings"
	"testing"
)

func TestNewClientIPAcceptsValidValue(
	t *testing.T,
) {
	t.Parallel()

	value := "192.168.0.1"

	clientIP, err := NewClientIP(&value)
	if err != nil {
		t.Fatalf("expected valid client ip, got error: %v", err)
	}
	if clientIP.GetValue() != value {
		t.Fatalf("expected %q, got %q", value, clientIP.GetValue())
	}
}

func TestNewClientIPRejectsTooLongValue(
	t *testing.T,
) {
	t.Parallel()

	value := strings.Repeat("1", 31)

	_, err := NewClientIP(&value)
	if err == nil {
		t.Fatal("expected too long client ip to return an error")
	}
}
