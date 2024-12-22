package primitive_object

import (
	"fmt"
	"unicode/utf8"
)

type PrimitiveString struct {
	Err       error
	Value     string
	IsNil     bool
	MaxLength int
	MinLength int
}

type PrimitiveStringOption func(*PrimitiveString)

func WithError(err error) PrimitiveStringOption {
	return func(s *PrimitiveString) {
		s.Err = err
	}
}

func WithValue(value string) PrimitiveStringOption {
	return func(s *PrimitiveString) {
		s.Value = value
	}
}

func WithIsNil(value string) PrimitiveStringOption {
	return func(s *PrimitiveString) {
		s.Value = value
	}
}

func WithMaxLength(length int) PrimitiveStringOption {
	return func(s *PrimitiveString) {
		s.MaxLength = length
	}
}

func WithMinLength(length int) PrimitiveStringOption {
	return func(s *PrimitiveString) {
		s.MinLength = length
	}
}

func NewPrimitiveString(
	options ...PrimitiveStringOption,
) (
	primitiveString *PrimitiveString,
) {
	// デフォルト値を設定
	primitiveString = &PrimitiveString{
		Err:       nil,
		Value:     "",
		IsNil:     false,
		MaxLength: -1,
		MinLength: -1,
	}

	// オプションを適用
	for _, option := range options {
		option(primitiveString)
	}

	return
}

// --------------------------------------
func (receiver *PrimitiveString) SetIsNil(isNil bool) {
	receiver.IsNil = isNil
}

// --------------------------------------
func (receiver *PrimitiveString) GetError() error {
	return receiver.Err
}

func (receiver *PrimitiveString) SetError(errString string) {
	receiver.Err = fmt.Errorf("error: %s", errString)
}

// --------------------------------------
func (receiver *PrimitiveString) GetValue() string {
	if receiver.IsNil {
		receiver.SetError("is nil")
		return ""
	}
	return receiver.Value
}

func (receiver *PrimitiveString) SetValue(value string) {
	if receiver.IsNil {
		receiver.SetError("is nil")
		return
	}
	receiver.Value = value
}

// --------------------------------------
func (receiver *PrimitiveString) Validation() {
	receiver.ValidationMax()
	if receiver.Err != nil {
		return
	}

	receiver.ValidationMin()
	if receiver.Err != nil {
		return
	}
}

func (receiver *PrimitiveString) ValidationMax() {
	if receiver.MaxLength < 0 {
		receiver.SetError("max length no defined")
		return
	}

	if receiver.IsNil {
		receiver.SetError("is nil")
		return
	}

	if utf8.RuneCountInString(receiver.Value) > receiver.MaxLength {
		receiver.SetError("over max limitation")
		return
	}
}

func (receiver *PrimitiveString) ValidationMin() {
	if receiver.MinLength < 0 {
		receiver.SetError("min length no defined")
		return
	}

	if receiver.IsNil {
		receiver.SetError("is nil")
		return
	}

	if utf8.RuneCountInString(receiver.Value) < receiver.MaxLength {
		receiver.SetError("over max limitation")
		return
	}
}
