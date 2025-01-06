package primitive_object

import "fmt"

type PrimitiveSliceInt struct {
	err       error          // バリデーションエラーを格納
	value     []PrimitiveInt // primitive_object.PrimitiveInt
	isNil     bool           // nil状態を示すフラグ
	MaxLength int            // 最大配列 (-1は制限なし)
	MinLength int            // 最小配列 (-1は制限なし)
}

// NewPrimitiveSliceInt creates a new PrimitiveSliceInt instance
func NewPrimitiveSliceInt(
	valueList []int,
) PrimitiveSliceInt {
	if valueList == nil {
		return PrimitiveSliceInt{
			isNil: true,
		}
	}
	primitiveInt := PrimitiveInt{}

	primitiveIntSlice := make([]PrimitiveInt, len(valueList))
	for index, value := range valueList {
		primitiveIntSlice[index] = *NewPrimitiveInt(
			primitiveInt.WithValue(value),
			primitiveInt.WithIsNil(false),
			// primitiveInt.WithMaxValue(),
			// primitiveInt.WithMinValue(),
		)
	}

	return PrimitiveSliceInt{
		value: primitiveIntSlice,
		isNil: false,
	}
}
func (receiver *PrimitiveSliceInt) SetIsNil(
	isNil bool,
) {
	receiver.isNil = isNil
}

// --------------------------------------
func (receiver *PrimitiveSliceInt) GetError() error {
	return receiver.err
}

func (receiver *PrimitiveSliceInt) SetError(
	err error,
) {
	receiver.err = err
}

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

// ValidationMax は最大文字列長のチェックを行います
func (receiver *PrimitiveSliceInt) ValidationMax() {
	if receiver.MaxLength != -1 && len(receiver.value) > receiver.MaxLength {
		receiver.SetErrorString("max limitation")
	}
}

func (receiver *PrimitiveSliceInt) ValidationMin() {
	if receiver.MinLength != -1 && len(receiver.value) < receiver.MinLength {
		receiver.SetErrorString("min limitation")
	}
}

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

// SetMaxLength sets the maximum allowed length
func (receiver *PrimitiveSliceInt) SetMaxLength(
	maxLength int,
) {
	receiver.MaxLength = maxLength
}

// SetMinLength sets the minimum allowed length
func (receiver *PrimitiveSliceInt) SetMinLength(
	minLength int,
) {
	receiver.MinLength = minLength
}
