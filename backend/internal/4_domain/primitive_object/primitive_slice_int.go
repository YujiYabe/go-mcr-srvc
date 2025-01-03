package primitive_object

import "fmt"

type PrimitiveSliceInt struct {
	Err       error          // バリデーションエラーを格納
	Value     []PrimitiveInt // primitive_object.PrimitiveInt
	IsNil     bool           // nil状態を示すフラグ
	MaxLength int            // 最大配列 (-1は制限なし)
	MinLength int            // 最小配列 (-1は制限なし)
}

// NewPrimitiveSliceInt creates a new PrimitiveSliceInt instance
func NewPrimitiveSliceInt(
	valueList []int,
) PrimitiveSliceInt {
	if valueList == nil {
		return PrimitiveSliceInt{
			IsNil: true,
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
		Value: primitiveIntSlice,
		IsNil: false,
	}
}
func (receiver *PrimitiveSliceInt) SetIsNil(
	isNil bool,
) {
	receiver.IsNil = isNil
}

// --------------------------------------
func (receiver *PrimitiveSliceInt) GetError() error {
	return receiver.Err
}

func (receiver *PrimitiveSliceInt) SetError(
	errString string,
) {
	receiver.Err = fmt.Errorf(
		"error: %s",
		errString,
	)
}

// Validate validates the PrimitiveSliceInt
func (receiver *PrimitiveSliceInt) Validation() error {
	if receiver.IsNil {
		return nil
	}

	receiver.ValidationMax()
	if receiver.Err != nil {
		return receiver.Err
	}

	receiver.ValidationMin()
	if receiver.Err != nil {
		return receiver.Err
	}

	for _, value := range receiver.Value {
		if err := value.Validation(); err != nil {
			return err
		}
	}

	return nil
}

// ValidationMax は最大文字列長のチェックを行います
func (receiver *PrimitiveSliceInt) ValidationMax() {
	if receiver.MaxLength != -1 && len(receiver.Value) > receiver.MaxLength {
		receiver.SetError("max limitation")
	}
}

func (receiver *PrimitiveSliceInt) ValidationMin() {
	if receiver.MinLength != -1 && len(receiver.Value) < receiver.MinLength {
		receiver.SetError("min limitation")
	}
}

// GetValue returns the underlying int slice
func (receiver *PrimitiveSliceInt) GetValue() []int {
	if receiver.IsNil {
		return nil
	}

	result := make([]int, len(receiver.Value))
	for i, v := range receiver.Value {
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
