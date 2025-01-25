package primitive_object

import (
	"fmt"
	"strconv"
)

// ______________________________________
type PrimitiveUint struct {
	err      error
	value    int
	isNil    bool
	maxDigit *uint
	minDigit *uint
}

// ______________________________________
type PrimitiveUintOption func(*PrimitiveUint)

// ______________________________________
func (receiver *PrimitiveUint) WithError(
	err error,
) PrimitiveUintOption {
	return func(s *PrimitiveUint) {
		s.err = err
	}
}

// ______________________________________
func (receiver *PrimitiveUint) WithValue(
	value *int,
) PrimitiveUintOption {
	isNil := receiver.CheckNil(value)
	valueInt := 0
	if !isNil {
		valueInt = *value
	}

	return func(s *PrimitiveUint) {
		s.value = valueInt
		s.isNil = isNil
	}
}

// ______________________________________
func (receiver *PrimitiveUint) WithIsNil(
	isNil bool,
) PrimitiveUintOption {
	return func(s *PrimitiveUint) {
		s.isNil = isNil
	}
}

// ______________________________________
func (receiver *PrimitiveUint) WithMaxDigit(
	value *uint,
) PrimitiveUintOption {
	return func(s *PrimitiveUint) {
		s.maxDigit = value
	}
}

// ______________________________________
func (receiver *PrimitiveUint) WithMinDigit(
	value *uint,
) PrimitiveUintOption {
	return func(s *PrimitiveUint) {
		s.minDigit = value
	}
}

// ______________________________________
func NewPrimitiveUint(
	options ...PrimitiveUintOption,
) (
	primitiveUint *PrimitiveUint,
) {
	// デフォルト値を設定
	primitiveUint = &PrimitiveUint{
		err:      nil,
		value:    0,
		isNil:    true,
		maxDigit: nil,
		minDigit: nil,
	}

	// オプションを適用
	for _, option := range options {
		option(primitiveUint)
	}

	return
}

// ______________________________________
func (receiver *PrimitiveUint) GetIsNil() bool {
	return receiver.isNil
}

// ______________________________________
func (receiver *PrimitiveUint) SetIsNil(
	isNil bool,
) {
	receiver.isNil = isNil
}

// ______________________________________
func (receiver *PrimitiveUint) CheckNil(
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
func (receiver *PrimitiveUint) GetError() error {
	return receiver.err
}

// ______________________________________
func (receiver *PrimitiveUint) SetError(
	err error,
) {
	receiver.err = err
}

// ______________________________________
func (receiver *PrimitiveUint) SetErrorString(
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
func (receiver *PrimitiveUint) GetValue() int {
	if receiver.GetIsNil() {
		return 0
	}

	return receiver.value
}

// ______________________________________
func (receiver *PrimitiveUint) SetValue(
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
func (receiver *PrimitiveUint) Validation() {

	if receiver.GetIsNil() {
		return
	}

	receiver.ValidationMaxDigit()
	if receiver.err != nil {
		return
	}

	receiver.ValidationMinDigit()
	if receiver.err != nil {
		return
	}
}

// ______________________________________
func (receiver *PrimitiveUint) ValidationMaxDigit() {
	if receiver.maxDigit == nil { //上限値なし
		return
	}

	// 上限値ありでかつnilの場合エラーとする
	if receiver.GetIsNil() {
		// receiver.SetErrorString("is nil")
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
}

// ______________________________________
func (receiver *PrimitiveUint) ValidationMinDigit() {
	if receiver.minDigit == nil { // 下限値なし
		return
	}

	// 下限値ありでかつnilの場合エラーとする
	if receiver.GetIsNil() {
		// receiver.SetErrorString("is nil")
		return
	}

	strValue := strconv.Itoa(receiver.value)

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
