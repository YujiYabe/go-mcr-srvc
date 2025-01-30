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
	err      error
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
		err:      nil,
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
func (receiver *PrimitiveIntX[T]) GetValue() T {
	if receiver.GetIsNil() {
		return 0
	}
	return receiver.value
}

func (receiver *PrimitiveIntX[T]) Validation() {
	if receiver.GetIsNil() {
		return
	}

	receiver.ValidationMaxDigit()
	if receiver.err != nil {
		return
	}

	receiver.ValidationMinDigit()
}

func (receiver *PrimitiveIntX[T]) ValidationMaxDigit() {
	if receiver.maxDigit == nil {
		return
	}

	if receiver.GetIsNil() {
		return
	}

	strValue := strconv.FormatInt(int64(receiver.value), 10)
	digitCount := uint(len(strValue))

	if receiver.value < 0 {
		digitCount--
	}

	if digitCount > *receiver.maxDigit {
		receiver.SetErrorString("max limitation")
	}
}

// ______________________________________
func (receiver *PrimitiveIntX[T]) GetIsNil() bool {
	return receiver.isNil
}

// ______________________________________
func (receiver *PrimitiveIntX[T]) SetErrorString(
	errString string,
) {
	receiver.SetError(
		fmt.Errorf(
			"error: %s",
			errString,
		),
	)
}

// ______________________________________
func (receiver *PrimitiveIntX[T]) SetError(
	err error,
) {
	receiver.err = err
}

// ______________________________________
func (receiver *PrimitiveIntX[T]) ValidationMinDigit() {
	if receiver.minDigit == nil { // 下限値なし
		return
	}

	// 下限値ありでかつnilの場合エラーとする
	if receiver.GetIsNil() {
		// receiver.SetErrorString("is nil")
		return
	}

	strValue := strconv.FormatInt(int64(receiver.value), 10)

	// 桁数を取得
	digitCount := uint(len(strValue))

	// 負の値の場合、マイナス記号を除いた桁数を計算
	if receiver.value < 0 {
		digitCount-- // マイナス符号を引く
	}

	if digitCount < *receiver.minDigit {
		receiver.SetErrorString("min limitation")
		return
	}
}

// ______________________________________
func (receiver *PrimitiveIntX[T]) CheckNil(
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
func (receiver *PrimitiveIntX[T]) GetError() error {
	return receiver.err
}

// ______________________________________
func (receiver *PrimitiveIntX[T]) SetIsNil(
	isNil bool,
) {
	receiver.isNil = isNil
}

// ______________________________________
func (receiver *PrimitiveIntX[T]) SetValue(
	value *T,
) {
	if value == nil {
		receiver.SetIsNil(true)
		return
	}
	receiver.SetIsNil(false)
	receiver.value = *value
}

// ______________________________________
func (receiver *PrimitiveIntX[T]) WithError(
	err error,
) PrimitiveIntXOption[T] {
	return func(s *PrimitiveIntX[T]) {
		s.err = err
	}
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
	var valueIntX T = 0
	if !isNil {
		valueIntX = *value
	}

	return func(s *PrimitiveIntX[T]) {
		s.value = valueIntX
		s.isNil = isNil
	}
}

// ______________________________________
func (receiver *PrimitiveIntX[T]) GetString() string {
	if receiver.GetIsNil() {
		return ""
	}
	return fmt.Sprintf("%d", receiver.value)
}
