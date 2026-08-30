package type_object

import (
	"strings"
	"testing"
)

func TestNewUserAgentAcceptsValidValue(
	t *testing.T,
) {
	t.Parallel()

	value := "Mozilla/5.0"

	userAgent, err := NewUserAgent(&value)
	if err != nil {
		t.Fatalf("expected valid user agent, got error: %v", err)
	}
	if userAgent.GetValue() != value {
		t.Fatalf("expected %q, got %q", value, userAgent.GetValue())
	}
}

func TestNewUserAgentRejectsTooLongValue(
	t *testing.T,
) {
	t.Parallel()

	value := strings.Repeat("a", 51)

	_, err := NewUserAgent(&value)
	if err == nil {
		t.Fatal("expected too long user agent to return an error")
	}
}
