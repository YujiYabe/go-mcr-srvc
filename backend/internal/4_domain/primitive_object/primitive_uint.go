package primitive_object

import (
	"fmt"
	"strconv"
)

// ______________________________________
type PrimitiveUint struct {
	value    int
	isNil    bool
	maxDigit *uint
	minDigit *uint
}

// ______________________________________
type PrimitiveUintOption func(*PrimitiveUint)

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
func (receiver PrimitiveUint) GetIsNil() bool {
	return receiver.isNil
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

func (receiver PrimitiveUint) newErrorString(
	errString string,
) error {
	return fmt.Errorf(
		"error: %s",
		errString,
	)
}

// ______________________________________
func (receiver PrimitiveUint) GetValue() int {
	if receiver.GetIsNil() {
		return 0
	}

	return receiver.value
}

// ______________________________________
func (receiver PrimitiveUint) Validation() error {
	if receiver.GetIsNil() {
		return nil
	}

	if err := receiver.ValidationMaxDigit(); err != nil {
		return err
	}

	return receiver.ValidationMinDigit()
}

// ______________________________________
func (receiver PrimitiveUint) ValidationMaxDigit() error {
	if receiver.maxDigit == nil { // 上限値なし
		return nil
	}

	// 上限値ありでかつnilの場合エラーとする
	if receiver.GetIsNil() {
		// receiver.setErrorString("is nil")
		return nil
	}

	strValue := strconv.Itoa(receiver.value)

	// 桁数を取得
	digitCount := uint(len(strValue))

	// 負の値の場合、マイナス記号を除いた桁数を計算
	if receiver.value < 0 {
		digitCount-- // マイナス符号を引く
	}

	if digitCount > *receiver.maxDigit {
		return receiver.newErrorString("max limitation")
	}

	return nil
}

// ______________________________________
func (receiver PrimitiveUint) ValidationMinDigit() error {
	if receiver.minDigit == nil { // 下限値なし
		return nil
	}

	// 下限値ありでかつnilの場合エラーとする
	if receiver.GetIsNil() {
		// receiver.setErrorString("is nil")
		return nil
	}

	strValue := strconv.Itoa(receiver.value)

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
