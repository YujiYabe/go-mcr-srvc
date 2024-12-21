package primitive_object

import (
	"fmt"
)

type PrimitiveInt struct {
	value  *int
	max    *int
	min    *int
	canNil *bool
}

func NewPrimitiveInt() *PrimitiveInt {
	return &PrimitiveInt{}
}

func (receiver *PrimitiveInt) GetValue() int {
	if receiver.value != nil {
		return *receiver.value
	}

	return 0

}

func (receiver *PrimitiveInt) GetPointer() *int {
	return receiver.value
}

func (receiver *PrimitiveInt) SetValue(v *int) (*PrimitiveInt, error) {
	receiver.value = v

	// nullがnilであれば判定対象外にヌル
	canNil := receiver.GetCanNil()
	if canNil != nil && // canNilに値が入っており
		!*canNil && // canNilがfalseだが[*canNil == false]
		receiver.value == nil { // valueがnilの場合
		return receiver, fmt.Errorf("error: %s", "validation error")
	}

	if canNil != nil && // canNilに値が入っており
		*canNil && // canNilがtrueだが[*canNil == true
		receiver.value == nil { // valueがnilの場合
		return receiver, nil
	}

	if canNil == nil && // canNilに値が入っておらず
		receiver.value == nil { // valueがnilの場合
		return receiver, nil // あえて設定しない要素の為このまま返す
	}

	// max があればmax判定確認
	if receiver.GetMax() != nil {
		if err := receiver.ValidationMax(); err != nil {
			return receiver, fmt.Errorf("error: %s", "validation error")
		}
	}

	// min があればmin判定確認
	if receiver.GetMin() != nil {
		if err := receiver.ValidationMin(); err != nil {
			return receiver, fmt.Errorf("error: %s", "validation error")
		}
	}

	return receiver, nil
}

func (receiver *PrimitiveInt) GetMax() *int {
	return receiver.max
}

func (receiver *PrimitiveInt) GetMin() *int {
	return receiver.min
}

func (receiver *PrimitiveInt) GetCanNil() *bool {
	return receiver.canNil
}

func (receiver *PrimitiveInt) SetMax(v int) *PrimitiveInt {
	receiver.max = &v
	return receiver
}

func (receiver *PrimitiveInt) SetMin(v int) *PrimitiveInt {
	receiver.min = &v
	return receiver
}

func (receiver *PrimitiveInt) SetCanNil(v bool) *PrimitiveInt {
	receiver.canNil = &v
	return receiver
}

func (receiver *PrimitiveInt) ValidationNil() error {
	if receiver.value == nil {
		return fmt.Errorf("error: %s", "is nil")
	}
	return nil
}

func (receiver *PrimitiveInt) ValidationMax() error {
	if err := receiver.ValidationNil(); err != nil {
		return err
	}
	if *receiver.value > *receiver.max {
		return fmt.Errorf("error: %s", "over max limitation")
	}

	return nil
}

func (receiver *PrimitiveInt) ValidationMin() error {
	if err := receiver.ValidationNil(); err != nil {
		return err
	}
	if *receiver.value < *receiver.min {
		return fmt.Errorf("error: %s", "over min limitation")
	}

	return nil
}

func (receiver *PrimitiveInt) IsNil() bool {
	return receiver.value == nil
}
