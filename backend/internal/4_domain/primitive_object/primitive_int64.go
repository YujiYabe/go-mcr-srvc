package primitive_object

import (
	"fmt"
)

type PrimitiveInt64 struct {
	Err      error
	Value    int64
	IsNil    bool
	MaxValue int64
	MinValue int64
}

type PrimitiveInt64Option func(*PrimitiveInt64)

func (receiver *PrimitiveInt64) WithError(
	err error,
) PrimitiveInt64Option {
	return func(s *PrimitiveInt64) {
		s.Err = err
	}
}

func (receiver *PrimitiveInt64) WithValue(
	value int64,
) PrimitiveInt64Option {
	return func(s *PrimitiveInt64) {
		s.Value = value
	}
}

func (receiver *PrimitiveInt64) WithIsNil(
	isNil bool,
) PrimitiveInt64Option {
	return func(s *PrimitiveInt64) {
		s.IsNil = isNil
	}
}

func (receiver *PrimitiveInt64) WithMaxValue(
	value int64,
) PrimitiveInt64Option {
	return func(s *PrimitiveInt64) {
		s.MaxValue = value
	}
}

func (receiver *PrimitiveInt64) WithMinValue(
	value int64,
) PrimitiveInt64Option {
	return func(s *PrimitiveInt64) {
		s.MinValue = value
	}
}

func NewPrimitiveInt64(
	options ...PrimitiveInt64Option,
) (
	primitiveInt64 *PrimitiveInt64,
) {
	// デフォルト値を設定
	primitiveInt64 = &PrimitiveInt64{
		Err:      nil,
		Value:    0,
		IsNil:    false,
		MaxValue: -1,
		MinValue: -1,
	}

	// オプションを適用
	for _, option := range options {
		option(primitiveInt64)
	}

	return
}

// --------------------------------------
func (receiver *PrimitiveInt64) SetIsNil(
	isNil bool,
) {
	receiver.IsNil = isNil
}

// --------------------------------------
func (receiver *PrimitiveInt64) GetError() error {
	return receiver.Err
}

func (receiver *PrimitiveInt64) SetError(
	errString string,
) {
	receiver.Err = fmt.Errorf(
		"PrimitiveInt64: %s",
		errString,
	)
}

// --------------------------------------
func (receiver *PrimitiveInt64) GetValue() int64 {
	if receiver.IsNil {
		receiver.SetError("is nil")
		return 0
	}
	return receiver.Value
}

func (receiver *PrimitiveInt64) SetValue(
	value int64,
) {
	if receiver.IsNil {
		receiver.SetError("is nil")
		return
	}
	receiver.Value = value
}

// --------------------------------------
func (receiver *PrimitiveInt64) Validation() error {
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

func (receiver *PrimitiveInt64) ValidationMax() {
	if receiver.MaxValue < 0 {
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

func (receiver *PrimitiveInt64) ValidationMin() {
	if receiver.MinValue < 0 {
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
