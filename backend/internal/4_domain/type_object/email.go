package type_object

import (
	"fmt"
	"regexp"

	primitiveObject "backend/internal/4_domain/primitive_object"
)

var (
	emailMaxLength uint = 30
	emailMinLength uint = 1
)

var emailCheckSpell = []string{}

const emailPattern = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

var emailRegexp = regexp.MustCompile(emailPattern)

type Email struct {
	content *primitiveObject.PrimitiveString
}

func NewEmail(
	value *string,
) (
	email Email,
	err error,
) {
	email = Email{}
	err = email.setValue(value)

	return
}

func (receiver *Email) setValue(
	value *string,
) (
	err error,
) {
	primitiveString := &primitiveObject.PrimitiveString{}

	receiver.content = primitiveObject.NewPrimitiveString(
		primitiveString.WithValue(value),
		primitiveString.WithMaxLength(&emailMaxLength),
		primitiveString.WithMinLength(&emailMinLength),
		primitiveString.WithCheckSpell(emailCheckSpell),
	)
	if returnedErr := receiver.content.Validation(); returnedErr != nil {
		err = returnedErr

		return
	}

	err = receiver.Validation()

	return
}
func (receiver Email) GetValue() (
	value string,
) {
	value = receiver.content.GetValue()

	return
}

func (receiver *Email) ErrorString(
	errString string,
) (
	err error,
) {
	err = fmt.Errorf("error: %s", errString)

	return
}

func (receiver Email) GetIsNil() (
	ok bool,
) {
	ok = receiver.content.GetIsNil()

	return
}

func (receiver Email) Validation() (
	err error,
) {
	if receiver.GetIsNil() {
		err = nil

		return
	}

	if !emailRegexp.MatchString(receiver.GetValue()) {
		err = fmt.Errorf("invalid email format: %s", receiver.GetValue())

		return
	}
	err = nil

	return
}
