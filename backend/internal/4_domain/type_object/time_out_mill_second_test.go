package type_object

import "testing"

func TestNewTimeOutMillSecondReturnsStringValue(t *testing.T) {
	t.Parallel()

	value := int64(1500)

	timeout, err := NewTimeOutMillSecond(&value)
	if err != nil {
		t.Fatalf("expected valid timeout, got error: %v", err)
	}
	if timeout.GetValue() != value {
		t.Fatalf("expected %d, got %d", value, timeout.GetValue())
	}
	if timeout.GetString() != "1500" {
		t.Fatalf("expected %q, got %q", "1500", timeout.GetString())
	}
}

func TestNewTimeOutMillSecondRejectsTooManyDigits(t *testing.T) {
	t.Parallel()

	value := int64(10000000000)

	_, err := NewTimeOutMillSecond(&value)
	if err == nil {
		t.Fatal("expected timeout with too many digits to return an error")
	}
}
