package primitive_object

import (
	"fmt"
	"unicode/utf8"
)

type PrimitiveSliceString struct {
	value []string
	max   *int
	min   *int
}

func NewPrimitiveSliceString() *PrimitiveSliceString {
	return &PrimitiveSliceString{}
}

func (receiver *PrimitiveSliceString) GetLength() int {
	return len(receiver.value)
}

func (receiver *PrimitiveSliceString) SortAsc() *PrimitiveSliceString {
	// sort.Strings(receiver.value)
	return receiver
}

func (receiver *PrimitiveSliceString) SortDesc() *PrimitiveSliceString {
	// sort.Sort(sort.Reverse(sort.StringSlice(receiver.value)))
	return receiver
}

func (receiver *PrimitiveSliceString) RemoveDuplicate() *PrimitiveSliceString {
	return receiver
}

func (receiver *PrimitiveSliceString) SetMax(v int) *PrimitiveSliceString {
	receiver.max = &v
	return receiver
}

func (receiver *PrimitiveSliceString) SetMin(v int) *PrimitiveSliceString {
	receiver.min = &v
	return receiver
}

func (receiver *PrimitiveSliceString) GetValue() []string {
	return receiver.value
}

func (receiver *PrimitiveSliceString) SetValue(v []string) (*PrimitiveSliceString, error) {

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

func (receiver *PrimitiveSliceString) GetMax() *int {
	return receiver.max
}

func (receiver *PrimitiveSliceString) GetMin() *int {
	return receiver.min
}

func (receiver *PrimitiveSliceString) ValidationMax(v string) error {
	if utf8.RuneCountInString(v) > *receiver.max {
		return fmt.Errorf("error: %s", "over max limitation")
	}

	return nil
}

func (receiver *PrimitiveSliceString) ValidationMin(v string) error {
	if utf8.RuneCountInString(v) < *receiver.min {
		return fmt.Errorf("error: %s", "over min limitation")
	}

	return nil
}
