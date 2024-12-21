package primitive_object

import (
	"fmt"
)

type PrimitiveInt64 struct {
	value  *int64
	max    *int64
	min    *int64
	canNil *bool
}

func NewPrimitiveInt64() *PrimitiveInt64 {
	return &PrimitiveInt64{}
}

func (receiver *PrimitiveInt64) GetValue() int64 {
	if receiver.value != nil {
		return *receiver.value
	}

	return 0

}

func (receiver *PrimitiveInt64) GetPointer() *int64 {
	return receiver.value
}

func (receiver *PrimitiveInt64) SetValue(v *int64) (*PrimitiveInt64, error) {
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

func (receiver *PrimitiveInt64) GetMax() *int64 {
	return receiver.max
}

func (receiver *PrimitiveInt64) GetMin() *int64 {
	return receiver.min
}

func (receiver *PrimitiveInt64) GetCanNil() *bool {
	return receiver.canNil
}

func (receiver *PrimitiveInt64) SetMax(v int64) *PrimitiveInt64 {
	receiver.max = &v
	return receiver
}

func (receiver *PrimitiveInt64) SetMin(v int64) *PrimitiveInt64 {
	receiver.min = &v
	return receiver
}

func (receiver *PrimitiveInt64) SetCanNil(v bool) *PrimitiveInt64 {
	receiver.canNil = &v
	return receiver
}

func (receiver *PrimitiveInt64) ValidationNil() error {
	if receiver.value == nil {
		return fmt.Errorf("error: %s", "is nil")
	}
	return nil
}

func (receiver *PrimitiveInt64) ValidationMax() error {
	if err := receiver.ValidationNil(); err != nil {
		return err
	}
	if *receiver.value > *receiver.max {
		return fmt.Errorf("error: %s", "over max limitation")
	}

	return nil
}

func (receiver *PrimitiveInt64) ValidationMin() error {
	if err := receiver.ValidationNil(); err != nil {
		return err
	}
	if *receiver.value < *receiver.min {
		return fmt.Errorf("error: %s", "over min limitation")
	}

	return nil
}

func (receiver *PrimitiveInt64) IsNil() bool {
	return receiver.value == nil
}
