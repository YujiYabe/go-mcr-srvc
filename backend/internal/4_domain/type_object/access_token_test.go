package type_object

import (
	"strings"
	"testing"
)

func TestNewAccessTokenAcceptsValidValue(t *testing.T) {
	t.Parallel()

	value := "access-token"

	accessToken, err := NewAccessToken(&value)
	if err != nil {
		t.Fatalf("expected valid access token, got error: %v", err)
	}
	if accessToken.GetValue() != value {
		t.Fatalf("expected %q, got %q", value, accessToken.GetValue())
	}
}

func TestNewAccessTokenRejectsTooLongValue(t *testing.T) {
	t.Parallel()

	value := strings.Repeat("a", 10000)

	_, err := NewAccessToken(&value)
	if err == nil {
		t.Fatal("expected too long access token to return an error")
	}
}
