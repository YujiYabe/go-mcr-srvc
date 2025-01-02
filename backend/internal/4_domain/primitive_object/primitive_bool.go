package primitive_object

import (
	"fmt"
)

type PrimitiveBool struct {
	Err   error
	Value bool
	IsNil bool
}

type PrimitiveBoolOption func(*PrimitiveBool)

func (receiver *PrimitiveBool) WithError(
	err error,
) PrimitiveBoolOption {
	return func(s *PrimitiveBool) {
		s.Err = err
	}
}

func (receiver *PrimitiveBool) WithValue(
	value bool,
) PrimitiveBoolOption {
	return func(s *PrimitiveBool) {
		s.Value = value
	}
}

func (receiver *PrimitiveBool) WithIsNil(
	isNil bool,
) PrimitiveBoolOption {
	return func(s *PrimitiveBool) {
		s.IsNil = isNil
	}
}

func NewPrimitiveBool(
	options ...PrimitiveBoolOption,
) (
	primitiveBool *PrimitiveBool,
) {
	primitiveBool = &PrimitiveBool{
		Err:   nil,
		Value: false,
		IsNil: false,
	}

	for _, option := range options {
		option(primitiveBool)
	}

	return
}

func (receiver *PrimitiveBool) SetIsNil(
	isNil bool,
) {
	receiver.IsNil = isNil
}

func (receiver *PrimitiveBool) GetError() error {
	return receiver.Err
}

func (receiver *PrimitiveBool) SetError(
	errString string,
) {
	receiver.Err = fmt.Errorf(
		"PrimitiveBool: %s",
		errString,
	)
}

func (receiver *PrimitiveBool) GetValue() bool {
	if receiver.IsNil {
		receiver.SetError("is nil")
		return false
	}
	return receiver.Value
}

func (receiver *PrimitiveBool) SetValue(
	value bool,
) {
	if receiver.IsNil {
		receiver.SetError("is nil")
		return
	}
	receiver.Value = value
}

func (receiver *PrimitiveBool) Validation() error {
	if receiver.IsNil {
		return nil
	}
	return receiver.Err
}
