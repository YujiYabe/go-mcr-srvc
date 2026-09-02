package primitive_object

import (
	"math"
	"strconv"
	"testing"
)

func TestPrimitiveIntXToUint32(
	t *testing.T,
) {
	tests := []struct {
		name          string
		value         int
		expectedValue uint32
		expectError   bool
	}{
		{
			name:          "minimum value",
			value:         0,
			expectedValue: 0,
		},
		{
			name:          "maximum value",
			value:         int(math.MaxUint32),
			expectedValue: math.MaxUint32,
		},
		{
			name:        "negative value",
			value:       -1,
			expectError: true,
		},
	}

	if strconv.IntSize > 32 {
		tests = append(
			tests,
			struct {
				name          string
				value         int
				expectedValue uint32
				expectError   bool
			}{
				name:        "value above maximum",
				value:       int(uint64(math.MaxUint32) + 1),
				expectError: true,
			},
		)
	}

	for _, test := range tests {
		t.Run(
			test.name,
			func(t *testing.T) {
				primitiveIntX := &PrimitiveIntX[int]{}
				value := test.value
				primitive := NewPrimitiveIntX(
					primitiveIntX.WithValue(&value),
				)
				convertedValue, err := primitive.ToUint32()
				if test.expectError {
					if err == nil {
						t.Fatal("expected an error")
					}
					return
				}

				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
				if convertedValue != test.expectedValue {
					t.Fatalf(
						"expected %d, got %d",
						test.expectedValue,
						convertedValue,
					)
				}
			},
		)
	}
}

func TestPrimitiveIntXToSignedIntegerTypes(
	t *testing.T,
) {
	primitive := newPrimitiveInt64ForConversion(math.MaxInt32 + 1)

	int64Value, err := primitive.ToInt64()
	if err != nil {
		t.Fatalf("ToInt64 returned an unexpected error: %v", err)
	}
	if int64Value != math.MaxInt32+1 {
		t.Fatalf("ToInt64 expected %d, got %d", int64(math.MaxInt32+1), int64Value)
	}

	_, err = primitive.ToInt32()
	if err == nil {
		t.Fatal("ToInt32 expected an overflow error")
	}

	intValue, err := newPrimitiveInt64ForConversion(1).ToInt()
	if err != nil {
		t.Fatalf("ToInt returned an unexpected error: %v", err)
	}
	if intValue != 1 {
		t.Fatalf("ToInt expected 1, got %d", intValue)
	}
}

func TestPrimitiveIntXToUnsignedIntegerTypes(
	t *testing.T,
) {
	primitive := newPrimitiveInt64ForConversion(math.MaxUint32)

	uintValue, err := primitive.ToUint()
	if err != nil {
		t.Fatalf("ToUint returned an unexpected error: %v", err)
	}
	if uint64(uintValue) != math.MaxUint32 {
		t.Fatalf("ToUint expected %d, got %d", uint64(math.MaxUint32), uintValue)
	}

	uint64Value, err := primitive.ToUint64()
	if err != nil {
		t.Fatalf("ToUint64 returned an unexpected error: %v", err)
	}
	if uint64Value != math.MaxUint32 {
		t.Fatalf("ToUint64 expected %d, got %d", uint64(math.MaxUint32), uint64Value)
	}

	negativePrimitive := newPrimitiveInt64ForConversion(-1)
	if _, err = negativePrimitive.ToUint(); err == nil {
		t.Fatal("ToUint expected an error for a negative value")
	}
	if _, err = negativePrimitive.ToUint64(); err == nil {
		t.Fatal("ToUint64 expected an error for a negative value")
	}
}

func newPrimitiveInt64ForConversion(
	value int64,
) (
	primitive *PrimitiveIntX[int64],
) {
	primitiveIntX := &PrimitiveIntX[int64]{}
	primitive = NewPrimitiveIntX(
		primitiveIntX.WithValue(&value),
	)

	return
}
