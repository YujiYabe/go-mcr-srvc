package value_object

import (
	"backend/internal/4_domain/primitive_object"
	"fmt"
	"regexp"
)

const (
	mailAddressLengthMax = 30
	mailAddressLengthMin = 1
)

var mailAddressCheckSpell = []string{}

type MailAddress struct {
	Err     error
	Content *primitive_object.PrimitiveString
}

func NewMailAddress(
	value *string,
) (
	mailAddress MailAddress,
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
	mailAddress.Validation()
	if mailAddress.GetError() != nil {
		mailAddress.SetError(mailAddress.Content.GetError())
		return
	}

	// メールアドレスのバリデーション
	mailAddress.Validation()
	if mailAddress.GetError() != nil {
		return
	}

	return
}

func (receiver MailAddress) Validation() {
	if receiver.Content.IsNil {
		return
	}

	// メールアドレスの正規表現パターン
	emailPattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`

	matched, err := regexp.MatchString(emailPattern, receiver.Content.Value)
	if err != nil {
		receiver.SetError(
			fmt.Errorf(
				"failed to validate email format: %w", err,
			),
		)
		return
	}

	if !matched {
		receiver.SetError(
			fmt.Errorf(
				"invalid email format: %s", receiver.Content.Value,
			),
		)
		return
	}
}

func (receiver *MailAddress) GetError() error {
	return receiver.Err
}

func (receiver *MailAddress) SetError(
	err error,
) {
	receiver.Err = err
}

func (receiver *MailAddress) SetErrorString(
	errString string,
) {
	receiver.SetError(
		fmt.Errorf(
			"error: %s",
			errString,
		),
	)
}
