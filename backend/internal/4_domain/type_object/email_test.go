package type_object

import "testing"

func TestNewEmailRejectsInvalidFormat(
	t *testing.T,
) {
	t.Parallel()

	values := []string{
		"invalid-mail-address",
		"user@example.com invalid",
		"long-long-long-long-user@example.com",
	}

	for _, value := range values {
		t.Run(value, func(t *testing.T) {
			t.Parallel()

			_, err := NewEmail(&value)
			if err == nil {
				t.Fatal("expected invalid email to return an error")
			}
		})
	}
}

func TestNewEmailAcceptsValidFormat(
	t *testing.T,
) {
	t.Parallel()

	value := "user@example.com"

	email, err := NewEmail(&value)
	if err != nil {
		t.Fatalf("expected valid email, got error: %v", err)
	}
	if email.GetValue() != value {
		t.Fatalf("expected %q, got %q", value, email.GetValue())
	}
}
