package primitive_object

import (
	"fmt"
	"math"
)

// --------------------------------------
type PrimitiveInt32 struct {
	err      error
	value    int32
	isNil    bool
	MaxValue int32
	MinValue int32
}

// --------------------------------------
type PrimitiveInt32Option func(*PrimitiveInt32)

// --------------------------------------
func (receiver *PrimitiveInt32) WithError(
	err error,
) PrimitiveInt32Option {
	return func(s *PrimitiveInt32) {
		s.err = err
	}
}

// --------------------------------------
func (receiver *PrimitiveInt32) WithValue(
	value *int32,
) PrimitiveInt32Option {
	receiver.SetIsNil(true)
	var resValue int32
	if value != nil {
		receiver.SetIsNil(false)
		resValue = *value
	}
	return func(s *PrimitiveInt32) {
		s.value = resValue
	}
}

// --------------------------------------
func (receiver *PrimitiveInt32) WithIsNil(
	isNil bool,
) PrimitiveInt32Option {
	return func(s *PrimitiveInt32) {
		s.isNil = isNil
	}
}

// --------------------------------------
func (receiver *PrimitiveInt32) WithMaxValue(
	value int32,
) PrimitiveInt32Option {
	return func(s *PrimitiveInt32) {
		s.MaxValue = value
	}
}

// --------------------------------------
func (receiver *PrimitiveInt32) WithMinValue(
	value int32,
) PrimitiveInt32Option {
	return func(s *PrimitiveInt32) {
		s.MinValue = value
	}
}

// --------------------------------------
func NewPrimitiveInt32(
	options ...PrimitiveInt32Option,
) (
	primitiveInt32 *PrimitiveInt32,
) {
	// デフォルト値を設定
	primitiveInt32 = &PrimitiveInt32{
		err:      nil,
		value:    0,
		isNil:    false,
		MaxValue: -1,
		MinValue: -1,
	}

	// オプションを適用
	for _, option := range options {
		option(primitiveInt32)
	}

	return
}

// --------------------------------------
func (receiver *PrimitiveInt32) GetIsNil() bool {
	return receiver.isNil
}

// --------------------------------------
func (receiver *PrimitiveInt32) SetIsNil(
	isNil bool,
) {
	receiver.isNil = isNil
}

// --------------------------------------
func (receiver *PrimitiveInt32) CheckNil(
	value *int32,
) (
	isNil bool,
) {
	isNil = true
	if value != nil {
		isNil = false
	}
	return
}

// --------------------------------------
func (receiver *PrimitiveInt32) GetError() error {
	return receiver.err
}

// --------------------------------------
func (receiver *PrimitiveInt32) SetError(
	err error,
) {
	receiver.err = err
}

// --------------------------------------
func (receiver *PrimitiveInt32) SetErrorString(
	errString string,
) {
	receiver.SetError(
		fmt.Errorf(
			"error: %s",
			errString,
		),
	)
}

// --------------------------------------
func (receiver *PrimitiveInt32) GetValue() int32 {
	if receiver.isNil {
		receiver.SetErrorString("is nil")
		return 0
	}
	return receiver.value
}

// --------------------------------------
func (receiver *PrimitiveInt32) SetValue(
	value *int32,
) {
	if value == nil {
		receiver.SetIsNil(true)
		return
	}
	receiver.SetIsNil(false)
	receiver.value = *value
}

// 指定した桁で四捨五入するメソッド
// --------------------------------------
func (receiver *PrimitiveInt32) RoundToDigit(
	digit int32,
) {
	if receiver.isNil {
		receiver.SetErrorString("is nil")
		return
	}

	multiplier := int32(1)
	for i := int32(0); i < digit; i++ {
		multiplier *= 10
	}

	value := float64(receiver.value) / float64(multiplier)
	rounded := int32(math.Round(value) * float64(multiplier))
	receiver.value = rounded
}

// --------------------------------------
func (receiver *PrimitiveInt32) Validation() {
	if receiver.isNil {
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

// --------------------------------------
func (receiver *PrimitiveInt32) ValidationMax() {
	if receiver.MaxValue < 0 {
		// receiver.SetErrorString("max length no defined")
		return
	}

	if receiver.isNil {
		receiver.SetErrorString("is nil")
		return
	}

	if receiver.value > receiver.MaxValue {
		receiver.SetErrorString("max limitation")
		return
	}
}

// --------------------------------------
func (receiver *PrimitiveInt32) ValidationMin() {
	if receiver.MinValue < 0 {
		// receiver.SetError("min length no defined")
		return
	}

	if receiver.isNil {
		receiver.SetErrorString("is nil")
		return
	}

	if receiver.value < receiver.MinValue {
		receiver.SetErrorString("min limitation")
		return
	}
}
