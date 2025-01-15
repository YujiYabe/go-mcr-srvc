package primitive_object

import (
	"fmt"
	"strconv"
)

// ______________________________________
type PrimitiveInt64 struct {
	err      error
	value    int64
	isNil    bool
	maxDigit *uint
	minDigit *uint
}

// ______________________________________
type PrimitiveInt64Option func(*PrimitiveInt64)

// ______________________________________
func (receiver *PrimitiveInt64) WithError(
	err error,
) PrimitiveInt64Option {
	return func(s *PrimitiveInt64) {
		s.err = err
	}
}

// ______________________________________
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
		err:      nil,
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
func (receiver *PrimitiveInt64) SetIsNil(
	isNil bool,
) {
	receiver.isNil = isNil
}

// ______________________________________
func (receiver *PrimitiveInt64) GetIsNil() bool {
	return receiver.isNil
}

// ______________________________________
func (receiver *PrimitiveInt64) GetError() error {
	return receiver.err
}

// ______________________________________
func (receiver *PrimitiveInt64) SetError(
	err error,
) {
	receiver.err = err
}

// ______________________________________
func (receiver *PrimitiveInt64) SetErrorString(
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
func (receiver *PrimitiveInt64) GetValue() int64 {
	if receiver.GetIsNil() {
		return 0
	}
	return receiver.value
}

// ______________________________________
func (receiver *PrimitiveInt64) GetString() string {
	if receiver.GetIsNil() {
		return ""
	}
	return fmt.Sprintf("%d", receiver.value)
}

// ______________________________________
func (receiver *PrimitiveInt64) SetValue(
	value *int64,
) {
	if value == nil {
		receiver.SetIsNil(true)
		return
	}
	receiver.SetIsNil(false)
	receiver.value = *value
}

// ______________________________________
func (receiver *PrimitiveInt64) Validation() {
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
func (receiver *PrimitiveInt64) ValidationMaxDigit() {
	if receiver.maxDigit == nil { //上限値なし
		return
	}

	// 上限値ありでかつnilの場合エラーとする
	if receiver.GetIsNil() {
		// receiver.SetErrorString("is nil")
		return
	}

	strValue := strconv.FormatInt(receiver.value, 10)

	// 桁数を取得
	digitCount := uint(len(strValue))
	debug := digitCount
	fmt.Println(" ----------------------------------- ")
	fmt.Printf("%+v\n", debug)

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
func (receiver *PrimitiveInt64) ValidationMinDigit() {
	if receiver.minDigit == nil { //上限値なし
		return
	}

	// 下限値ありでかつnilの場合エラーとする
	if receiver.GetIsNil() {
		// receiver.SetErrorString("is nil")
		return
	}

	strValue := strconv.FormatInt(receiver.value, 10)

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
