package primitive_object

import "fmt"

// --------------------------------------
type PrimitiveInt struct {
	err      error
	value    int
	isNil    bool
	MaxValue int
	MinValue int
}

// --------------------------------------
type PrimitiveIntOption func(*PrimitiveInt)

// --------------------------------------
func (receiver *PrimitiveInt) WithError(
	err error,
) PrimitiveIntOption {
	return func(s *PrimitiveInt) {
		s.err = err
	}
}

// --------------------------------------
func (receiver *PrimitiveInt) WithValue(
	value *int,
) PrimitiveIntOption {
	receiver.SetIsNil(true)
	var resValue int
	if value != nil {
		receiver.SetIsNil(false)
		resValue = *value
	}
	return func(s *PrimitiveInt) {
		s.value = resValue
	}
}

// --------------------------------------
func (receiver *PrimitiveInt) WithIsNil(
	isNil bool,
) PrimitiveIntOption {
	return func(s *PrimitiveInt) {
		s.isNil = isNil
	}
}

// --------------------------------------
func (receiver *PrimitiveInt) WithMaxValue(
	value int,
) PrimitiveIntOption {
	return func(s *PrimitiveInt) {
		s.MaxValue = value
	}
}

// --------------------------------------
func (receiver *PrimitiveInt) WithMinValue(
	value int,
) PrimitiveIntOption {
	return func(s *PrimitiveInt) {
		s.MinValue = value
	}
}

// --------------------------------------
func NewPrimitiveInt(
	options ...PrimitiveIntOption,
) (
	primitiveInt *PrimitiveInt,
) {
	// デフォルト値を設定
	primitiveInt = &PrimitiveInt{
		err:      nil,
		value:    0,
		isNil:    false,
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
func (receiver *PrimitiveInt) GetIsNil() bool {
	return receiver.isNil
}

// --------------------------------------
func (receiver *PrimitiveInt) SetIsNil(
	isNil bool,
) {
	receiver.isNil = isNil
}

// --------------------------------------
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

// --------------------------------------
func (receiver *PrimitiveInt) GetError() error {
	return receiver.err
}

// --------------------------------------
func (receiver *PrimitiveInt) SetError(
	err error,
) {
	receiver.err = err
}

// --------------------------------------
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
	if receiver.isNil {
		return 0
	}

	return receiver.value
}

// --------------------------------------
func (receiver *PrimitiveInt) SetValue(
	value *int,
) {
	if value == nil {
		receiver.SetIsNil(true)
		return
	}
	receiver.SetIsNil(false)
	receiver.value = *value
}

// --------------------------------------
func (receiver *PrimitiveInt) Validation() {

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

// --------------------------------------
func (receiver *PrimitiveInt) ValidationMax() {
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

// --------------------------------------
func (receiver *PrimitiveInt) ValidationMin() {
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
