package type_object

import (
	"context"
	"testing"
)

func TestNewMailAddressRejectsInvalidFormat(t *testing.T) {
	t.Parallel()

	values := []string{
		"invalid-mail-address",
		"user@example.com invalid",
	}

	for _, value := range values {
		value := value
		t.Run(value, func(t *testing.T) {
			t.Parallel()

			mailAddress := NewMailAddress(context.Background(), &value)
			if mailAddress.GetError() == nil {
				t.Fatal("expected invalid mail address to return an error")
			}
		})
	}
}

func TestNewMailAddressAcceptsValidFormat(t *testing.T) {
	t.Parallel()

	value := "user@example.com"

	mailAddress := NewMailAddress(context.Background(), &value)
	if mailAddress.GetError() != nil {
		t.Fatalf("expected valid mail address, got error: %v", mailAddress.GetError())
	}
	if mailAddress.GetValue() != value {
		t.Fatalf("expected %q, got %q", value, mailAddress.GetValue())
	}
}
