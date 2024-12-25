package primitive_object

import (
	"fmt"
)

type PrimitiveSliceInt struct {
	value []int64
	max   *int64
	min   *int64
}

func NewPrimitiveSliceInt() *PrimitiveSliceInt {
	return &PrimitiveSliceInt{}
}

func (receiver *PrimitiveSliceInt) GetLength() int {
	return len(receiver.value)
}

func (receiver *PrimitiveSliceInt) SortAsc() *PrimitiveSliceInt {
	return receiver
}

func (receiver *PrimitiveSliceInt) SortDesc() *PrimitiveSliceInt {
	return receiver
}

func (receiver *PrimitiveSliceInt) RemoveDuplicate() *PrimitiveSliceInt {
	return receiver
}

func (receiver *PrimitiveSliceInt) GetValue() []int64 {
	return receiver.value
}

func (receiver *PrimitiveSliceInt) SetValue(v []int64) (*PrimitiveSliceInt, error) {

	for _, value := range v {
		// max があればmax判定確認
		if receiver.GetMax() != nil {
			if err := receiver.ValidationMax(value); err != nil {
				return receiver, fmt.Errorf("error: %s", "validation error")
			}
		}

		// min があればmin判定確認
		if receiver.GetMin() != nil {
			if err := receiver.ValidationMin(value); err != nil {
				return receiver, fmt.Errorf("error: %s", "validation error")
			}
		}
	}

	receiver.value = v
	return receiver, nil
}

func (receiver *PrimitiveSliceInt) GetMax() *int64 {
	return receiver.max
}

func (receiver *PrimitiveSliceInt) GetMin() *int64 {
	return receiver.min
}

func (receiver *PrimitiveSliceInt) SetMax(v int64) *PrimitiveSliceInt {
	receiver.max = &v
	return receiver
}

func (receiver *PrimitiveSliceInt) SetMin(v int64) *PrimitiveSliceInt {
	receiver.min = &v
	return receiver
}

func (receiver *PrimitiveSliceInt) ValidationMax(v int64) error {
	if v >= *receiver.max {
		return fmt.Errorf("error: %s", "max limitation")
	}

	return nil
}

func (receiver *PrimitiveSliceInt) ValidationMin(v int64) error {
	if v >= *receiver.min {
		return fmt.Errorf("error: %s", "min limitation")
	}

	return nil
}
