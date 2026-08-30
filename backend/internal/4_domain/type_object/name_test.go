package type_object

import "testing"

func TestNewNameAcceptsValidValue(
	t *testing.T,
) {
	t.Parallel()

	value := "Alice"

	name, err := NewName(&value)
	if err != nil {
		t.Fatalf("expected valid name, got error: %v", err)
	}
	if name.GetValue() != value {
		t.Fatalf("expected %q, got %q", value, name.GetValue())
	}
}

func TestNewNameRejectsInvalidValues(
	t *testing.T,
) {
	t.Parallel()

	values := []string{
		"",
		"1234567890123456789012345678901",
	}

	for _, value := range values {
		t.Run(value, func(t *testing.T) {
			t.Parallel()

			_, err := NewName(&value)
			if err == nil {
				t.Fatal("expected invalid name to return an error")
			}
		})
	}
}

func TestNewNameRejectsConfiguredCheckSpell(
	t *testing.T,
) {
	t.Parallel()

	value := "暴力"
	checkSpell := []string{
		"盗む",
		"暴力",
	}

	_, err := NewName(&value, checkSpell)
	if err == nil {
		t.Fatal("expected configured check spell to return an error")
	}
}
