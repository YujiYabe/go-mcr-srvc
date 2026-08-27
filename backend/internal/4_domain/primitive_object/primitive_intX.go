package primitive_object

import (
	"fmt"
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
) *PrimitiveIntX[T] {
	primitive := &PrimitiveIntX[T]{
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
func (receiver PrimitiveIntX[T]) GetValue() T {
	if receiver.GetIsNil() {
		return 0
	}
	return receiver.value
}

// ______________________________________
func (receiver PrimitiveIntX[T]) IsZero() bool {
	return receiver.GetValue() == 0
}

// ______________________________________
func (receiver PrimitiveIntX[T]) HasValue() bool {
	return !receiver.GetIsNil()
}

// ______________________________________
func (receiver PrimitiveIntX[T]) Equal(value T) bool {
	return !receiver.GetIsNil() && receiver.value == value
}

// ______________________________________
func (receiver PrimitiveIntX[T]) DigitCount() uint {
	if receiver.GetIsNil() {
		return 0
	}

	strValue := strconv.FormatInt(int64(receiver.value), 10)
	digitCount := uint(len(strValue))
	if receiver.value < 0 {
		digitCount--
	}

	return digitCount
}

func (receiver PrimitiveIntX[T]) Validation() error {
	if receiver.GetIsNil() {
		return nil
	}

	if err := receiver.ValidationMaxDigit(); err != nil {
		return err
	}

	return receiver.ValidationMinDigit()
}

func (receiver PrimitiveIntX[T]) ValidationMaxDigit() error {
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
func (receiver PrimitiveIntX[T]) GetIsNil() bool {
	return receiver.isNil
}

// ______________________________________
func (receiver PrimitiveIntX[T]) newErrorString(
	errString string,
) error {
	return fmt.Errorf(
		"error: %s",
		errString,
	)
}

// ______________________________________
func (receiver PrimitiveIntX[T]) ValidationMinDigit() error {
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
) PrimitiveIntXOption[T] {
	return func(s *PrimitiveIntX[T]) {
		s.isNil = isNil
	}
}

// ______________________________________
func (receiver *PrimitiveIntX[T]) WithMaxDigit(
	value *uint,
) PrimitiveIntXOption[T] {
	return func(s *PrimitiveIntX[T]) {
		s.maxDigit = value
	}
}

// ______________________________________
func (receiver *PrimitiveIntX[T]) WithMinDigit(
	value *uint,
) PrimitiveIntXOption[T] {
	return func(s *PrimitiveIntX[T]) {
		s.minDigit = value
	}
}

// ______________________________________
func (receiver *PrimitiveIntX[T]) WithValue(
	value *T,
) PrimitiveIntXOption[T] {
	isNil := receiver.CheckNil(value)
	var valueIntX T
	if !isNil {
		valueIntX = *value
	}

	return func(s *PrimitiveIntX[T]) {
		s.value = valueIntX
		s.isNil = isNil
	}
}

// ______________________________________
func (receiver PrimitiveIntX[T]) GetString() string {
	if receiver.GetIsNil() {
		return ""
	}
	return fmt.Sprintf("%d", receiver.value)
}
