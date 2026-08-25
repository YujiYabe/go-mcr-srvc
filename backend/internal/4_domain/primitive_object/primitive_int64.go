package primitive_object

import (
	"fmt"
	"strconv"
)

// ______________________________________
type PrimitiveInt64 struct {
	value    int64
	isNil    bool
	maxDigit *uint
	minDigit *uint
}

// ______________________________________
type PrimitiveInt64Option func(*PrimitiveInt64)

func (receiver *PrimitiveInt64) WithValue(
	value *int64,
) PrimitiveInt64Option {
	isNil := receiver.CheckNil(value)
	valueInt64 := int64(0)
	if !isNil {
		valueInt64 = *value
	}

	return func(s *PrimitiveInt64) {
		s.value = valueInt64
		s.isNil = isNil
	}
}

// ______________________________________
func (receiver *PrimitiveInt64) WithIsNil(
	isNil bool,
) PrimitiveInt64Option {
	return func(s *PrimitiveInt64) {
		s.isNil = isNil
	}
}

// ______________________________________
func (receiver *PrimitiveInt64) WithMaxDigit(
	value *uint,
) PrimitiveInt64Option {
	return func(s *PrimitiveInt64) {
		s.maxDigit = value
	}
}

// ______________________________________
func (receiver *PrimitiveInt64) WithMinDigit(
	value *uint,
) PrimitiveInt64Option {
	return func(s *PrimitiveInt64) {
		s.minDigit = value
	}
}

// ______________________________________
func NewPrimitiveInt64(
	options ...PrimitiveInt64Option,
) (
	primitiveInt64 *PrimitiveInt64,
) {
	// デフォルト値を設定
	primitiveInt64 = &PrimitiveInt64{
		value:    0,
		isNil:    true,
		maxDigit: nil,
		minDigit: nil,
	}

	// オプションを適用
	for _, option := range options {
		option(primitiveInt64)
	}

	return
}

// ______________________________________
func (receiver PrimitiveInt64) GetIsNil() bool {
	return receiver.isNil
}

func (receiver PrimitiveInt64) newErrorString(
	errString string,
) error {
	return fmt.Errorf(
		"error: %s",
		errString,
	)
}

// ______________________________________
func (receiver PrimitiveInt64) GetValue() int64 {
	if receiver.GetIsNil() {
		return 0
	}
	return receiver.value
}

// ______________________________________
func (receiver PrimitiveInt64) GetString() string {
	if receiver.GetIsNil() {
		return ""
	}
	return fmt.Sprintf("%d", receiver.value)
}

// ______________________________________
func (receiver PrimitiveInt64) Validation() error {
	if receiver.GetIsNil() {
		return nil
	}

	if err := receiver.ValidationMaxDigit(); err != nil {
		return err
	}

	return receiver.ValidationMinDigit()
}

// ______________________________________
func (receiver PrimitiveInt64) ValidationMaxDigit() error {
	if receiver.maxDigit == nil { //上限値なし
		return nil
	}

	// 上限値ありでかつnilの場合エラーとする
	if receiver.GetIsNil() {
		// receiver.setErrorString("is nil")
		return nil
	}

	strValue := strconv.FormatInt(receiver.value, 10)

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
func (receiver PrimitiveInt64) ValidationMinDigit() error {
	if receiver.minDigit == nil { //上限値なし
		return nil
	}

	// 下限値ありでかつnilの場合エラーとする
	if receiver.GetIsNil() {
		// receiver.setErrorString("is nil")
		return nil
	}

	strValue := strconv.FormatInt(receiver.value, 10)

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
func (receiver *PrimitiveInt64) CheckNil(
	value *int64,
) (
	isNil bool,
) {
	isNil = true
	if value != nil {
		isNil = false
	}

	return
}
