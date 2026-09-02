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

	return primitive
}

// 共通メソッド
func (receiver PrimitiveUIntX[T]) GetValue() (
	value T,
) {
	if receiver.GetIsNil() {
		return 0
	}
	return receiver.value
}

// ______________________________________
func (receiver PrimitiveUIntX[T]) IsZero() (
	isZero bool,
) {
	return receiver.GetValue() == 0
}

// ______________________________________
func (receiver PrimitiveUIntX[T]) HasValue() (
	hasValue bool,
) {
	return !receiver.GetIsNil()
}

// ______________________________________
func (receiver PrimitiveUIntX[T]) Equal(
	value T,
) (
	ok bool,
) {
	return !receiver.GetIsNil() && receiver.value == value
}

// ______________________________________
func (receiver PrimitiveUIntX[T]) DigitCount() (
	value uint,
) {
	if receiver.GetIsNil() {
		return 0
	}

	return uint(len(strconv.FormatUint(uint64(receiver.value), 10)))
}

func (receiver PrimitiveUIntX[T]) Validation() (
	err error,
) {
	if receiver.GetIsNil() {
		return nil
	}

	if err := receiver.ValidationMaxDigit(); err != nil {
		return err
	}

	return receiver.ValidationMinDigit()
}

func (receiver PrimitiveUIntX[T]) ValidationMaxDigit() (
	err error,
) {
	if receiver.maxDigit == nil {
		return nil
	}

	if receiver.GetIsNil() {
		return nil
	}

	if receiver.DigitCount() > *receiver.maxDigit {
		return receiver.newErrorString("max limitation")
	}

	return nil
}

// ______________________________________
func (receiver PrimitiveUIntX[T]) GetIsNil() (
	ok bool,
) {
	return receiver.isNil
}

// ______________________________________
func (receiver PrimitiveUIntX[T]) newErrorString(
	errString string,
) (
	err error,
) {
	return fmt.Errorf(
		"error: %s",
		errString,
	)
}

// ______________________________________
func (receiver PrimitiveUIntX[T]) ValidationMinDigit() (
	err error,
) {
	if receiver.minDigit == nil { // 下限値なし
		return nil
	}

	// 下限値ありでかつnilの場合エラーとする
	if receiver.GetIsNil() {
		// receiver.setErrorString("is nil")
		return nil
	}

	if receiver.DigitCount() < *receiver.minDigit {
		return receiver.newErrorString("min limitation")
	}

	return nil
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
	return func(s *PrimitiveUIntX[T]) {
		s.isNil = isNil
	}
}

// ______________________________________
func (receiver *PrimitiveUIntX[T]) WithMaxDigit(
	value *uint,
) (
	option PrimitiveUIntXOption[T],
) {
	return func(s *PrimitiveUIntX[T]) {
		s.maxDigit = value
	}
}

// ______________________________________
func (receiver *PrimitiveUIntX[T]) WithMinDigit(
	value *uint,
) (
	option PrimitiveUIntXOption[T],
) {
	return func(s *PrimitiveUIntX[T]) {
		s.minDigit = value
	}
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

	return func(s *PrimitiveUIntX[T]) {
		s.value = valueIntX
		s.isNil = isNil
	}
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
	if receiver.GetIsNil() {
		return
	}
	value = strconv.FormatUint(uint64(receiver.value), 10)
	return
}
