package primitive_object

import (
	"fmt"
)

type PrimitiveUint64 struct {
	err      error
	value    uint64
	isNil    bool
	MaxValue uint64
	MinValue uint64
}

type PrimitiveUint64Option func(*PrimitiveUint64)

func (receiver *PrimitiveUint64) WithError(
	err error,
) PrimitiveUint64Option {
	return func(s *PrimitiveUint64) {
		s.err = err
	}
}

func (receiver *PrimitiveUint64) WithValue(
	value uint64,
) PrimitiveUint64Option {
	return func(s *PrimitiveUint64) {
		s.value = value
	}
}

func (receiver *PrimitiveUint64) WithIsNil(
	isNil bool,
) PrimitiveUint64Option {
	return func(s *PrimitiveUint64) {
		s.isNil = isNil
	}
}

func (receiver *PrimitiveUint64) WithMaxValue(
	value uint64,
) PrimitiveUint64Option {
	return func(s *PrimitiveUint64) {
		s.MaxValue = value
	}
}

func (receiver *PrimitiveUint64) WithMinValue(
	value uint64,
) PrimitiveUint64Option {
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
		err:      nil,
		value:    0,
		isNil:    false,
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
func (receiver *PrimitiveUint64) SetIsNil(
	isNil bool,
) {
	receiver.isNil = isNil
}

// --------------------------------------
func (receiver *PrimitiveUint64) GetError() error {
	return receiver.err
}

func (receiver *PrimitiveUint64) SetError(
	err error,
) {
	receiver.err = err
}

func (receiver *PrimitiveUint64) SetErrorString(
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
func (receiver *PrimitiveUint64) GetValue() uint64 {
	if receiver.isNil {
		receiver.SetErrorString("is nil")
		return 0
	}
	return receiver.value
}

func (receiver *PrimitiveUint64) SetValue(value uint64) {
	if receiver.isNil {
		receiver.SetErrorString("is nil")
		return
	}
	receiver.value = value
}

// --------------------------------------
func (receiver *PrimitiveUint64) Validation() error {
	if receiver.isNil {
		return nil
	}

	receiver.ValidationMax()
	if receiver.err != nil {
		return receiver.err
	}

	receiver.ValidationMin()
	if receiver.err != nil {
		return receiver.err
	}

	return nil
}

func (receiver *PrimitiveUint64) ValidationMax() {
	if receiver.MaxValue == 0 {
		// receiver.SetError("max length no defined")
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

func (receiver *PrimitiveUint64) ValidationMin() {
	if receiver.MinValue == 0 {
		// receiver.SetErrorString("min length no defined")
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
