package type_object

import (
	"strings"
	"testing"
)

func TestNewLocaleAcceptsValidValue(t *testing.T) {
	t.Parallel()

	value := "ja-JP"

	locale, err := NewLocale(&value)
	if err != nil {
		t.Fatalf("expected valid locale, got error: %v", err)
	}
	if locale.GetValue() != value {
		t.Fatalf("expected %q, got %q", value, locale.GetValue())
	}
}

func TestNewLocaleRejectsTooLongValue(t *testing.T) {
	t.Parallel()

	value := strings.Repeat("a", 21)

	_, err := NewLocale(&value)
	if err == nil {
		t.Fatal("expected too long locale to return an error")
	}
}
