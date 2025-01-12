package primitive_object

import "fmt"

// ______________________________________
type PrimitiveInt struct {
	err      error
	value    int
	isNil    bool
	maxValue int
	minValue int
}

// ______________________________________
type PrimitiveIntOption func(*PrimitiveInt)

// ______________________________________
func (receiver *PrimitiveInt) WithError(
	err error,
) PrimitiveIntOption {
	return func(s *PrimitiveInt) {
		s.err = err
	}
}

// ______________________________________
func (receiver *PrimitiveInt) WithValue(
	value *int,
) PrimitiveIntOption {
	isNil := receiver.CheckNil(value)
	valueInt := 0
	if !isNil {
		valueInt = *value
	}

	return func(s *PrimitiveInt) {
		s.value = valueInt
		s.isNil = isNil
	}
}

// ______________________________________
func (receiver *PrimitiveInt) WithIsNil(
	isNil bool,
) PrimitiveIntOption {
	return func(s *PrimitiveInt) {
		s.isNil = isNil
	}
}

// ______________________________________
func (receiver *PrimitiveInt) WithMaxValue(
	value int,
) PrimitiveIntOption {
	return func(s *PrimitiveInt) {
		s.maxValue = value
	}
}

// ______________________________________
func (receiver *PrimitiveInt) WithMinValue(
	value int,
) PrimitiveIntOption {
	return func(s *PrimitiveInt) {
		s.minValue = value
	}
}

// ______________________________________
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
		maxValue: -1,
		minValue: -1,
	}

	// オプションを適用
	for _, option := range options {
		option(primitiveInt)
	}

	return
}

// ______________________________________
func (receiver *PrimitiveInt) GetIsNil() bool {
	return receiver.isNil
}

// ______________________________________
func (receiver *PrimitiveInt) SetIsNil(
	isNil bool,
) {
	receiver.isNil = isNil
}

// ______________________________________
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

// ______________________________________
func (receiver *PrimitiveInt) GetError() error {
	return receiver.err
}

// ______________________________________
func (receiver *PrimitiveInt) SetError(
	err error,
) {
	receiver.err = err
}

// ______________________________________
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

// ______________________________________
func (receiver *PrimitiveInt) GetValue() int {
	if receiver.GetIsNil() {
		return 0
	}

	return receiver.value
}

// ______________________________________
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

// ______________________________________
func (receiver *PrimitiveInt) Validation() {

	if receiver.GetIsNil() {
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

// ______________________________________
func (receiver *PrimitiveInt) ValidationMax() {
	if receiver.maxValue < 0 { //上限値なし
		return
	}

	if receiver.GetIsNil() {
		receiver.SetErrorString("is nil")
		return
	}

	if receiver.value > receiver.maxValue {
		receiver.SetErrorString("max limitation")
		return
	}
}

// ______________________________________
func (receiver *PrimitiveInt) ValidationMin() {
	if receiver.minValue < 0 { //下限値なし
		return
	}

	if receiver.GetIsNil() {
		receiver.SetErrorString("is nil")
		return
	}

	if receiver.value < receiver.minValue {
		receiver.SetErrorString("min limitation")
		return
	}
}
