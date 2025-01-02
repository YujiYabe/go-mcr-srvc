package primitive_object

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// PrimitiveString は文字列値に対してバリデーション機能を提供する構造体です。
// nil チェック、長さ制限、禁止文字列のチェックなどの機能を備えています。
type PrimitiveString struct {
	Err       error    // バリデーションエラーを格納
	Value     string   // 実際の文字列値
	IsNil     bool     // nil状態を示すフラグ
	MaxLength int      // 最大文字列長 (-1は制限なし)
	MinLength int      // 最小文字列長 (-1は制限なし)
	SpellList []string // チェック対象の禁止文字列リスト
}

// NewPrimitiveString は指定されたオプションで新しいPrimitiveStringインスタンスを生成します
type PrimitiveStringOption func(*PrimitiveString)

// WithError はエラーを設定するオプションを返します
func (receiver *PrimitiveString) WithError(
	err error,
) PrimitiveStringOption {
	return func(s *PrimitiveString) {
		s.Err = err
	}
}

// WithValue は文字列値を設定するオプションを返します
func (receiver *PrimitiveString) WithValue(value string) PrimitiveStringOption {
	return func(s *PrimitiveString) {
		s.Value = value
	}
}

// WithIsNil はnil状態を設定するオプションを返します
func (receiver *PrimitiveString) WithIsNil(isNil bool) PrimitiveStringOption {
	return func(s *PrimitiveString) {
		s.IsNil = isNil
	}
}

// WithMaxLength は最大文字列長を設定するオプションを返します

func (receiver *PrimitiveString) WithMaxLength(length int) PrimitiveStringOption {
	return func(s *PrimitiveString) {
		s.MaxLength = length
	}
}

// WithMinLength は最小文字列長を設定するオプションを返します

func (receiver *PrimitiveString) WithMinLength(length int) PrimitiveStringOption {
	return func(s *PrimitiveString) {
		s.MinLength = length
	}
}

// WithCheckSpell は禁止文字列リストを設定するオプションを返します
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

func (receiver *PrimitiveString) SetError(
	errString string,
) {
	receiver.Err = fmt.Errorf(
		"error: %s",
		errString,
	)
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

// Validation は全てのバリデーションチェックを実行します
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

// ValidationMax は最大文字列長のチェックを行います
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

// ValidationMin は最小文字列長のチェックを行います
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

// ValidationSpell は禁止文字列のチェックを行います
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

// CheckNil は文字列ポインタのnilチェックを行い、適切な値を返します
func (receiver *PrimitiveString) CheckNil(
	value *string,
) (
	isNil bool,
) {
	isNil = true
	if value != nil {
		isNil = false
	}

	return
}
