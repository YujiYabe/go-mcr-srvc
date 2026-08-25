package type_object

import "testing"

func TestNewIDAcceptsValidDigitRange(t *testing.T) {
	t.Parallel()

	value := 123456789

	id, err := NewID(&value)
	if err != nil {
		t.Fatalf("expected valid id, got error: %v", err)
	}
	if id.GetValue() != value {
		t.Fatalf("expected %d, got %d", value, id.GetValue())
	}
}

func TestNewIDRejectsTooManyDigits(t *testing.T) {
	t.Parallel()

	value := 1234567890

	_, err := NewID(&value)
	if err == nil {
		t.Fatal("expected id with too many digits to return an error")
	}
}
