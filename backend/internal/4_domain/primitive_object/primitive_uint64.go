package primitive_object

import (
	"fmt"
)

type PrimitiveUint64 struct {
	Err      error
	Value    uint64
	IsNil    bool
	MaxValue uint64
	MinValue uint64
}

type PrimitiveUint64Option func(*PrimitiveUint64)

func (receiver *PrimitiveUint64) WithError(err error) PrimitiveUint64Option {
	return func(s *PrimitiveUint64) {
		s.Err = err
	}
}

func (receiver *PrimitiveUint64) WithValue(value uint64) PrimitiveUint64Option {
	return func(s *PrimitiveUint64) {
		s.Value = value
	}
}

func (receiver *PrimitiveUint64) WithIsNil(isNil bool) PrimitiveUint64Option {
	return func(s *PrimitiveUint64) {
		s.IsNil = isNil
	}
}

func (receiver *PrimitiveUint64) WithMaxValue(value uint64) PrimitiveUint64Option {
	return func(s *PrimitiveUint64) {
		s.MaxValue = value
	}
}

func (receiver *PrimitiveUint64) WithMinValue(value uint64) PrimitiveUint64Option {
	return func(s *PrimitiveUint64) {
		s.MinValue = value
	}
}

func NewPrimitiveUint64(
	options ...PrimitiveUint64Option,
) (
	primitiveInt64 *PrimitiveUint64,
) {
	// デフォルト値を設定
	primitiveInt64 = &PrimitiveUint64{
		Err:      nil,
		Value:    0,
		IsNil:    false,
		MaxValue: 0,
		MinValue: 0,
	}

	// オプションを適用
	for _, option := range options {
		option(primitiveInt64)
	}

	return
}

// --------------------------------------
func (receiver *PrimitiveUint64) SetIsNil(isNil bool) {
	receiver.IsNil = isNil
}

// --------------------------------------
func (receiver *PrimitiveUint64) GetError() error {
	return receiver.Err
}

func (receiver *PrimitiveUint64) SetError(errString string) {
	receiver.Err = fmt.Errorf("PrimitiveUint64: %s", errString)
}

// --------------------------------------
func (receiver *PrimitiveUint64) GetValue() uint64 {
	if receiver.IsNil {
		receiver.SetError("is nil")
		return 0
	}
	return receiver.Value
}

func (receiver *PrimitiveUint64) SetValue(value uint64) {
	if receiver.IsNil {
		receiver.SetError("is nil")
		return
	}
	receiver.Value = value
}

// --------------------------------------
func (receiver *PrimitiveUint64) Validation() error {
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

func (receiver *PrimitiveUint64) ValidationMax() {
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

func (receiver *PrimitiveUint64) ValidationMin() {
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

func (receiver *PrimitiveUint64) CheckNil(
	value *uint64,
) (
	isNil bool,
) {
	isNil = true
	if value != nil {
		isNil = false
	}

	return
}
