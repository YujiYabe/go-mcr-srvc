package primitive_object

import (
	"fmt"
)

type PrimitiveUint32 struct {
	Err      error
	Value    uint32
	IsNil    bool
	MaxValue uint32
	MinValue uint32
}

type PrimitiveUint32Option func(*PrimitiveUint32)

func (receiver *PrimitiveUint32) WithError(
	err error,
) PrimitiveUint32Option {
	return func(s *PrimitiveUint32) {
		s.Err = err
	}
}

func (receiver *PrimitiveUint32) WithValue(value uint32) PrimitiveUint32Option {
	return func(s *PrimitiveUint32) {
		s.Value = value
	}
}

func (receiver *PrimitiveUint32) WithIsNil(isNil bool) PrimitiveUint32Option {
	return func(s *PrimitiveUint32) {
		s.IsNil = isNil
	}
}

func (receiver *PrimitiveUint32) WithMaxValue(value uint32) PrimitiveUint32Option {
	return func(s *PrimitiveUint32) {
		s.MaxValue = value
	}
}

func (receiver *PrimitiveUint32) WithMinValue(value uint32) PrimitiveUint32Option {
	return func(s *PrimitiveUint32) {
		s.MinValue = value
	}
}

func NewPrimitiveUint32(
	options ...PrimitiveUint32Option,
) (
	primitiveInt32 *PrimitiveUint32,
) {
	// デフォルト値を設定
	primitiveInt32 = &PrimitiveUint32{
		Err:      nil,
		Value:    0,
		IsNil:    false,
		MaxValue: 0,
		MinValue: 0,
	}

	// オプションを適用
	for _, option := range options {
		option(primitiveInt32)
	}

	return
}

// --------------------------------------
func (receiver *PrimitiveUint32) SetIsNil(isNil bool) {
	receiver.IsNil = isNil
}

// --------------------------------------
func (receiver *PrimitiveUint32) GetError() error {
	return receiver.Err
}

func (receiver *PrimitiveUint32) SetError(
	errString string,
) {
	receiver.Err = fmt.Errorf(
		"PrimitiveUint32: %s",
		errString,
	)
}

// --------------------------------------
func (receiver *PrimitiveUint32) GetValue() uint32 {
	if receiver.IsNil {
		receiver.SetError("is nil")
		return 0
	}
	return receiver.Value
}

func (receiver *PrimitiveUint32) SetValue(value uint32) {
	if receiver.IsNil {
		receiver.SetError("is nil")
		return
	}
	receiver.Value = value
}

// --------------------------------------
func (receiver *PrimitiveUint32) Validation() error {
	if receiver.IsNil {
		return nil
	}

	receiver.ValidationMax()
	if receiver.Err != nil {
		return receiver.Err
	}

	receiver.ValidationMin()
	if receiver.Err != nil {
		return receiver.Err
	}

	return nil
}

func (receiver *PrimitiveUint32) ValidationMax() {
	if receiver.MaxValue == 0 {
		// receiver.SetError("max length no defined")
		return
	}

	if receiver.IsNil {
		receiver.SetError("is nil")
		return
	}

	if receiver.Value > receiver.MaxValue {
		receiver.SetError("max limitation")
		return
	}
}

func (receiver *PrimitiveUint32) ValidationMin() {
	if receiver.MinValue == 0 {
		// receiver.SetError("min length no defined")
		return
	}

	if receiver.IsNil {
		receiver.SetError("is nil")
		return
	}

	if receiver.Value < receiver.MinValue {
		receiver.SetError("min limitation")
		return
	}
}

func (receiver *PrimitiveUint32) CheckNil(
	value *uint32,
) (
	isNil bool,
) {
	isNil = true
	if value != nil {
		isNil = false
	}

	return
}
