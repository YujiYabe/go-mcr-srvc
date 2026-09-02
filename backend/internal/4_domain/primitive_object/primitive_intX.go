package primitive_object

import (
	"fmt"
	"math"
	"strconv"
)

// 数値型のインターフェース制約
// ______________________________________
type IntX interface {
	~int | ~int32 | ~int64
}

// ______________________________________
type PrimitiveIntX[T IntX] struct {
	value    T
	isNil    bool
	maxDigit *uint
	minDigit *uint
}

// ______________________________________
type PrimitiveIntXOption[T IntX] func(*PrimitiveIntX[T])

// ______________________________________
func NewPrimitiveIntX[T IntX](
	options ...func(*PrimitiveIntX[T]),
) (
	primitiveIntX *PrimitiveIntX[T],
) {
	primitive := &PrimitiveIntX[T]{
		value:    0,
		isNil:    true,
		maxDigit: nil,
		minDigit: nil,
	}

	for _, option := range options {
		option(primitive)
	}

	primitiveIntX = primitive

	return
}

// 共通メソッド
func (receiver PrimitiveIntX[T]) GetValue() (
	value T,
) {
	if receiver.GetIsNil() {
		value = 0

		return
	}
	value = receiver.value

	return
}

func (receiver PrimitiveIntX[T]) ToInt() (
	value int,
	err error,
) {
	value = 0
	err = nil
	intValue := int64(receiver.GetValue())
	if strconv.IntSize == 32 && (intValue < math.MinInt32 || intValue > math.MaxInt32) {
		err = newIntegerConversionError("int", intValue)

		return
	}
	value = int(intValue)

	return
}

func (receiver PrimitiveIntX[T]) ToInt32() (
	value int32,
	err error,
) {
	value = 0
	err = nil
	intValue := int64(receiver.GetValue())
	if intValue < math.MinInt32 || intValue > math.MaxInt32 {
		err = newIntegerConversionError("int32", intValue)

		return
	}
	value = int32(intValue)

	return
}

func (receiver PrimitiveIntX[T]) ToInt64() (
	value int64,
	err error,
) {
	err = nil
	value = int64(receiver.GetValue())

	return
}

func (receiver PrimitiveIntX[T]) ToUint() (
	value uint,
	err error,
) {
	value = 0
	err = nil
	intValue := int64(receiver.GetValue())
	if intValue < 0 || (strconv.IntSize == 32 && uint64(intValue) > math.MaxUint32) {
		err = newIntegerConversionError("uint", intValue)

		return
	}
	value = uint(intValue)

	return
}

func (receiver PrimitiveIntX[T]) ToUint32() (
	value uint32,
	err error,
) {
	value = 0
	err = nil
	intValue := int64(receiver.GetValue())
	if intValue < 0 || uint64(intValue) > math.MaxUint32 {
		err = newIntegerConversionError("uint32", intValue)

		return
	}
	value = uint32(intValue)

	return
}

func (receiver PrimitiveIntX[T]) ToUint64() (
	value uint64,
	err error,
) {
	value = 0
	err = nil
	intValue := int64(receiver.GetValue())
	if intValue < 0 {
		err = newIntegerConversionError("uint64", intValue)

		return
	}
	value = uint64(intValue)

	return
}

func newIntegerConversionError(
	targetType string,
	value int64,
) (
	err error,
) {
	err = fmt.Errorf("value is outside the %s range: %d", targetType, value)

	return
}

// ______________________________________
func (receiver PrimitiveIntX[T]) IsZero() (
	isZero bool,
) {
	isZero = receiver.GetValue() == 0

	return
}

// ______________________________________
func (receiver PrimitiveIntX[T]) HasValue() (
	hasValue bool,
) {
	hasValue = !receiver.GetIsNil()

	return
}

// ______________________________________
func (receiver PrimitiveIntX[T]) Equal(
	value T,
) (
	ok bool,
) {
	ok = !receiver.GetIsNil() && receiver.value == value

	return
}

// ______________________________________
func (receiver PrimitiveIntX[T]) DigitCount() (
	value uint,
) {
	if receiver.GetIsNil() {
		value = 0

		return
	}

	strValue := strconv.FormatInt(int64(receiver.value), 10)
	digitCount := uint(len(strValue))
	if receiver.value < 0 {
		digitCount--
	}

	value = digitCount

	return
}

func (receiver PrimitiveIntX[T]) Validation() (
	err error,
) {
	if receiver.GetIsNil() {
		err = nil

		return
	}

	if returnedErr := receiver.ValidationMaxDigit(); returnedErr != nil {
		err = returnedErr

		return
	}

	err = receiver.ValidationMinDigit()

	return
}

func (receiver PrimitiveIntX[T]) ValidationMaxDigit() (
	err error,
) {
	if receiver.maxDigit == nil {
		err = nil

		return
	}

	if receiver.GetIsNil() {
		err = nil

		return
	}

	if receiver.DigitCount() > *receiver.maxDigit {
		err = receiver.newErrorString("max limitation")

		return
	}

	err = nil

	return
}

// ______________________________________
func (receiver PrimitiveIntX[T]) GetIsNil() (
	ok bool,
) {
	ok = receiver.isNil

	return
}

// ______________________________________
func (receiver PrimitiveIntX[T]) newErrorString(
	errString string,
) (
	err error,
) {
	err = fmt.Errorf(
		"error: %s",
		errString,
	)

	return
}

// ______________________________________
func (receiver PrimitiveIntX[T]) ValidationMinDigit() (
	err error,
) {
	if receiver.minDigit == nil { // 下限値なし
		err = nil

		return
	}

	// 下限値ありでかつnilの場合エラーとする
	if receiver.GetIsNil() {
		// receiver.setErrorString("is nil")
		err = nil

		return
	}

	if receiver.DigitCount() < *receiver.minDigit {
		err = receiver.newErrorString("min limitation")

		return
	}

	err = nil

	return
}

// ______________________________________
func (receiver PrimitiveIntX[T]) CheckNil(
	value *T,
) (
	isNil bool,
) {
	isNil = true
	if value != nil {
		isNil = false
	}

	return
}

// ______________________________________
func (receiver *PrimitiveIntX[T]) WithIsNil(
	isNil bool,
) (
	value PrimitiveIntXOption[T],
) {
	value = func(s *PrimitiveIntX[T]) {
		s.isNil = isNil
	}

	return
}

// ______________________________________
func (receiver *PrimitiveIntX[T]) WithMaxDigit(
	value *uint,
) (
	option PrimitiveIntXOption[T],
) {
	option = func(s *PrimitiveIntX[T]) {
		s.maxDigit = value
	}

	return
}

// ______________________________________
func (receiver *PrimitiveIntX[T]) WithMinDigit(
	value *uint,
) (
	option PrimitiveIntXOption[T],
) {
	option = func(s *PrimitiveIntX[T]) {
		s.minDigit = value
	}

	return
}

// ______________________________________
func (receiver *PrimitiveIntX[T]) WithValue(
	value *T,
) (
	option PrimitiveIntXOption[T],
) {
	isNil := receiver.CheckNil(value)
	var valueIntX T
	if !isNil {
		valueIntX = *value
	}

	option = func(s *PrimitiveIntX[T]) {
		s.value = valueIntX
		s.isNil = isNil
	}

	return
}

// ______________________________________
func (receiver PrimitiveIntX[T]) GetString() (
	value string,
) {
	value = receiver.ToString()

	return
}

func (receiver PrimitiveIntX[T]) ToString() (
	value string,
) {
	value = ""
	if receiver.GetIsNil() {

		return
	}
	value = strconv.FormatInt(int64(receiver.value), 10)

	return
}
