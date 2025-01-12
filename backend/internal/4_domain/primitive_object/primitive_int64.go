package primitive_object

import (
	"fmt"
)

// ______________________________________
type PrimitiveInt64 struct {
	err      error
	value    int64
	isNil    bool
	maxValue int64
	minValue int64
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
func (receiver *PrimitiveInt64) WithMaxValue(
	value int64,
) PrimitiveInt64Option {
	return func(s *PrimitiveInt64) {
		s.maxValue = value
	}
}

// ______________________________________
func (receiver *PrimitiveInt64) WithMinValue(
	value int64,
) PrimitiveInt64Option {
	return func(s *PrimitiveInt64) {
		s.minValue = value
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
		isNil:    false,
		maxValue: -1,
		minValue: -1,
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

	receiver.ValidationMax()
	if receiver.err != nil {
		return
	}

	receiver.ValidationMin()
	if receiver.err != nil {
		return
	}

}

// ______________________________________
func (receiver *PrimitiveInt64) ValidationMax() {
	if receiver.maxValue < 0 { //上限値なし
		return
	}

	if receiver.GetIsNil() {
		receiver.SetErrorString("is nil")
		return
	}

	if receiver.value > receiver.maxValue {
		receiver.SetErrorString("max limitation")
		return
	}
}

// ______________________________________
func (receiver *PrimitiveInt64) ValidationMin() {
	if receiver.minValue < 0 { //下限値なし
		return
	}

	if receiver.GetIsNil() {
		receiver.SetErrorString("is nil")
		return
	}

	if receiver.value < receiver.minValue {
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
