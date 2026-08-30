package primitive_object

import "fmt"

// ______________________________________
type PrimitiveSliceInt struct {
	value     []PrimitiveIntX[int]
	isNil     bool  // nil状態を示すフラグ
	maxLength *uint // 最大列長
	minLength *uint // 最小列長
}

// ______________________________________
type PrimitiveSliceIntOption func(*PrimitiveSliceInt)

// ______________________________________
func (receiver *PrimitiveSliceInt) WithValue(
	value []PrimitiveIntX[int],
) (
	option PrimitiveSliceIntOption,
) {
	isNil := value == nil
	var valueIntSlice []PrimitiveIntX[int]
	if value != nil {
		valueIntSlice = value
	}
	return func(s *PrimitiveSliceInt) {
		s.value = valueIntSlice
		s.isNil = isNil
	}
}

// ______________________________________
func (receiver *PrimitiveSliceInt) WithIsNil(
	isNil bool,
) (
	value PrimitiveSliceIntOption,
) {
	return func(s *PrimitiveSliceInt) {
		s.isNil = isNil
	}
}

// ______________________________________
func (receiver *PrimitiveSliceInt) WithMaxLength(
	value *uint,
) (
	option PrimitiveSliceIntOption,
) {
	return func(s *PrimitiveSliceInt) {
		s.maxLength = value
	}
}

// ______________________________________
func (receiver *PrimitiveSliceInt) WithMinLength(
	value *uint,
) (
	option PrimitiveSliceIntOption,
) {
	return func(s *PrimitiveSliceInt) {
		s.minLength = value
	}
}

func NewPrimitiveSliceInt(
	options ...PrimitiveSliceIntOption,
) (
	primitiveSliceInt *PrimitiveSliceInt,
) {
	primitiveSliceInt = &PrimitiveSliceInt{
		value:     []PrimitiveIntX[int]{},
		isNil:     true,
		maxLength: nil,
		minLength: nil,
	}

	for _, option := range options {
		option(primitiveSliceInt)
	}

	return
}

// ______________________________________
func (receiver PrimitiveSliceInt) GetIsNil() (
	ok bool,
) {
	return receiver.isNil
}

// ______________________________________
func (receiver PrimitiveSliceInt) Count() (
	value int,
) {
	if receiver.GetIsNil() {
		return 0
	}

	return len(receiver.value)
}

// ______________________________________
func (receiver PrimitiveSliceInt) IsEmpty() (
	isEmpty bool,
) {
	return receiver.Count() == 0
}

// ______________________________________
func (receiver PrimitiveSliceInt) HasValue(
	value int,
) (
	hasValue bool,
) {
	if receiver.GetIsNil() {
		return false
	}

	for _, content := range receiver.value {
		if content.Equal(value) {
			return true
		}
	}

	return false
}

// ______________________________________
func (receiver *PrimitiveSliceInt) CheckNil(
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
func (receiver PrimitiveSliceInt) newErrorString(
	errString string,
) (
	err error,
) {
	return fmt.Errorf(
		"error: %s",
		errString,
	)
}

// ______________________________________
// Validate validates the PrimitiveSliceInt
func (receiver PrimitiveSliceInt) Validation() (
	err error,
) {
	if receiver.isNil {
		return nil
	}

	if err := receiver.ValidationMax(); err != nil {
		return err
	}

	if err := receiver.ValidationMin(); err != nil {
		return err
	}

	for _, value := range receiver.value {
		if err := value.Validation(); err != nil {
			return err
		}
	}

	return nil
}

// ______________________________________
// ValidationMax は最大文字列長のチェックを行います
func (receiver PrimitiveSliceInt) ValidationMax() (
	err error,
) {
	if receiver.maxLength == nil {
		return nil
	}

	if receiver.GetIsNil() {
		return nil
	}

	if receiver.Count() > int(*receiver.maxLength) {
		return receiver.newErrorString("max limitation")
	}

	return nil
}

// ______________________________________
func (receiver PrimitiveSliceInt) ValidationMin() (
	err error,
) {
	if receiver.minLength == nil {
		return nil
	}

	if receiver.GetIsNil() {
		return nil
	}

	if receiver.Count() < int(*receiver.minLength) {
		return receiver.newErrorString("min limitation")
	}

	return nil
}

// ______________________________________
// GetValue returns the underlying int slice
func (receiver PrimitiveSliceInt) GetValue() (
	primitiveIntXs []PrimitiveIntX[int],
) {
	if receiver.isNil {
		return nil
	}

	return receiver.value
}

// ______________________________________
// SortDesc sorts the slice in descending order
func (receiver *PrimitiveSliceInt) SortDesc() {
	if receiver.isNil || len(receiver.value) == 0 {
		return
	}

	// Using bubble sort for demonstration
	valueCount := len(receiver.value)
	for outerIndex := 0; outerIndex < valueCount-1; outerIndex++ {
		for innerIndex := 0; innerIndex < valueCount-outerIndex-1; innerIndex++ {
			if receiver.value[innerIndex].GetValue() < receiver.value[innerIndex+1].GetValue() {
				receiver.value[innerIndex], receiver.value[innerIndex+1] = receiver.value[innerIndex+1], receiver.value[innerIndex]
			}
		}
	}
}

// ______________________________________
// SortAsc sorts the slice in ascending order
func (receiver *PrimitiveSliceInt) SortAsc() {
	if receiver.isNil || len(receiver.value) == 0 {
		return
	}

	valueCount := len(receiver.value)
	for outerIndex := 0; outerIndex < valueCount-1; outerIndex++ {
		for innerIndex := 0; innerIndex < valueCount-outerIndex-1; innerIndex++ {
			if receiver.value[innerIndex].GetValue() > receiver.value[innerIndex+1].GetValue() {
				receiver.value[innerIndex], receiver.value[innerIndex+1] = receiver.value[innerIndex+1], receiver.value[innerIndex]
			}
		}
	}
}

// PrimitiveSliceInt を []int に変換して出力する関数
func (receiver *PrimitiveSliceInt) ToSliceInt() (
	values []int,
) {
	if receiver.isNil {
		return nil
	}

	result := make([]int, len(receiver.value))
	for index, value := range receiver.value {
		result[index] = value.GetValue()
	}
	return result

}
