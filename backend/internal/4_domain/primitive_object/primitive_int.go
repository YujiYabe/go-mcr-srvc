package primitive_object

import (
	"fmt"
	"strconv"
)

// ______________________________________
type PrimitiveInt struct {
	value    int
	isNil    bool
	maxDigit *uint
	minDigit *uint
}

// ______________________________________
type PrimitiveIntOption func(*PrimitiveInt)

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
		value:    0,
		isNil:    true,
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
func (receiver PrimitiveInt) GetIsNil() bool {
	return receiver.isNil
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

func (receiver PrimitiveInt) newErrorString(
	errString string,
) error {
	return fmt.Errorf(
		"error: %s",
		errString,
	)
}

// ______________________________________
func (receiver PrimitiveInt) GetValue() int {
	if receiver.GetIsNil() {
		return 0
	}

	return receiver.value
}

// ______________________________________
func (receiver PrimitiveInt) Validation() error {
	if receiver.GetIsNil() {
		return nil
	}

	if err := receiver.ValidationMaxDigit(); err != nil {
		return err
	}

	return receiver.ValidationMinDigit()
}

// ______________________________________
func (receiver PrimitiveInt) ValidationMaxDigit() error {
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
func (receiver PrimitiveInt) ValidationMinDigit() error {
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
