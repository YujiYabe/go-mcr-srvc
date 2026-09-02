package primitive_object

import (
	"math"
	"testing"
)

func TestPrimitiveStringToSignedIntegerTypes(
	t *testing.T,
) {
	primitive := newPrimitiveStringForConversion("2147483647")

	intValue, err := primitive.ToInt()
	if err != nil || int64(intValue) != math.MaxInt32 {
		t.Fatalf("ToInt expected %d, got %d, err: %v", math.MaxInt32, intValue, err)
	}
	int32Value, err := primitive.ToInt32()
	if err != nil || int32Value != math.MaxInt32 {
		t.Fatalf("ToInt32 expected %d, got %d, err: %v", math.MaxInt32, int32Value, err)
	}
	int64Value, err := primitive.ToInt64()
	if err != nil || int64Value != math.MaxInt32 {
		t.Fatalf("ToInt64 expected %d, got %d, err: %v", math.MaxInt32, int64Value, err)
	}

	if _, err = newPrimitiveStringForConversion("2147483648").ToInt32(); err == nil {
		t.Fatal("ToInt32 expected an overflow error")
	}
	if _, err = newPrimitiveStringForConversion("not-a-number").ToInt64(); err == nil {
		t.Fatal("ToInt64 expected a syntax error")
	}
}

func TestPrimitiveStringToUnsignedIntegerTypes(
	t *testing.T,
) {
	primitive := newPrimitiveStringForConversion("4294967295")

	uintValue, err := primitive.ToUint()
	if err != nil || uint64(uintValue) != math.MaxUint32 {
		t.Fatalf("ToUint expected %d, got %d, err: %v", uint64(math.MaxUint32), uintValue, err)
	}
	uint32Value, err := primitive.ToUint32()
	if err != nil || uint32Value != math.MaxUint32 {
		t.Fatalf("ToUint32 expected %d, got %d, err: %v", uint64(math.MaxUint32), uint32Value, err)
	}
	uint64Value, err := primitive.ToUint64()
	if err != nil || uint64Value != math.MaxUint32 {
		t.Fatalf("ToUint64 expected %d, got %d, err: %v", uint64(math.MaxUint32), uint64Value, err)
	}

	if _, err = newPrimitiveStringForConversion("4294967296").ToUint32(); err == nil {
		t.Fatal("ToUint32 expected an overflow error")
	}
	if _, err = newPrimitiveStringForConversion("-1").ToUint64(); err == nil {
		t.Fatal("ToUint64 expected an error for a negative value")
	}
}

func TestPrimitiveIntegerTypesToString(
	t *testing.T,
) {
	intValue := int64(-123)
	primitiveIntX := &PrimitiveIntX[int64]{}
	intPrimitive := NewPrimitiveIntX(primitiveIntX.WithValue(&intValue))
	if value := intPrimitive.ToString(); value != "-123" {
		t.Fatalf("ToString expected %q, got %q", "-123", value)
	}

	uintValue := uint64(123)
	primitiveUintX := &PrimitiveUIntX[uint64]{}
	uintPrimitive := NewPrimitiveUIntX(primitiveUintX.WithValue(&uintValue))
	if value := uintPrimitive.ToString(); value != "123" {
		t.Fatalf("ToString expected %q, got %q", "123", value)
	}
}

func newPrimitiveStringForConversion(
	value string,
) (
	primitive *PrimitiveString,
) {
	primitiveString := &PrimitiveString{}
	primitive = NewPrimitiveString(
		primitiveString.WithValue(&value),
	)
	return
}
