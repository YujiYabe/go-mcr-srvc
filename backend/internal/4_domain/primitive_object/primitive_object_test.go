package primitive_object

import "testing"

func TestPrimitiveStringConvenienceMethods(
	t *testing.T,
) {
	t.Parallel()

	value := "東京"
	primitiveString := &PrimitiveString{}
	content := NewPrimitiveString(
		primitiveString.WithValue(&value),
	)

	if content.Length() != 2 {
		t.Fatalf("expected length 2, got %d", content.Length())
	}
	if !content.HasValue() {
		t.Fatal("expected content to have value")
	}
	if content.IsEmpty() {
		t.Fatal("expected content not to be empty")
	}
	if !content.Equal(value) {
		t.Fatalf("expected content to equal %q", value)
	}
}

func TestPrimitiveStringNilConvenienceMethods(
	t *testing.T,
) {
	t.Parallel()

	content := NewPrimitiveString()

	if content.Length() != 0 {
		t.Fatalf("expected length 0, got %d", content.Length())
	}
	if content.HasValue() {
		t.Fatal("expected nil content not to have value")
	}
	if !content.IsEmpty() {
		t.Fatal("expected nil content to be empty")
	}
	if content.Equal("") {
		t.Fatal("expected nil content not to equal empty string")
	}
}

func TestPrimitiveIntXConvenienceMethods(
	t *testing.T,
) {
	t.Parallel()

	value := -123
	primitiveIntX := &PrimitiveIntX[int]{}
	content := NewPrimitiveIntX(
		primitiveIntX.WithValue(&value),
	)

	if !content.HasValue() {
		t.Fatal("expected content to have value")
	}
	if content.IsZero() {
		t.Fatal("expected content not to be zero")
	}
	if content.DigitCount() != 3 {
		t.Fatalf("expected digit count 3, got %d", content.DigitCount())
	}
	if !content.Equal(value) {
		t.Fatalf("expected content to equal %d", value)
	}
}

func TestPrimitiveUIntXConvenienceMethods(
	t *testing.T,
) {
	t.Parallel()

	value := uint64(123)
	primitiveUIntX := &PrimitiveUIntX[uint64]{}
	content := NewPrimitiveUIntX(
		primitiveUIntX.WithValue(&value),
	)

	if !content.HasValue() {
		t.Fatal("expected content to have value")
	}
	if content.IsZero() {
		t.Fatal("expected content not to be zero")
	}
	if content.DigitCount() != 3 {
		t.Fatalf("expected digit count 3, got %d", content.DigitCount())
	}
	if !content.Equal(value) {
		t.Fatalf("expected content to equal %d", value)
	}
}

func TestPrimitiveSliceIntConvenienceMethods(
	t *testing.T,
) {
	t.Parallel()

	firstValue := 1
	secondValue := 2
	primitiveIntX := &PrimitiveIntX[int]{}
	primitiveSliceInt := &PrimitiveSliceInt{}
	content := NewPrimitiveSliceInt(
		primitiveSliceInt.WithValue([]PrimitiveIntX[int]{
			*NewPrimitiveIntX(primitiveIntX.WithValue(&firstValue)),
			*NewPrimitiveIntX(primitiveIntX.WithValue(&secondValue)),
		}),
	)

	if content.Count() != 2 {
		t.Fatalf("expected count 2, got %d", content.Count())
	}
	if content.IsEmpty() {
		t.Fatal("expected content not to be empty")
	}
	if !content.HasValue(secondValue) {
		t.Fatalf("expected content to have value %d", secondValue)
	}
}

func TestPrimitiveSliceStringConvenienceMethods(
	t *testing.T,
) {
	t.Parallel()

	firstValue := "read"
	secondValue := "write"
	primitiveString := &PrimitiveString{}
	primitiveSliceString := &PrimitiveSliceString{}
	content := NewPrimitiveSliceString(
		primitiveSliceString.WithValue([]PrimitiveString{
			*NewPrimitiveString(primitiveString.WithValue(&firstValue)),
			*NewPrimitiveString(primitiveString.WithValue(&secondValue)),
		}),
	)

	if content.Count() != 2 {
		t.Fatalf("expected count 2, got %d", content.Count())
	}
	if content.IsEmpty() {
		t.Fatal("expected content not to be empty")
	}
	if !content.HasValue(secondValue) {
		t.Fatalf("expected content to have value %q", secondValue)
	}
}
