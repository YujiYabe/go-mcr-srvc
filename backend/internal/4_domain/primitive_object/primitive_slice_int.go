package primitive_object

import "fmt"

// ______________________________________
type PrimitiveSliceInt struct {
	value     []PrimitiveInt // primitive_object.PrimitiveInt
	isNil     bool           // nil状態を示すフラグ
	maxLength *uint          // 最大列長
	minLength *uint          // 最小列長
}

// ______________________________________
type PrimitiveSliceIntOption func(*PrimitiveSliceInt)

// ______________________________________
func (receiver *PrimitiveSliceInt) WithValue(
	value []PrimitiveInt,
) PrimitiveSliceIntOption {
	isNil := value == nil
	var valueIntSlice []PrimitiveInt
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
) PrimitiveSliceIntOption {
	return func(s *PrimitiveSliceInt) {
		s.isNil = isNil
	}
}

// ______________________________________
func (receiver *PrimitiveSliceInt) WithMaxLength(
	value *uint,
) PrimitiveSliceIntOption {
	return func(s *PrimitiveSliceInt) {
		s.maxLength = value
	}
}

// ______________________________________
func (receiver *PrimitiveSliceInt) WithMinLength(
	value *uint,
) PrimitiveSliceIntOption {
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
		value:     []PrimitiveInt{},
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
func (receiver PrimitiveSliceInt) GetIsNil() bool {
	return receiver.isNil
}

// ______________________________________
func (receiver *PrimitiveSliceInt) setIsNil(
	isNil bool,
) {
	receiver.isNil = isNil
}

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
) error {
	return fmt.Errorf(
		"error: %s",
		errString,
	)
}

// ______________________________________
// Validate validates the PrimitiveSliceInt
func (receiver PrimitiveSliceInt) Validation() error {
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
func (receiver PrimitiveSliceInt) ValidationMax() error {
	if receiver.maxLength == nil {
		return nil
	}

	if receiver.GetIsNil() {
		return nil
	}

	if len(receiver.value) > int(*receiver.maxLength) {
		return receiver.newErrorString("max limitation")
	}

	return nil
}

// ______________________________________
func (receiver PrimitiveSliceInt) ValidationMin() error {
	if receiver.minLength == nil {
		return nil
	}

	if receiver.GetIsNil() {
		return nil
	}

	if len(receiver.value) < int(*receiver.minLength) {
		return receiver.newErrorString("min limitation")
	}

	return nil
}

// ______________________________________
// GetValue returns the underlying int slice
func (receiver PrimitiveSliceInt) GetValue() []PrimitiveInt {
	if receiver.isNil {
		return nil
	}

	return receiver.value
}

// ______________________________________
func (receiver *PrimitiveSliceInt) setMaxLength(
	maxLength *uint,
) {
	receiver.maxLength = maxLength
}

// ______________________________________
func (receiver *PrimitiveSliceInt) setMinLength(
	minLength *uint,
) {
	receiver.minLength = minLength
}

// ______________________________________
// SortDesc sorts the slice in descending order
func (receiver *PrimitiveSliceInt) SortDesc() {
	if receiver.isNil || len(receiver.value) == 0 {
		return
	}

	// Using bubble sort for demonstration
	n := len(receiver.value)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if receiver.value[j].GetValue() < receiver.value[j+1].GetValue() {
				receiver.value[j], receiver.value[j+1] = receiver.value[j+1], receiver.value[j]
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

	n := len(receiver.value)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if receiver.value[j].GetValue() > receiver.value[j+1].GetValue() {
				receiver.value[j], receiver.value[j+1] = receiver.value[j+1], receiver.value[j]
			}
		}
	}
}

// PrimitiveSliceInt を []int に変換して出力する関数
func (receiver *PrimitiveSliceInt) ToSliceInt() []int {
	if receiver.isNil {
		return nil
	}

	result := make([]int, len(receiver.value))
	for i, v := range receiver.value {
		result[i] = v.GetValue()
	}
	return result

}
