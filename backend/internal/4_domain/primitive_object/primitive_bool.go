package primitive_object

import (
	"fmt"
)

type PrimitiveBool struct {
	value  *bool
	canNil *bool
}

func NewPrimitiveBool() *PrimitiveBool {
	return &PrimitiveBool{}
}

func (receiver *PrimitiveBool) GetPointer() *bool {
	return receiver.value
}

func (receiver *PrimitiveBool) GetValue() bool {
	if receiver.value != nil {
		return *receiver.value
	}

	return false
}

func (receiver *PrimitiveBool) SetValue(v *bool) (*PrimitiveBool, error) {
	receiver.value = v

	// nullがnilであれば判定対象外にヌル
	canNil := receiver.GetCanNil()
	if canNil != nil && // canNilに値が入っており
		!*canNil && // canNilがfalseだが[*canNil == false]
		receiver.value == nil { // valueがnilの場合
		return receiver, fmt.Errorf("error: %s", "validation error")
	}

	if canNil != nil && // canNilに値が入っており
		*canNil && // canNilがtrueだが[*canNil == true]
		receiver.value == nil { // valueがnilの場合
		return receiver, nil
	}

	if canNil == nil && // canNilに値が入っておらず
		receiver.value == nil { // valueがnilの場合
		return receiver, nil // あえて設定しない要素の為このまま返す
	}

	return receiver, nil
}

func (receiver *PrimitiveBool) GetCanNil() *bool {
	return receiver.canNil
}

func (receiver *PrimitiveBool) SetCanNil(v bool) *PrimitiveBool {
	receiver.canNil = &v
	return receiver
}

func (receiver *PrimitiveBool) ValidationNil() error {
	if receiver.value == nil {
		return fmt.Errorf("error: %s", " is nil")
	}
	return nil
}

func (receiver *PrimitiveBool) IsNil() bool {
	return receiver.value == nil
}
