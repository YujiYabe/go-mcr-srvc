package primitive_object

import "fmt"

// ______________________________________
type PrimitiveSliceInt struct {
	err       error          // バリデーションエラーを格納
	value     []PrimitiveInt // primitive_object.PrimitiveInt
	isNil     bool           // nil状態を示すフラグ
	maxLength int            // 最大配列 (-1は制限なし)
	minLength int            // 最小配列 (-1は制限なし)
}

// ______________________________________
type PrimitiveSliceIntOption func(*PrimitiveSliceInt)

// ______________________________________
func (receiver *PrimitiveSliceInt) WithError(
	err error,
) PrimitiveSliceIntOption {
	return func(s *PrimitiveSliceInt) {
		s.err = err
	}
}

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
	value int,
) PrimitiveSliceIntOption {
	return func(s *PrimitiveSliceInt) {
		s.maxLength = value
	}
}

// ______________________________________
func (receiver *PrimitiveSliceInt) WithMinLength(
	value int,
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
		err:       nil,
		value:     []PrimitiveInt{},
		isNil:     false,
		maxLength: -1,
		minLength: -1,
	}

	for _, option := range options {
		option(primitiveSliceInt)
	}

	return
}

// ______________________________________
func (receiver *PrimitiveSliceInt) GetIsNil() bool {
	return receiver.isNil
}

// ______________________________________
func (receiver *PrimitiveSliceInt) SetIsNil(
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
func (receiver *PrimitiveSliceInt) GetError() error {
	return receiver.err
}

// ______________________________________
func (receiver *PrimitiveSliceInt) SetError(
	err error,
) {
	receiver.err = err
}

// ______________________________________
func (receiver *PrimitiveSliceInt) SetErrorString(
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
// Validate validates the PrimitiveSliceInt
func (receiver *PrimitiveSliceInt) Validation() error {
	if receiver.isNil {
		return nil
	}

	receiver.ValidationMax()
	if receiver.err != nil {
		return receiver.err
	}

	receiver.ValidationMin()
	if receiver.err != nil {
		return receiver.err
	}

	for _, value := range receiver.value {
		value.Validation()
		if value.GetError() != nil {
			receiver.SetError(value.GetError())
			break
		}
	}

	return nil
}

// ______________________________________
// ValidationMax は最大文字列長のチェックを行います
func (receiver *PrimitiveSliceInt) ValidationMax() {
	if receiver.maxLength != -1 && len(receiver.value) > receiver.maxLength {
		receiver.SetErrorString("max limitation")
	}
}

// ______________________________________
func (receiver *PrimitiveSliceInt) ValidationMin() {
	if receiver.minLength != -1 && len(receiver.value) < receiver.minLength {
		receiver.SetErrorString("min limitation")
	}
}

// ______________________________________
// GetValue returns the underlying int slice
func (receiver *PrimitiveSliceInt) GetValue() []int {
	if receiver.isNil {
		return nil
	}

	result := make([]int, len(receiver.value))
	for i, v := range receiver.value {
		result[i] = v.GetValue()
	}
	return result
}

// ______________________________________
// SetMaxLength sets the maximum allowed length
func (receiver *PrimitiveSliceInt) SetMaxLength(
	maxLength int,
) {
	receiver.maxLength = maxLength
}

// ______________________________________
// SetMinLength sets the minimum allowed length
func (receiver *PrimitiveSliceInt) SetMinLength(
	minLength int,
) {
	receiver.minLength = minLength
}

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
