package primitive_object

import (
	"fmt"
)

type PrimitiveBool struct {
	err   error
	value bool
	isNil bool
}

type PrimitiveBoolOption func(*PrimitiveBool)

func (receiver *PrimitiveBool) WithError(
	err error,
) PrimitiveBoolOption {
	return func(s *PrimitiveBool) {
		s.err = err
	}
}

func (receiver *PrimitiveBool) WithValue(
	value bool,
) PrimitiveBoolOption {
	return func(s *PrimitiveBool) {
		s.value = value
	}
}

func (receiver *PrimitiveBool) WithIsNil(
	isNil bool,
) PrimitiveBoolOption {
	return func(s *PrimitiveBool) {
		s.isNil = isNil
	}
}

func NewPrimitiveBool(
	options ...PrimitiveBoolOption,
) (
	primitiveBool *PrimitiveBool,
) {
	primitiveBool = &PrimitiveBool{
		err:   nil,
		value: false,
		isNil: false,
	}

	for _, option := range options {
		option(primitiveBool)
	}

	return
}

func (receiver *PrimitiveBool) SetIsNil(
	isNil bool,
) {
	receiver.isNil = isNil
}

func (receiver *PrimitiveBool) GetError() error {
	return receiver.err
}

func (receiver *PrimitiveBool) SetError(
	err error,
) {
	receiver.err = err
}

func (receiver *PrimitiveBool) SetErrorString(
	errString string,
) {
	receiver.SetError(
		fmt.Errorf(
			"error: %s",
			errString,
		),
	)
}

func (receiver *PrimitiveBool) GetValue() bool {
	if receiver.isNil {
		receiver.SetErrorString("is nil")
		return false
	}
	return receiver.value
}

func (receiver *PrimitiveBool) SetValue(
	value bool,
) {
	if receiver.isNil {
		receiver.SetErrorString("is nil")
		return
	}
	receiver.value = value
}

func (receiver *PrimitiveBool) Validation() {
	// TODO : バリデーションチェックを実装する
}
