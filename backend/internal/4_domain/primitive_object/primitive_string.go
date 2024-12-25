package primitive_object

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

type PrimitiveString struct {
	Err       error
	Value     string
	IsNil     bool
	MaxLength int
	MinLength int
	SpellList []string
}

type PrimitiveStringOption func(*PrimitiveString)

func (receiver *PrimitiveString) WithError(err error) PrimitiveStringOption {
	return func(s *PrimitiveString) {
		s.Err = err
	}
}

func (receiver *PrimitiveString) WithValue(value string) PrimitiveStringOption {
	return func(s *PrimitiveString) {
		s.Value = value
	}
}

func (receiver *PrimitiveString) WithIsNil(isNil bool) PrimitiveStringOption {
	return func(s *PrimitiveString) {
		s.IsNil = isNil
	}
}

func (receiver *PrimitiveString) WithMaxLength(length int) PrimitiveStringOption {
	return func(s *PrimitiveString) {
		s.MaxLength = length
	}
}

func (receiver *PrimitiveString) WithMinLength(length int) PrimitiveStringOption {
	return func(s *PrimitiveString) {
		s.MinLength = length
	}
}

func (receiver *PrimitiveString) WithCheckSpell(spellList []string) PrimitiveStringOption {
	return func(s *PrimitiveString) {
		s.SpellList = spellList
	}
}

func NewPrimitiveString(
	options ...PrimitiveStringOption,
) (
	primitiveString *PrimitiveString,
) {
	// デフォルト値を設定
	primitiveString = &PrimitiveString{
		Err:       nil,
		Value:     "",
		IsNil:     false,
		MaxLength: -1,
		MinLength: -1,
		SpellList: []string{},
	}

	// オプションを適用
	for _, option := range options {
		option(primitiveString)
	}

	return
}

// --------------------------------------
func (receiver *PrimitiveString) SetIsNil(isNil bool) {
	receiver.IsNil = isNil
}

// --------------------------------------
func (receiver *PrimitiveString) GetError() error {
	return receiver.Err
}

func (receiver *PrimitiveString) SetError(errString string) {
	receiver.Err = fmt.Errorf("error: %s", errString)
}

// --------------------------------------
func (receiver *PrimitiveString) GetValue() string {
	if receiver.IsNil {
		receiver.SetError("is nil")
		return ""
	}
	return receiver.Value
}

func (receiver *PrimitiveString) SetValue(value string) {
	// if receiver.IsNil {
	// 	receiver.SetError("is nil")
	// 	return
	// }
	receiver.IsNil = false
	receiver.Value = value
}

// --------------------------------------
func (receiver *PrimitiveString) Validation() error {

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

	receiver.ValidationSpell()
	if receiver.Err != nil {
		return receiver.Err
	}

	return nil
}

func (receiver *PrimitiveString) ValidationMax() {
	if receiver.MaxLength < 0 {
		// receiver.SetError("max length no defined")
		return
	}

	if receiver.IsNil {
		receiver.SetError("is nil")
		return
	}

	if utf8.RuneCountInString(receiver.Value) > receiver.MaxLength {
		receiver.SetError("max limitation")
		return
	}
}

func (receiver *PrimitiveString) ValidationMin() {
	if receiver.MinLength < 0 {
		// receiver.SetError("min length no defined")
		return
	}

	if receiver.IsNil {
		receiver.SetError("is nil")
		return
	}

	if utf8.RuneCountInString(receiver.Value) < receiver.MinLength {
		receiver.SetError("min limitation")
		return
	}
}

func (receiver *PrimitiveString) ValidationSpell() {
	if len(receiver.SpellList) == 0 {
		return
	}
	for _, spell := range receiver.SpellList {
		if strings.Contains(receiver.Value, spell) {
			receiver.SetError("detect target spell : " + spell)
			return
		}
	}
}

func (receiver *PrimitiveString) CheckNil(
	value *string,
) (
	valueString string,
	isNil bool,
) {
	valueString = ""
	isNil = true
	if value != nil {
		valueString = *value
		isNil = false
	}

	return
}
