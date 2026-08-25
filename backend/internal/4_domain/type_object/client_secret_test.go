package type_object

import (
	"strings"
	"testing"
)

func TestNewClientSecretAcceptsValidValue(t *testing.T) {
	t.Parallel()

	value := "client-secret"

	clientSecret, err := NewClientSecret(&value)
	if err != nil {
		t.Fatalf("expected valid client secret, got error: %v", err)
	}
	if clientSecret.GetValue() != value {
		t.Fatalf("expected %q, got %q", value, clientSecret.GetValue())
	}
}

func TestNewClientSecretRejectsTooLongValue(t *testing.T) {
	t.Parallel()

	value := strings.Repeat("a", 1000)

	_, err := NewClientSecret(&value)
	if err == nil {
		t.Fatal("expected too long client secret to return an error")
	}
}
