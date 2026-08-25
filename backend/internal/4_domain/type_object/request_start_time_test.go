package type_object

import "testing"

func TestNewRequestStartTimeAcceptsProvidedValue(t *testing.T) {
	t.Parallel()

	value := int64(1700000000000)

	requestStartTime, err := NewRequestStartTime(&value)
	if err != nil {
		t.Fatalf("expected valid request start time, got error: %v", err)
	}
	if requestStartTime.GetValue() != value {
		t.Fatalf("expected %d, got %d", value, requestStartTime.GetValue())
	}
	if requestStartTime.GetString() != "1700000000000" {
		t.Fatalf("expected %q, got %q", "1700000000000", requestStartTime.GetString())
	}
}

func TestNewRequestStartTimeUsesDefaultWhenNil(t *testing.T) {
	t.Parallel()

	requestStartTime, err := NewRequestStartTime(nil)
	if err != nil {
		t.Fatalf("expected default request start time, got error: %v", err)
	}
	if requestStartTime.GetValue() == 0 {
		t.Fatal("expected default request start time to be set")
	}
}
