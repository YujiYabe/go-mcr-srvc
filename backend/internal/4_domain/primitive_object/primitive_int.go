package primitive_object

import (
	"fmt"
	"strconv"
)

// ______________________________________
type PrimitiveInt struct {
	err      error
	value    int
	isNil    bool
	maxDigit *uint
	minDigit *uint
}

// ______________________________________
type PrimitiveIntOption func(*PrimitiveInt)

// ______________________________________
func (receiver *PrimitiveInt) WithError(
	err error,
) PrimitiveIntOption {
	return func(s *PrimitiveInt) {
		s.err = err
	}
}

// ______________________________________
func (receiver *PrimitiveInt) WithValue(
	value *int,
) PrimitiveIntOption {
	isNil := receiver.CheckNil(value)
	valueInt := 0
	if !isNil {
		valueInt = *value
	}

	return func(s *PrimitiveInt) {
		s.value = valueInt
		s.isNil = isNil
	}
}

// ______________________________________
func (receiver *PrimitiveInt) WithIsNil(
	isNil bool,
) PrimitiveIntOption {
	return func(s *PrimitiveInt) {
		s.isNil = isNil
	}
}

// ______________________________________
func (receiver *PrimitiveInt) WithMaxDigit(
	value *uint,
) PrimitiveIntOption {
	return func(s *PrimitiveInt) {
		s.maxDigit = value
	}
}

// ______________________________________
func (receiver *PrimitiveInt) WithMinDigit(
	value *uint,
) PrimitiveIntOption {
	return func(s *PrimitiveInt) {
		s.minDigit = value
	}
}

// ______________________________________
func NewPrimitiveInt(
	options ...PrimitiveIntOption,
) (
	primitiveInt *PrimitiveInt,
) {
	// デフォルト値を設定
	primitiveInt = &PrimitiveInt{
		err:      nil,
		value:    0,
		isNil:    false,
		maxDigit: nil,
		minDigit: nil,
	}

	// オプションを適用
	for _, option := range options {
		option(primitiveInt)
	}

	return
}

// ______________________________________
func (receiver *PrimitiveInt) GetIsNil() bool {
	return receiver.isNil
}

// ______________________________________
func (receiver *PrimitiveInt) SetIsNil(
	isNil bool,
) {
	receiver.isNil = isNil
}

// ______________________________________
func (receiver *PrimitiveInt) CheckNil(
	value *int,
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
func (receiver *PrimitiveInt) GetError() error {
	return receiver.err
}

// ______________________________________
func (receiver *PrimitiveInt) SetError(
	err error,
) {
	receiver.err = err
}

// ______________________________________
func (receiver *PrimitiveInt) SetErrorString(
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
func (receiver *PrimitiveInt) GetValue() int {
	if receiver.GetIsNil() {
		return 0
	}

	return receiver.value
}

// ______________________________________
func (receiver *PrimitiveInt) SetValue(
	value *int,
) {
	if value == nil {
		receiver.SetIsNil(true)
		return
	}
	receiver.SetIsNil(false)
	receiver.value = *value
}

// ______________________________________
func (receiver *PrimitiveInt) Validation() {

	if receiver.GetIsNil() {
		return
	}

	receiver.ValidationDigit()
	if receiver.err != nil {
		return
	}
}

// ______________________________________
func (receiver *PrimitiveInt) ValidationDigit() {
	if receiver.maxDigit == nil { //上限値なし
		return
	}

	// 上限値ありでかつnilの場合エラーとする
	if receiver.GetIsNil() {
		receiver.SetErrorString("is nil")
		return
	}

	strValue := strconv.Itoa(receiver.value)

	// 桁数を取得
	digitCount := uint(len(strValue))

	// 負の値の場合、マイナス記号を除いた桁数を計算
	if receiver.value < 0 {
		digitCount-- // マイナス符号を引く
	}

	if digitCount > *receiver.maxDigit {
		receiver.SetErrorString("max limitation")
		return
	}

	if digitCount < *receiver.minDigit {
		receiver.SetErrorString("min limitation")
		return
	}

}
