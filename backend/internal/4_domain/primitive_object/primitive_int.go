package primitive_object

import (
	"fmt"
)

type PrimitiveInt struct {
	Err       error
	Value     int
	IsNil     bool
	MaxLength int
	MinLength int
}

type PrimitiveIntOption func(*PrimitiveInt)

func (receiver *PrimitiveInt) WithError(err error) PrimitiveIntOption {
	return func(s *PrimitiveInt) {
		s.Err = err
	}
}

func (receiver *PrimitiveInt) WithValue(value int) PrimitiveIntOption {
	return func(s *PrimitiveInt) {
		s.Value = value
	}
}

func (receiver *PrimitiveInt) WithIsNil(isNil bool) PrimitiveIntOption {
	return func(s *PrimitiveInt) {
		s.IsNil = isNil
	}
}

func (receiver *PrimitiveInt) WithMaxLength(length int) PrimitiveIntOption {
	return func(s *PrimitiveInt) {
		s.MaxLength = length
	}
}

func (receiver *PrimitiveInt) WithMinLength(length int) PrimitiveIntOption {
	return func(s *PrimitiveInt) {
		s.MinLength = length
	}
}

func NewPrimitiveInt(
	options ...PrimitiveIntOption,
) (
	primitiveInt *PrimitiveInt,
) {
	// デフォルト値を設定
	primitiveInt = &PrimitiveInt{
		Err:       nil,
		Value:     0,
		IsNil:     false,
		MaxLength: -1,
		MinLength: -1,
	}

	// オプションを適用
	for _, option := range options {
		option(primitiveInt)
	}

	return
}

// --------------------------------------
func (receiver *PrimitiveInt) SetIsNil(isNil bool) {
	receiver.IsNil = isNil
}

// --------------------------------------
func (receiver *PrimitiveInt) GetError() error {
	return receiver.Err
}

func (receiver *PrimitiveInt) SetError(errString string) {
	receiver.Err = fmt.Errorf("error: %s", errString)
}

// --------------------------------------
func (receiver *PrimitiveInt) GetValue() int {
	if receiver.IsNil {
		receiver.SetError("is nil")
		return 0
	}
	return receiver.Value
}

func (receiver *PrimitiveInt) SetValue(value int) {
	if receiver.IsNil {
		receiver.SetError("is nil")
		return
	}
	receiver.Value = value
}

// --------------------------------------
func (receiver *PrimitiveInt) Validation() error {
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

func (receiver *PrimitiveInt) ValidationMax() {
	if receiver.MaxLength < 0 {
		// receiver.SetError("max length no defined")
		return
	}

	if receiver.IsNil {
		receiver.SetError("is nil")
		return
	}

	if receiver.Value > receiver.MaxLength {
		receiver.SetError("over max limitation")
		return
	}
}

func (receiver *PrimitiveInt) ValidationMin() {
	if receiver.MinLength < 0 {
		// receiver.SetError("min length no defined")
		return
	}

	if receiver.IsNil {
		receiver.SetError("is nil")
		return
	}

	if receiver.Value < receiver.MinLength {
		receiver.SetError("over min limitation")
		return
	}
}
