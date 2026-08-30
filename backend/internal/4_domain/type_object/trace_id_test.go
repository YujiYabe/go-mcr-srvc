package type_object

import "testing"

func TestNewTraceIDAcceptsValidValue(
	t *testing.T,
) {
	t.Parallel()

	value := "550e8400-e29b-41d4-a716-446655440000"

	traceID, err := NewTraceID(&value)
	if err != nil {
		t.Fatalf("expected valid trace id, got error: %v", err)
	}
	if traceID.GetValue() != value {
		t.Fatalf("expected %q, got %q", value, traceID.GetValue())
	}
}

func TestNewTraceIDRejectsInvalidLength(
	t *testing.T,
) {
	t.Parallel()

	value := "short-trace-id"

	_, err := NewTraceID(&value)
	if err == nil {
		t.Fatal("expected invalid trace id length to return an error")
	}
}

func TestNewTraceIDUsesDefaultWhenNil(
	t *testing.T,
) {
	t.Parallel()

	traceID, err := NewTraceID(nil)
	if err != nil {
		t.Fatalf("expected default trace id, got error: %v", err)
	}
	if len(traceID.GetValue()) != 36 {
		t.Fatalf("expected default trace id length 36, got %d", len(traceID.GetValue()))
	}
}
