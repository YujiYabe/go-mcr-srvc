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
type PrimitiveUIntX[T IntX] struct {
	value    T
	isNil    bool
	maxDigit *uint
	minDigit *uint
}

// ______________________________________
type PrimitiveUIntXOption[T IntX] func(*PrimitiveUIntX[T])

// ______________________________________
func NewPrimitiveUIntX[T IntX](
	options ...func(*PrimitiveUIntX[T]),
) *PrimitiveUIntX[T] {
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
func (receiver PrimitiveUIntX[T]) GetValue() T {
	if receiver.GetIsNil() {
		return 0
	}
	return receiver.value
}

func (receiver PrimitiveUIntX[T]) Validation() error {
	if receiver.GetIsNil() {
		return nil
	}

	if err := receiver.ValidationMaxDigit(); err != nil {
		return err
	}

	return receiver.ValidationMinDigit()
}

func (receiver PrimitiveUIntX[T]) ValidationMaxDigit() error {
	if receiver.maxDigit == nil {
		return nil
	}

	if receiver.GetIsNil() {
		return nil
	}

	strValue := strconv.FormatInt(int64(receiver.value), 10)
	digitCount := uint(len(strValue))

	if receiver.value < 0 {
		digitCount--
	}

	if digitCount > *receiver.maxDigit {
		return receiver.newErrorString("max limitation")
	}

	return nil
}

// ______________________________________
func (receiver PrimitiveUIntX[T]) GetIsNil() bool {
	return receiver.isNil
}

// ______________________________________
func (receiver PrimitiveUIntX[T]) newErrorString(
	errString string,
) error {
	return fmt.Errorf(
		"error: %s",
		errString,
	)
}

// ______________________________________
func (receiver PrimitiveUIntX[T]) ValidationMinDigit() error {
	if receiver.minDigit == nil { // 下限値なし
		return nil
	}

	// 下限値ありでかつnilの場合エラーとする
	if receiver.GetIsNil() {
		// receiver.setErrorString("is nil")
		return nil
	}

	strValue := strconv.FormatInt(int64(receiver.value), 10)

	// 桁数を取得
	digitCount := uint(len(strValue))

	// 負の値の場合、マイナス記号を除いた桁数を計算
	if receiver.value < 0 {
		digitCount-- // マイナス符号を引く
	}

	if digitCount < *receiver.minDigit {
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
) PrimitiveUIntXOption[T] {
	return func(s *PrimitiveUIntX[T]) {
		s.isNil = isNil
	}
}

// ______________________________________
func (receiver *PrimitiveUIntX[T]) WithMaxDigit(
	value *uint,
) PrimitiveUIntXOption[T] {
	return func(s *PrimitiveUIntX[T]) {
		s.maxDigit = value
	}
}

// ______________________________________
func (receiver *PrimitiveUIntX[T]) WithMinDigit(
	value *uint,
) PrimitiveUIntXOption[T] {
	return func(s *PrimitiveUIntX[T]) {
		s.minDigit = value
	}
}

// ______________________________________
func (receiver *PrimitiveUIntX[T]) WithValue(
	value *T,
) PrimitiveUIntXOption[T] {
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
func (receiver PrimitiveUIntX[T]) GetString() string {
	if receiver.GetIsNil() {
		return ""
	}
	return fmt.Sprintf("%d", receiver.value)
}
