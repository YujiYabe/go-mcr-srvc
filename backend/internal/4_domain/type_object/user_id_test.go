package type_object

import (
	"strings"
	"testing"
)

func TestNewUserIDAcceptsValidValue(
	t *testing.T,
) {
	t.Parallel()

	value := "user-123"

	userID, err := NewUserID(&value)
	if err != nil {
		t.Fatalf("expected valid user id, got error: %v", err)
	}
	if userID.GetValue() != value {
		t.Fatalf("expected %q, got %q", value, userID.GetValue())
	}
}

func TestNewUserIDRejectsTooLongValue(
	t *testing.T,
) {
	t.Parallel()

	value := strings.Repeat("a", 10)

	_, err := NewUserID(&value)
	if err == nil {
		t.Fatal("expected too long user id to return an error")
	}
}
