package primitive_object

import "fmt"

type PrimitiveInt struct {
	Err      error
	Value    int
	IsNil    bool
	MaxValue int
	MinValue int
}

type PrimitiveIntOption func(*PrimitiveInt)

func (receiver *PrimitiveInt) WithError(
	err error,
) PrimitiveIntOption {
	return func(s *PrimitiveInt) {
		s.Err = err
	}
}

func (receiver *PrimitiveInt) WithValue(
	value int,
) PrimitiveIntOption {
	return func(s *PrimitiveInt) {
		s.Value = value
	}
}

func (receiver *PrimitiveInt) WithIsNil(
	isNil bool,
) PrimitiveIntOption {
	return func(s *PrimitiveInt) {
		s.IsNil = isNil
	}
}

func (receiver *PrimitiveInt) WithMaxValue(
	value int,
) PrimitiveIntOption {
	return func(s *PrimitiveInt) {
		s.MaxValue = value
	}
}

func (receiver *PrimitiveInt) WithMinValue(
	value int,
) PrimitiveIntOption {
	return func(s *PrimitiveInt) {
		s.MinValue = value
	}
}

func NewPrimitiveInt(
	options ...PrimitiveIntOption,
) (
	primitiveInt *PrimitiveInt,
) {
	// デフォルト値を設定
	primitiveInt = &PrimitiveInt{
		Err:      nil,
		Value:    0,
		IsNil:    false,
		MaxValue: -1,
		MinValue: -1,
	}

	// オプションを適用
	for _, option := range options {
		option(primitiveInt)
	}

	return
}

// --------------------------------------
func (receiver *PrimitiveInt) SetIsNil(
	isNil bool,
) {
	receiver.IsNil = isNil
}

// --------------------------------------
func (receiver *PrimitiveInt) GetError() error {
	return receiver.Err
}

func (receiver *PrimitiveInt) SetError(
	err error,
) {
	receiver.Err = err
}

func (receiver *PrimitiveInt) SetErrorString(
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
func (receiver *PrimitiveInt) GetValue() int {
	if receiver.IsNil {
		receiver.SetErrorString("is nil")
		return 0
	}
	return receiver.Value
}

func (receiver *PrimitiveInt) SetValue(
	value int,
) {
	if receiver.IsNil {
		receiver.SetErrorString("is nil")
		return
	}
	receiver.Value = value
}

// --------------------------------------
func (receiver *PrimitiveInt) Validation() {

	if receiver.IsNil {
		return
	}

	receiver.ValidationMax()
	if receiver.Err != nil {
		return
	}

	receiver.ValidationMin()
	if receiver.Err != nil {
		return
	}

}

func (receiver *PrimitiveInt) ValidationMax() {
	if receiver.MaxValue < 0 {
		// receiver.SetError("max length no defined")
		return
	}

	if receiver.IsNil {
		receiver.SetErrorString("is nil")
		return
	}

	if receiver.Value > receiver.MaxValue {
		receiver.SetErrorString("max limitation")
		return
	}
}

func (receiver *PrimitiveInt) ValidationMin() {
	if receiver.MinValue < 0 {
		// receiver.SetErrorString("min length no defined")
		return
	}

	if receiver.IsNil {
		receiver.SetErrorString("is nil")
		return
	}

	if receiver.Value < receiver.MinValue {
		receiver.SetErrorString("min limitation")
		return
	}
}

func (receiver *PrimitiveInt) CheckNil(
	value *int,
) (
	isNil bool,
) {
	isNil = true
	if value != nil {
		isNil = false
	}

	return
}
