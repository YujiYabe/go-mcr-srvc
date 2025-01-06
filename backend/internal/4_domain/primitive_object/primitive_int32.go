package primitive_object

import (
	"fmt"
	"math"
)

type PrimitiveInt32 struct {
	Err      error
	Value    int32
	IsNil    bool
	MaxValue int32
	MinValue int32
}

type PrimitiveInt32Option func(*PrimitiveInt32)

func (receiver *PrimitiveInt32) WithError(
	err error,
) PrimitiveInt32Option {
	return func(s *PrimitiveInt32) {
		s.Err = err
	}
}

func (receiver *PrimitiveInt32) WithValue(
	value int32,
) PrimitiveInt32Option {
	return func(s *PrimitiveInt32) {
		s.Value = value
	}
}

func (receiver *PrimitiveInt32) WithIsNil(
	isNil bool,
) PrimitiveInt32Option {
	return func(s *PrimitiveInt32) {
		s.IsNil = isNil
	}
}

func (receiver *PrimitiveInt32) WithMaxValue(
	value int32,
) PrimitiveInt32Option {
	return func(s *PrimitiveInt32) {
		s.MaxValue = value
	}
}

func (receiver *PrimitiveInt32) WithMinValue(
	value int32,
) PrimitiveInt32Option {
	return func(s *PrimitiveInt32) {
		s.MinValue = value
	}
}

func NewPrimitiveInt32(
	options ...PrimitiveInt32Option,
) (
	primitiveInt32 *PrimitiveInt32,
) {
	// デフォルト値を設定
	primitiveInt32 = &PrimitiveInt32{
		Err:      nil,
		Value:    0,
		IsNil:    false,
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
func (receiver *PrimitiveInt32) SetIsNil(
	isNil bool,
) {
	receiver.IsNil = isNil
}

// --------------------------------------
func (receiver *PrimitiveInt32) GetError() error {
	return receiver.Err
}

func (receiver *PrimitiveInt32) SetError(
	err error,
) {
	receiver.Err = err
}

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

func (receiver *PrimitiveInt32) GetValue() int32 {
	if receiver.IsNil {
		receiver.SetErrorString("is nil")
		return 0
	}
	return receiver.Value
}

func (receiver *PrimitiveInt32) SetValue(
	value int32,
) {
	if receiver.IsNil {
		receiver.SetErrorString("is nil")
		return
	}
	receiver.Value = value
}

// 指定した桁で四捨五入するメソッド
func (receiver *PrimitiveInt32) RoundToDigit(
	digit int32,
) {
	if receiver.IsNil {
		receiver.SetErrorString("is nil")
		return
	}

	multiplier := int32(1)
	for i := int32(0); i < digit; i++ {
		multiplier *= 10
	}

	value := float64(receiver.Value) / float64(multiplier)
	rounded := int32(math.Round(value) * float64(multiplier))
	receiver.Value = rounded
}

func (receiver *PrimitiveInt32) Validation() {
	if receiver.IsNil {
		return
	}

	receiver.ValidationMax()
	if receiver.Err != nil {
		return
	}

	receiver.ValidationMin()
	if receiver.Err != nil {
		return
	}

}

func (receiver *PrimitiveInt32) ValidationMax() {
	if receiver.MaxValue < 0 {
		// receiver.SetErrorString("max length no defined")
		return
	}

	if receiver.IsNil {
		receiver.SetErrorString("is nil")
		return
	}

	if receiver.Value > receiver.MaxValue {
		receiver.SetErrorString("max limitation")
		return
	}
}

func (receiver *PrimitiveInt32) ValidationMin() {
	if receiver.MinValue < 0 {
		// receiver.SetError("min length no defined")
		return
	}

	if receiver.IsNil {
		receiver.SetErrorString("is nil")
		return
	}

	if receiver.Value < receiver.MinValue {
		receiver.SetErrorString("min limitation")
		return
	}
}

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
