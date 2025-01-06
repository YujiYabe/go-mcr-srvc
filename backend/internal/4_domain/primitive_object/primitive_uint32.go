package primitive_object

import (
	"fmt"
)

type PrimitiveUint32 struct {
	err      error
	value    uint32
	isNil    bool
	MaxValue uint32
	MinValue uint32
}

type PrimitiveUint32Option func(*PrimitiveUint32)

func (receiver *PrimitiveUint32) WithError(
	err error,
) PrimitiveUint32Option {
	return func(s *PrimitiveUint32) {
		s.err = err
	}
}

func (receiver *PrimitiveUint32) WithValue(
	value uint32,
) PrimitiveUint32Option {
	return func(s *PrimitiveUint32) {
		s.value = value
	}
}

func (receiver *PrimitiveUint32) WithIsNil(
	isNil bool,
) PrimitiveUint32Option {
	return func(s *PrimitiveUint32) {
		s.isNil = isNil
	}
}

func (receiver *PrimitiveUint32) WithMaxValue(
	value uint32,
) PrimitiveUint32Option {
	return func(s *PrimitiveUint32) {
		s.MaxValue = value
	}
}

func (receiver *PrimitiveUint32) WithMinValue(
	value uint32,
) PrimitiveUint32Option {
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
		err:      nil,
		value:    0,
		isNil:    false,
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
	receiver.isNil = isNil
}
// --------------------------------------
func (receiver *PrimitiveUint32) GetIsNil() bool {
	return receiver.isNil
}

// --------------------------------------
func (receiver *PrimitiveUint32) GetError() error {
	return receiver.err
}

func (receiver *PrimitiveUint32) SetError(
	err error,
) {
	receiver.err = err
}

func (receiver *PrimitiveUint32) SetErrorString(
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
func (receiver *PrimitiveUint32) GetValue() uint32 {
	if receiver.isNil {
		receiver.SetErrorString("is nil")
		return 0
	}
	return receiver.value
}

func (receiver *PrimitiveUint32) SetValue(value uint32) {
	if receiver.isNil {
		receiver.SetErrorString("is nil")
		return
	}
	receiver.value = value
}

// --------------------------------------
func (receiver *PrimitiveUint32) Validation() error {
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

func (receiver *PrimitiveUint32) ValidationMax() {
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

func (receiver *PrimitiveUint32) ValidationMin() {
	if receiver.MinValue == 0 {
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
