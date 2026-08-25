package type_object

import "testing"

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

			_, err := NewMailAddress(&value)
			if err == nil {
				t.Fatal("expected invalid mail address to return an error")
			}
		})
	}
}

func TestNewMailAddressAcceptsValidFormat(t *testing.T) {
	t.Parallel()

	value := "user@example.com"

	mailAddress, err := NewMailAddress(&value)
	if err != nil {
		t.Fatalf("expected valid mail address, got error: %v", err)
	}
	if mailAddress.GetValue() != value {
		t.Fatalf("expected %q, got %q", value, mailAddress.GetValue())
	}
}
