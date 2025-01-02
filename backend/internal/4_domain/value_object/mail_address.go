package value_object

import (
	"fmt"
	"regexp"

	"backend/internal/4_domain/primitive_object"
)

const (
	mailAddressLengthMax = 30
	mailAddressLengthMin = 1
)

var mailAddressCheckSpell = []string{}

type MailAddress struct {
	Content *primitive_object.PrimitiveString
}

func NewMailAddress(
	value *string,
) (
	mailAddress MailAddress,
	err error,
) {
	mailAddress = MailAddress{}
	primitiveString := &primitive_object.PrimitiveString{}

	isNil := primitiveString.CheckNil(value)
	valueString := ""
	if !isNil {
		valueString = *value
	}

	mailAddress.Content = primitive_object.NewPrimitiveString(
		primitiveString.WithValue(valueString),
		primitiveString.WithIsNil(isNil),
		primitiveString.WithMaxLength(mailAddressLengthMax),
		primitiveString.WithMinLength(mailAddressLengthMin),
		primitiveString.WithCheckSpell(mailAddressCheckSpell),
	)

	// 文字列そのもののバリデーション
	err = mailAddress.Content.Validation()
	if err != nil {
		return
	}

	// メールアドレスのバリデーション
	err = mailAddress.Validation()
	if err != nil {
		return
	}

	return
}

func (receiver MailAddress) Validation() error {
	if receiver.Content.IsNil {
		return nil
	}

	// メールアドレスの正規表現パターン
	emailPattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	matched, err := regexp.MatchString(emailPattern, receiver.Content.Value)
	if err != nil {
		return fmt.Errorf(
			"failed to validate email format: %w", err)
	}

	if !matched {
		return fmt.Errorf(
			"invalid email format: %s", receiver.Content.Value)
	}

	return nil
}
