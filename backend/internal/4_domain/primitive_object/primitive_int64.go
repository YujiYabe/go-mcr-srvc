package primitive_object

import (
	"fmt"
)

type PrimitiveInt64 struct {
	err      error
	value    int64
	isNil    bool
	MaxValue int64
	MinValue int64
}

type PrimitiveInt64Option func(*PrimitiveInt64)

func (receiver *PrimitiveInt64) WithError(
	err error,
) PrimitiveInt64Option {
	return func(s *PrimitiveInt64) {
		s.err = err
	}
}

func (receiver *PrimitiveInt64) WithValue(
	value int64,
) PrimitiveInt64Option {
	return func(s *PrimitiveInt64) {
		s.value = value
	}
}

func (receiver *PrimitiveInt64) WithIsNil(
	isNil bool,
) PrimitiveInt64Option {
	return func(s *PrimitiveInt64) {
		s.isNil = isNil
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
		err:      nil,
		value:    0,
		isNil:    false,
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
	receiver.isNil = isNil
}

// --------------------------------------
func (receiver *PrimitiveInt64) GetError() error {
	return receiver.err
}

func (receiver *PrimitiveInt64) SetError(
	err error,
) {
	receiver.err = err
}

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

// --------------------------------------
func (receiver *PrimitiveInt64) GetValue() int64 {
	if receiver.isNil {
		receiver.SetErrorString("is nil")
		return 0
	}
	return receiver.value
}

func (receiver *PrimitiveInt64) SetValue(
	value int64,
) {
	if receiver.isNil {
		receiver.SetErrorString("is nil")
		return
	}
	receiver.value = value
}

// --------------------------------------
func (receiver *PrimitiveInt64) Validation() {
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

func (receiver *PrimitiveInt64) ValidationMax() {
	if receiver.MaxValue < 0 {
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

func (receiver *PrimitiveInt64) ValidationMin() {
	if receiver.MinValue < 0 {
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
