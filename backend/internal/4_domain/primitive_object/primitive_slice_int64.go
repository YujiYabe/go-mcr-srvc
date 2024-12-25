package primitive_object

import (
	"fmt"
)

type PrimitiveSliceInt64 struct {
	value []int64
	max   *int64
	min   *int64
}

func NewPrimitiveSliceInt64() *PrimitiveSliceInt64 {
	return &PrimitiveSliceInt64{}
}

func (receiver *PrimitiveSliceInt64) GetLength() int {
	return len(receiver.value)
}

func (receiver *PrimitiveSliceInt64) SortAsc() *PrimitiveSliceInt64 {
	return receiver
}

func (receiver *PrimitiveSliceInt64) SortDesc() *PrimitiveSliceInt64 {
	return receiver
}

func (receiver *PrimitiveSliceInt64) RemoveDuplicate() *PrimitiveSliceInt64 {
	return receiver
}

func (receiver *PrimitiveSliceInt64) GetValue() []int64 {
	return receiver.value
}

func (receiver *PrimitiveSliceInt64) SetValue(v []int64) (*PrimitiveSliceInt64, error) {

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

func (receiver *PrimitiveSliceInt64) GetMax() *int64 {
	return receiver.max
}

func (receiver *PrimitiveSliceInt64) GetMin() *int64 {
	return receiver.min
}

func (receiver *PrimitiveSliceInt64) SetMax(v int64) *PrimitiveSliceInt64 {
	receiver.max = &v
	return receiver
}

func (receiver *PrimitiveSliceInt64) SetMin(v int64) *PrimitiveSliceInt64 {
	receiver.min = &v
	return receiver
}

func (receiver *PrimitiveSliceInt64) ValidationMax(v int64) error {
	if v >= *receiver.max {
		return fmt.Errorf("error: %s", "max limitation")
	}

	return nil
}

func (receiver *PrimitiveSliceInt64) ValidationMin(v int64) error {
	if v >= *receiver.min {
		return fmt.Errorf("error: %s", "min limitation")
	}

	return nil
}
