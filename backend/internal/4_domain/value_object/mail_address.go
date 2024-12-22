package value_object

import (
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

	valueString, isNil := primitiveString.CheckNil(value)

	mailAddress.Content = primitive_object.NewPrimitiveString(
		primitiveString.WithValue(valueString),
		primitiveString.WithIsNil(isNil),
		primitiveString.WithMaxLength(mailAddressLengthMax),
		primitiveString.WithMinLength(mailAddressLengthMin),
		primitiveString.WithCheckSpell(mailAddressCheckSpell),
	)

	err = mailAddress.Content.Validation()
	// メールアドレス自体のバリデーション

	return
}
