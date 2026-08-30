package type_object

import "testing"

func TestNewTimeZoneAcceptsValidValue(
	t *testing.T,
) {
	t.Parallel()

	value := "UTC"

	timeZone, err := NewTimeZone(&value)
	if err != nil {
		t.Fatalf("expected valid timezone, got error: %v", err)
	}
	if timeZone.GetValue() != value {
		t.Fatalf("expected %q, got %q", value, timeZone.GetValue())
	}
}

func TestNewTimeZoneRejectsTooLongValue(
	t *testing.T,
) {
	t.Parallel()

	value := "America/New_York"

	_, err := NewTimeZone(&value)
	if err == nil {
		t.Fatal("expected too long timezone to return an error")
	}
}
