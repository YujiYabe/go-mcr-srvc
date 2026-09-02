package primitive_object

import (
	"fmt"
	"strconv"
)

// 数値型のインターフェース制約
// ______________________________________
type UIntX interface {
	~uint | ~uint32 | ~uint64
}

// ______________________________________
type PrimitiveUIntX[T UIntX] struct {
	value    T
	isNil    bool
	maxDigit *uint
	minDigit *uint
}

// ______________________________________
type PrimitiveUIntXOption[T UIntX] func(*PrimitiveUIntX[T])

// ______________________________________
func NewPrimitiveUIntX[T UIntX](
	options ...func(*PrimitiveUIntX[T]),
) (
	primitiveUIntX *PrimitiveUIntX[T],
) {
	primitive := &PrimitiveUIntX[T]{
		value:    0,
		isNil:    true,
		maxDigit: nil,
		minDigit: nil,
	}

	for _, option := range options {
		option(primitive)
	}

	primitiveUIntX = primitive
	return
}

// 共通メソッド
func (receiver PrimitiveUIntX[T]) GetValue() (
	value T,
) {
	if receiver.GetIsNil() {
		value = 0
		return
	}
	value = receiver.value
	return
}

// ______________________________________
func (receiver PrimitiveUIntX[T]) IsZero() (
	isZero bool,
) {
	isZero = receiver.GetValue() == 0
	return
}

// ______________________________________
func (receiver PrimitiveUIntX[T]) HasValue() (
	hasValue bool,
) {
	hasValue = !receiver.GetIsNil()
	return
}

// ______________________________________
func (receiver PrimitiveUIntX[T]) Equal(
	value T,
) (
	ok bool,
) {
	ok = !receiver.GetIsNil() && receiver.value == value
	return
}

// ______________________________________
func (receiver PrimitiveUIntX[T]) DigitCount() (
	value uint,
) {
	if receiver.GetIsNil() {
		value = 0
		return
	}

	value = uint(len(strconv.FormatUint(uint64(receiver.value), 10)))
	return
}

func (receiver PrimitiveUIntX[T]) Validation() (
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

func (receiver PrimitiveUIntX[T]) ValidationMaxDigit() (
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
func (receiver PrimitiveUIntX[T]) GetIsNil() (
	ok bool,
) {
	ok = receiver.isNil
	return
}

// ______________________________________
func (receiver PrimitiveUIntX[T]) newErrorString(
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
func (receiver PrimitiveUIntX[T]) ValidationMinDigit() (
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
func (receiver *PrimitiveUIntX[T]) CheckNil(
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

func (receiver *PrimitiveUIntX[T]) WithIsNil(
	isNil bool,
) (
	value PrimitiveUIntXOption[T],
) {
	value = func(s *PrimitiveUIntX[T]) {
		s.isNil = isNil
	}
	return
}

// ______________________________________
func (receiver *PrimitiveUIntX[T]) WithMaxDigit(
	value *uint,
) (
	option PrimitiveUIntXOption[T],
) {
	option = func(s *PrimitiveUIntX[T]) {
		s.maxDigit = value
	}
	return
}

// ______________________________________
func (receiver *PrimitiveUIntX[T]) WithMinDigit(
	value *uint,
) (
	option PrimitiveUIntXOption[T],
) {
	option = func(s *PrimitiveUIntX[T]) {
		s.minDigit = value
	}
	return
}

// ______________________________________
func (receiver *PrimitiveUIntX[T]) WithValue(
	value *T,
) (
	option PrimitiveUIntXOption[T],
) {
	isNil := receiver.CheckNil(value)
	var valueIntX T
	if !isNil {
		valueIntX = *value
	}

	option = func(s *PrimitiveUIntX[T]) {
		s.value = valueIntX
		s.isNil = isNil
	}
	return
}

// ______________________________________
func (receiver PrimitiveUIntX[T]) GetString() (
	value string,
) {
	value = receiver.ToString()
	return
}

func (receiver PrimitiveUIntX[T]) ToString() (
	value string,
) {
	value = ""
	if receiver.GetIsNil() {
		return
	}
	value = strconv.FormatUint(uint64(receiver.value), 10)
	return
}
